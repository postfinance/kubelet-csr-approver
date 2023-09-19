/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package controller contains the code of the CSR-approver controller.
package controller

import (
	"context"
	"fmt"
	"strings"

	"go4.org/netipx"
	certificatesv1 "k8s.io/api/certificates/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// HostResolver is used to resolve a Host with the LookupHost function
type HostResolver interface {
	LookupHost(context.Context, string) ([]string, error)
}

// Config holds all variables needed to configure the controller
type Config struct {
	LogLevel               int
	MetricsAddr            string
	ProbeAddr              string
	LeaderElection         bool
	RegexStr               string
	ProviderRegexp         func(string) bool
	IPPrefixesStr          string
	ProviderIPSet          *netipx.IPSet
	MaxExpirationSeconds   int32
	K8sConfig              *rest.Config
	DNSResolver            HostResolver
	BypassDNSResolution    bool
	IgnoreNonSystemNodeCsr bool
	AllowedDNSNames        int
	BypassHostnameCheck    bool
}

// CertificateSigningRequestReconciler reconciles a CertificateSigningRequest object
type CertificateSigningRequestReconciler struct {
	ClientSet *clientset.Clientset
	client.Client
	Scheme *runtime.Scheme
	Config
}

//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=get;watch;list
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests/approval,verbs=update
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,resourceNames="kubernetes.io/kubelet-serving",verbs=approve

// Reconcile will perform a series of checks before deciding whether the CSR should be approved or denied
// cyclomatic complexity is high (over 15), but this improves
// readibility for the programmer, therefore we ignore the linting error
//
//nolint:gocyclo // see above
func (r *CertificateSigningRequestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, returnErr error) {
	l := log.FromContext(ctx)

	var csr certificatesv1.CertificateSigningRequest
	if err := r.Client.Get(ctx, req.NamespacedName, &csr); err != nil {
		if apierrors.IsNotFound(err) {
			// we'll ignore not-found errors, since we can get them on deleted requests.
			return res, returnErr
		}

		l.Error(err, "Unable to get CSR with name: %s", req.Name)

		return res, returnErr
	}

	// baseline CSR checks - triage to ignore CSR we should process
	if csr.Spec.SignerName != certificatesv1.KubeletServingSignerName {
		l.V(4).Info("Ignoring non-kubelet-serving CSR.")
		return res, returnErr
	}

	if approved, denied := GetCertApprovalCondition(&csr.Status); approved || denied {
		l.V(3).Info("The CSR is already approved|denied. Ignoring", "approved", approved, "denied", denied)
		return res, returnErr
	}

	if len(csr.Status.Certificate) > 0 {
		l.V(3).Info("The CSR is already signed. No need to do anything else.")
		return res, returnErr
	}

	// actual CSR and x509 CR checks
	x509cr, err := ParseCSR(csr.Spec.Request)
	if err != nil {
		l.Error(err, fmt.Sprintf("unable to parse csr %q", csr.Name))
		return res, returnErr
	}

	if !strings.HasPrefix(csr.Spec.Username, "system:node:") {
		if r.IgnoreNonSystemNodeCsr {
			l.V(0).Info("Ignoring a CSR with username different than system:node:")
			return res, returnErr
		}

		reason := "CSR Spec.Username is not prefixed with system:node:"
		l.V(0).Info("Denying kubelet-serving CSR. Reason:" + reason)

		appendCondition(&csr, false, reason)
	} else if len(x509cr.DNSNames)+len(x509cr.IPAddresses) == 0 {
		reason := "The x509 Cert Request SAN contains neither an IP address nor a DNS name"
		l.V(0).Info("Denying kubelet-serving CSR. Reason:" + reason)

		appendCondition(&csr, false, reason)
	} else if x509cr.Subject.CommonName != csr.Spec.Username {
		reason := "CSR username does not match the parsed x509 certificate request commonname"
		l.V(0).Info("Denying kubelet-serving CSR. Reason:"+reason,
			"commonName", x509cr.Subject.CommonName, "specUsername", csr.Spec.Username)

		appendCondition(&csr, false, reason)
	} else if valid, reason, err := r.DNSCheck(ctx, &csr, x509cr); !valid {
		if err != nil {
			l.V(0).Error(err, reason)
			return res, err // returning a non-nil error to make this request be processed again in the reconcile function
		}
		l.V(0).Info("Denying kubelet-serving CSR. DNS checks failed. Reason:" + reason)

		appendCondition(&csr, false, reason)
	} else if valid, reason, err := r.WhitelistedIPCheck(&csr, x509cr); !valid {
		if err != nil {
			l.V(0).Error(err, reason)
			return res, err // returning a non-nil error to make this request be processed again in the reconcile function
		}
		l.V(0).Info("Denying kubelet-serving CSR. IP whitelist check failed. Reason:" + reason)
		appendCondition(&csr, false, reason)
	} else if csr.Spec.ExpirationSeconds != nil && *csr.Spec.ExpirationSeconds > r.MaxExpirationSeconds {
		reason := "CSR spec.expirationSeconds is longer than the maximum allowed expiration second"
		l.V(0).Info("Denying kubelet-serving CSR. Reason:" + reason)

		appendCondition(&csr, false, reason)
	} else if valid, reason := ProviderChecks(&csr, x509cr); !valid {
		l.V(0).Info("CSR request did not pass the provider-specific tests. Reason: " + reason)
		appendCondition(&csr, false, reason)
	} else {
		l.V(0).Info("CSR approved")
		appendCondition(&csr, true, "")
	}

	_, err = r.ClientSet.CertificatesV1().CertificateSigningRequests().UpdateApproval(ctx, req.Name, &csr, metav1.UpdateOptions{})

	if apierrors.IsConflict(err) || apierrors.IsNotFound(err) {
		// The CSR has been updated or deleted since we read it.
		// Requeue the CSR to try to reconciliate again.
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		l.Error(err, "Couldn't update the CSR to include the approval")
		return ctrl.Result{}, err
	}

	return res, nil
}

func appendCondition(csr *certificatesv1.CertificateSigningRequest, approved bool, reason string) {
	if approved {
		csr.Status.Conditions = append(csr.Status.Conditions, certificatesv1.CertificateSigningRequestCondition{
			Type:               certificatesv1.CertificateApproved,
			Status:             corev1.ConditionTrue,
			Reason:             "kubelet-serving cert validated",
			Message:            "CSR complied with kubelet-csr-approver validation process",
			LastUpdateTime:     metav1.Now(),
			LastTransitionTime: metav1.Time{},
		})
	} else {
		csr.Status.Conditions = append(csr.Status.Conditions, certificatesv1.CertificateSigningRequestCondition{
			Type:               certificatesv1.CertificateDenied,
			Status:             corev1.ConditionTrue,
			Reason:             "kubelet-serving cert denied",
			Message:            "CSR not complying with kubelet-csr-approver validation process. Reason: " + reason,
			LastUpdateTime:     metav1.Now(),
			LastTransitionTime: metav1.Time{},
		})
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *CertificateSigningRequestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&certificatesv1.CertificateSigningRequest{}).
		Complete(r)
}
