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
	}
	validCsr := createCsr(t, csrParams)

	_, nodeClientSet, _ := createControlPlaneUser(t, validCsr.Spec.Username, []string{"system:masters"})
	_, err = nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &validCsr, metav1.CreateOptions{})
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
	}
	csr := createCsr(t, csrParams)
	csr.Spec.SignerName = "example.com/not-kubelet-serving"

	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})
	_, err = nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.False(t, denied)
	assert.False(t, approved)
}

func TestNonMatchingCommonNameUsername(t *testing.T) {
	csrParams := CsrParams{
		csrName:    "csr-non-matching",
		commonName: "funny-common-name",

		ipAddresses: testNodeIpAddresses, nodeName: testNodeName}
	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err = nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
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

	_, err = nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
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

	_, err = nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
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
		ipAddresses: []net.IP{{9, 9, 9, 9}},
	}
	csr := createCsr(t, csrParams)
	_, nodeClientSet, _ := createControlPlaneUser(t, csr.Spec.Username, []string{"system:masters"})

	_, err = nodeClientSet.CertificatesV1().CertificateSigningRequests().Create(testContext, &csr, metav1.CreateOptions{})
	require.Nil(t, err, "Could not create the CSR.")

	approved, denied, err := waitCsrApprovalStatus(csr.Name)
	require.Nil(t, err, "Could not retrieve the CSR to check its approval status")
	assert.True(t, denied)
	assert.False(t, approved)

}
