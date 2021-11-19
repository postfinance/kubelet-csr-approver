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

	"github.com/go-logr/logr"

	certificatesv1 "k8s.io/api/certificates/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// HostResolver is used to resolve a Host with the LookupHost function
type HostResolver interface {
	LookupHost(context.Context, string) ([]string, error)
}

// CertificateSigningRequestReconciler reconciles a CertificateSigningRequest object
type CertificateSigningRequestReconciler struct {
	ClientSet *clientset.Clientset
	client.Client
	Scheme         *runtime.Scheme
	ProviderRegexp func(string) bool
	Resolver       HostResolver
}

//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests,verbs=get;watch;list
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=certificatesigningrequests/approval,verbs=update
//+kubebuilder:rbac:groups=certificates.k8s.io,resources=signers,resourceNames="kubernetes.io/kubelet-serving",verbs=approve

// Reconcile will perform a series of checks before deciding whether the CSR should be approved or denied
func (r *CertificateSigningRequestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, returnErr error) {
	l := log.FromContext(ctx)

	var csr certificatesv1.CertificateSigningRequest
	if err := r.Get(ctx, req.NamespacedName, &csr); err != nil {
		if apierrors.IsNotFound(err) {
			// we'll ignore not-found errors, since we can get them on deleted requests.
			return
		}

		l.Error(err, "Unable to get CSR with name: %s", req.Name)

		return
	}

	if passed := baselineCsrChecks(l, &csr); !passed {
		return
	}

	x509cr, err := ParseCSR(csr.Spec.Request)
	if err != nil {
		l.Error(err, fmt.Sprintf("unable to parse csr %q", csr.Name))
		return
	}

	if !strings.HasPrefix(csr.Spec.Username, "system:node:") {
		reason := "CSR Spec.Username is not prefixed with system:node:"
		l.V(0).Info("Denying kubelet-serving CSR. Reason:" + reason)

		appendCondition(&csr, false, reason)
	} else if x509cr.Subject.CommonName != csr.Spec.Username {
		reason := "CSR username does not match the parsed x509 certificate request commonname"
		l.V(0).Info("Denying kubelet-serving CSR. Reason:"+reason,
			"commonName", x509cr.Subject.CommonName, "specUsername", csr.Spec.Username)

		appendCondition(&csr, false, reason)
	} else if valid, reason, err := r.RegexIPChecks(ctx, &csr, x509cr); !valid {
		if err != nil {
			l.V(0).Error(err, reason)
			return res, err // returning a non-nil error to make this request be processed again in the reconcile function
		}
		l.V(0).Info("Denying kubelet-serving CSR. Regex/IP checks failed. Reason:" + reason)

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

func baselineCsrChecks(l logr.Logger, csr *certificatesv1.CertificateSigningRequest) (passed bool) {
	passed = false

	if csr.Spec.SignerName != certificatesv1.KubeletServingSignerName {
		l.V(4).Info("Ignoring non-kubelet-serving CSR.")
		return
	}

	if approved, denied := GetCertApprovalCondition(&csr.Status); approved || denied {
		l.V(3).Info("The CSR is already approved|denied. Ignoring", "approved", approved, "denied", denied)
		return
	}

	if len(csr.Status.Certificate) > 0 {
		l.V(3).Info("The CSR is already signed. No need to do anything else.")
		return
	}

	return true
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
