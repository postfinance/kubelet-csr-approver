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
	"strings"
	"testing"
	"time"

	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"

	mockdns "github.com/foxcpp/go-mockdns"
	"github.com/postfinance/kubelet-csr-approver/internal/cmd"
	"github.com/postfinance/kubelet-csr-approver/internal/controller"

	"github.com/thanhpk/randstr"
	capiv1 "k8s.io/api/certificates/v1"
	certificates_v1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

var testEnv *envtest.Environment
var cfg *rest.Config
var k8sClient client.Client
var adminClientset *clientset.Clientset
var dnsResolver mockdns.Resolver
var csrController *controller.CertificateSigningRequestReconciler

var testContext context.Context
var testContextCancel context.CancelFunc

func waitCsrApprovalStatus(csrName string) (approved, denied bool, reason string, err error) {
	for i := 0; i < 3; i++ {
		time.Sleep(250 * time.Millisecond)
		csr, err := adminClientset.CertificatesV1().CertificateSigningRequests().
			Get(testContext, csrName, metav1.GetOptions{})
		if err != nil {
			continue
		}

		for _, c := range csr.Status.Conditions {
			if c.Type == capiv1.CertificateApproved {
				approved = true
				reason = c.Message
			}

			if c.Type == capiv1.CertificateDenied {
				denied = true
				reason = c.Message

			}
		}
		if approved || denied {
			break
		}
	}
	return
}

type CsrParams struct {
	csrName           string
	commonName        string
	dnsName           string
	nodeName          string
	username          string
	ipAddresses       []net.IP
	expirationSeconds int32
}

var (
	testNodeName        string
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
		params.nodeName = randstr.String(4, "0123456789abcdefghijklmnopqrstuvwxyz")
	}

	csr.Spec.SignerName = certificates_v1.KubeletServingSignerName
	csr.Spec.Usages = append(csr.Spec.Usages,
		certificates_v1.UsageDigitalSignature,
		certificates_v1.UsageKeyEncipherment,
		certificates_v1.UsageServerAuth,
	)

	if len(params.username) == 0 {
		csr.Spec.Username = "system:node:" + params.nodeName
	} else {
		csr.Spec.Username = params.username
	}

	if len(params.commonName) == 0 {
		params.commonName = csr.Spec.Username
	}

	if params.expirationSeconds > 0 {
		csr.Spec.ExpirationSeconds = &params.expirationSeconds
	}

	_, priv, _ := ed25519.GenerateKey(rand.Reader)
	x509RequestTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			Organization: []string{"system:nodes"},
			CommonName:   params.commonName,
		},
		IPAddresses: params.ipAddresses,
	}
	if len(params.dnsName) > 0 {
		x509RequestTemplate.DNSNames = strings.Split(params.dnsName, ",")
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
	adminClientset = clientset.NewForConfigOrDie(cfg)

	testNodeIpv4 := []string{"192.168.14.34"}
	testNodeIpv6 := []string{"fc00:1291:feed::cafe"}
	for _, ip := range testNodeIpv4 {
		testNodeIpAddresses = append(testNodeIpAddresses, net.ParseIP(ip))
	}
	for _, ip := range testNodeIpv6 {
		testNodeIpAddresses = append(testNodeIpAddresses, net.ParseIP(ip))
	}
	testNodeName = randstr.String(4, "0123456789abcdefghijklmnopqrstuvwxyz")
	dnsResolver = mockdns.Resolver{
		Zones: map[string]mockdns.Zone{
			testNodeName + ".test.ch.": {
				A:    testNodeIpv4,
				AAAA: testNodeIpv6,
			},
		},
	}

	testingConfig := controller.Config{
		LeaderElection:         true,
		RegexStr:               `^[\w-]*\.test\.ch$`,
		MaxExpirationSeconds:   367 * 24 * 3600,
		AllowedDNSNames:        3,
		K8sConfig:              cfg,
		IgnoreNonSystemNodeCsr: true,
		DNSResolver:            &dnsResolver,
		IPPrefixesStr:          "192.168.0.0/16,fc00::/7",
	}

	csrCtrl, mgr, errorCode := cmd.CreateControllerManager(&testingConfig, controller.InitLogger(&testingConfig))
	csrController = csrCtrl
	if errorCode != 0 {
		log.Fatalf("unable to create controller-runtime manager. Error:\n%v", errorCode)
	}

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
