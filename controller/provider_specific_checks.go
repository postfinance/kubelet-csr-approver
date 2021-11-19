package controller

import (
	"crypto/x509"

	certificatesv1 "k8s.io/api/certificates/v1"
)

// ProviderChecks is a function in which the Cloud Provider specifies a series of checks
// to run against the CSRs. The out-of-band identity checks of the CSRs should happen here
func ProviderChecks(csr *certificatesv1.CertificateSigningRequest, x509csr *x509.CertificateRequest) (valid bool, reason string) {
	return true, ""
}
