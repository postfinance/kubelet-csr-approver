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
	"net"
	"testing"

	"github.com/foxcpp/go-mockdns"
	"github.com/stretchr/testify/require"
	"github.com/tj/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestValidCsrApproved(t *testing.T) {
	csrParams := CsrParams{
		ipAddresses: testNodeIpAddresses,
		nodeName:    testNodeName,
		dnsName:     testNodeName + ".test.ch",
	}
	validCsr := createCsr(t, csrParams)

	_, nodeClientSet, _ := createControlPlaneUser(t, validCsr.Spec.Username, []string{"system:masters"})
	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &validCsr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(validCsr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.False(t, denied)
	assert.True(t, approved)
}

func TestWrongSignerCsr(t *testing.T) {
	csrParams := CsrParams{
		csrName:     "csr-wrong-signer",
		ipAddresses: testNodeIpAddresses,
		nodeName:    testNodeName,
		dnsName:     testNodeName + ".test.ch",
	}
	csr := createCsr(t, csrParams)
	csr.Spec.SignerName = "example.com/not-kubelet-serving"

	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})
	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.False(t, denied)
	assert.False(t, approved)
}

func TestNonMatchingCommonNameUsername(t *testing.T) {
	csrParams := CsrParams{
		csrName:     "csr-non-matching",
		commonName:  "funny-common-name",
		ipAddresses: testNodeIpAddresses,
		nodeName:    testNodeName,
		dnsName:     testNodeName + ".test.ch",
	}
	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.True(t, denied)
	assert.False(t, approved)
}

func TestInvalidDNSName(t *testing.T) {
	csrParams := CsrParams{
		csrName:  "csr-invalid-dnsName",
		nodeName: testNodeName,
		dnsName:  "fishing.google.com",
	}
	dnsResolver.Zones[csrParams.dnsName+"."] = mockdns.Zone{
		A: []string{"1.2.3.14"},
	} // we mock the dns zone of this test, as we really only want the invalid dns name to make it fail
	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.True(t, denied)
	assert.False(t, approved)
}

func TestInvalidRegexName(t *testing.T) {
	csrParams := CsrParams{
		csrName:  "csr-invalid-regexName",
		nodeName: testNodeName,
		dnsName:  testNodeName + ".phishingTemptative.ch",
	}
	dnsResolver.Zones[csrParams.dnsName+"."] = mockdns.Zone{
		A: []string{"1.2.3.14"},
	} // we mock the dns zone of this test, as we really only want the invalid dns name to make it fail
	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.True(t, denied)
	assert.False(t, approved)
}
func TestUnresolvedDNSName(t *testing.T) {
	csrParams := CsrParams{
		csrName:  "csr-unresolved-dnsName",
		nodeName: "unresolved",
		dnsName:  "unresolved.test.ch",
	}
	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.True(t, denied)
	assert.False(t, approved)
}

func TestMismatchedResolvedIpsSANIps(t *testing.T) {
	csrParams := CsrParams{
		csrName:     "mismatched-san-ip-resolved-dns-ip",
		nodeName:    testNodeName,
		dnsName:     testNodeName + ".test.ch",
		ipAddresses: []net.IP{{9, 9, 9, 9}},
	}
	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.True(t, denied)
	assert.False(t, approved)

}

func TestExpirationSecondsTooLarge(t *testing.T) {
	csrParams := CsrParams{
		csrName:           "expiration-seconds",
		expirationSeconds: 368 * 24 * 3600, // one day more than the maximum of 367
		nodeName:          testNodeName,
		dnsName:           testNodeName + ".test.ch",
	}
	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.True(t, denied)
	assert.False(t, approved)
}

func TestBypassDNSResolution(t *testing.T) {
	csrParams := CsrParams{
		csrName:  "dns-bypass",
		nodeName: testNodeName,
		dnsName:  testNodeName + "-unresolved.test.ch",
	}
	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	csrController.BypassDNSResolution = true
	defer func() { csrController.BypassDNSResolution = false }()

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(
		testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.True(t, approved)
	assert.False(t, denied)
}

func TestIPv4NotWhitelisted(t *testing.T) {
	csrParams := CsrParams{
		csrName:     "ipv4-non-whitelisted",
		nodeName:    testNodeName,
		ipAddresses: []net.IP{{9, 9, 9, 9}},
		dnsName:     testNodeName + "-v4-non-whitelisted.test.ch",
	}
	dnsResolver.Zones[csrParams.dnsName+"."] = mockdns.Zone{
		A: []string{"9.9.9.9"},
	}

	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(
		testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.False(t, approved)
	assert.True(t, denied)
}

func TestIPv6NotWhitelisted(t *testing.T) {
	csrParams := CsrParams{
		csrName:     "ipv6-non-whitelisted",
		nodeName:    testNodeName,
		ipAddresses: []net.IP{{0x20, 0x01, 0xc0, 0xfe, 0xbe, 0xef, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x14}},
		dnsName:     testNodeName + "-v6-non-whitelisted.test.ch",
	}
	dnsResolver.Zones[csrParams.dnsName+"."] = mockdns.Zone{
		AAAA: []string{"2001:c0fe:beef::14"},
	}

	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(
		testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.False(t, approved)
	assert.True(t, denied)
}

func TestIPv6WithoutDNSNotWhitelisted(t *testing.T) {
	csrParams := CsrParams{
		csrName:     "ipv6-noDNS-non-whitelisted",
		nodeName:    testNodeName,
		ipAddresses: []net.IP{{0x20, 0x01, 0xc0, 0xfe, 0xbe, 0xef, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x14}},
	}
	dnsResolver.Zones[csrParams.dnsName+"."] = mockdns.Zone{
		AAAA: []string{"2001:c0fe:beef::14"},
	}

	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err := nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(
		testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.False(t, approved)
	assert.True(t, denied)
}
