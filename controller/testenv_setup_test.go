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

package controller_test

import (
	"context"
	"encoding/pem"
	"net"
	"os"
	"regexp"
	"testing"
	"time"

	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"

	mockdns "github.com/foxcpp/go-mockdns"
	"github.com/postfinance/kubelet-csr-approver/controller"
	"github.com/thanhpk/randstr"
	certificates_v1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

var testEnv *envtest.Environment
var cfg *rest.Config
var k8sClient client.Client
var adminClientset *clientset.Clientset
var dnsResolver mockdns.Resolver

var testContext context.Context
var testContextCancel context.CancelFunc

func waitCsrApprovalStatus(csrName string) (approved, denied bool, err error) {
	for i := 0; i < 3; i++ {
		time.Sleep(250 * time.Millisecond)
		csr, err := adminClientset.CertificatesV1().CertificateSigningRequests().
			Get(testContext, csrName, metav1.GetOptions{})
		if err != nil {
			continue
		}

		approved, denied = controller.GetCertApprovalCondition(&csr.Status)
		if approved || denied {
			break
		}
	}
	return
}

type CsrParams struct {
	csrName     string
	commonName  string
	dnsName     string
	nodeName    string
	ipAddresses []net.IP
}

var (
	testNodeName        string
	testNodeIps         []string
	testNodeIpAddresses []net.IP
)

func createCsr(t *testing.T, params CsrParams) certificates_v1.CertificateSigningRequest {
	csr := certificates_v1.CertificateSigningRequest{}
	if len(params.csrName) == 0 {
		csr.GenerateName = "csr-"
		csr.Name = csr.GenerateName + randstr.String(4)
	} else {
		csr.Name = params.csrName
	}

	if len(params.nodeName) == 0 {
		params.nodeName = randstr.String(4, "0123456789abcdefghijklmnopqrstuvwxyz") + ".test.ch"
	}

	if len(params.dnsName) == 0 {
		params.dnsName = params.nodeName
	}

	csr.Spec.SignerName = certificates_v1.KubeletServingSignerName
	csr.Spec.Usages = append(csr.Spec.Usages,
		certificates_v1.UsageDigitalSignature,
		certificates_v1.UsageKeyEncipherment,
		certificates_v1.UsageServerAuth,
	)
	csr.Spec.Username = "system:node:" + params.nodeName

	if len(params.commonName) == 0 {
		params.commonName = csr.Spec.Username
	}
	_, priv, _ := ed25519.GenerateKey(rand.Reader)
	x509RequestTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			Organization: []string{"system:nodes"},
			CommonName:   params.commonName,
		},
		DNSNames:    []string{params.dnsName},
		IPAddresses: params.ipAddresses,
	}
	x509Request, _ := x509.CreateCertificateRequest(rand.Reader, &x509RequestTemplate, priv)
	pemRequest := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: x509Request,
	})

	csr.Spec.Request = pemRequest
	return csr
}

func createControlPlaneUser(t *testing.T, username string, groups []string) (*rest.Config, *clientset.Clientset, error) {
	userInfo := envtest.User{Name: username, Groups: groups}
	userCfg, err := testEnv.ControlPlane.AddUser(userInfo, cfg)
	if err != nil {
		t.Fatalf("Could not create a ControlPlane User with username:%s, and groups: %v", username, groups)
	}

	userClientSet, err := clientset.NewForConfig(userCfg.Config())
	if err != nil {
		t.Fatalf("Could not create a k8s/ClientSet. Error message: %v", err)
	}
	return userCfg.Config(), userClientSet, err
}

func packageSetup() {

	testNodeIps := []string{"192.168.14.34"}
	for _, ip := range testNodeIps {
		testNodeIpAddresses = append(testNodeIpAddresses, net.ParseIP(ip))
	}
	testNodeName = randstr.String(4, "0123456789abcdefghijklmnopqrstuvwxyz") + ".test.ch"

	testContext, testContextCancel = context.WithCancel(context.Background())
	log.Println("Setting up the testing K8s Control plane -- envtest")
	testEnv = &envtest.Environment{}

	cfg, err := testEnv.Start()
	if err != nil {
		log.Fatalf("Could not start envtest, exiting. Error output:\n %v", err)
	}

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		log.Fatalf("Could not create a k8sClient, exiting. Error output:\n %v", err)
	}

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{})
	if err != nil {
		log.Fatalf("unable to create controller-runtime manager. Error:\n%v", err)
	}

	provRegexp := regexp.MustCompile(`^\w*\.test\.ch$`)
	dnsResolver = mockdns.Resolver{
		Zones: map[string]mockdns.Zone{
			testNodeName + ".": {
				A: testNodeIps,
			},
		},
	}

	adminClientset = clientset.NewForConfigOrDie(cfg)
	csrController := controller.CertificateSigningRequestReconciler{
		ClientSet:      adminClientset,
		Client:         mgr.GetClient(),
		Scheme:         mgr.GetScheme(),
		ProviderRegexp: provRegexp.MatchString,
		Resolver:       &dnsResolver,
	}
	csrController.SetupWithManager(mgr)

	go mgr.Start(testContext)
}

func packageTeardown() {
	log.Println("Teardown: Canceling testContext and stopping the testing CP (testEnv)")
	testContextCancel()
	testEnv.Stop()
}

func TestMain(m *testing.M) {
	packageSetup()
	resultCode := m.Run()
	packageTeardown()
	os.Exit(resultCode)
}
