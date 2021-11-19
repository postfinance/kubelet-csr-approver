package controller

import (
	"context"
	"crypto/x509"
	"fmt"
	"strings"
	"time"

	"inet.af/netaddr"

	certificatesv1 "k8s.io/api/certificates/v1"
)

// RegexIPChecks is a function checking that the DNS Name and the IP address comply with the provider-specific regex/ranges
func (r *CertificateSigningRequestReconciler) RegexIPChecks(ctx context.Context, csr *certificatesv1.CertificateSigningRequest, x509cr *x509.CertificateRequest) (valid bool, reason string, err error) {
	if valid = (len(x509cr.DNSNames) <= 1); !valid {
		reason = "The x509 Cert Request contains more than 1 DNS name"
		return
	}

	if valid = (len(x509cr.DNSNames)+len(x509cr.IPAddresses) != 0); !valid {
		reason = "The x509 Cert Request SAN doesn't contain IP address nor DNS name"
		return
	}

	if len(x509cr.DNSNames) == 1 {
		sanDNSName := x509cr.DNSNames[0]

		if valid = r.ProviderRegexp(sanDNSName); !valid {
			reason = "The DNS name in the x509 CSR is not allowed by the Cloud provider regex"
			return
		}

		hostname := strings.TrimPrefix(csr.Spec.Username, "system:node:")
		if valid = strings.HasPrefix(sanDNSName, hostname); !valid {
			reason = "The SAN DNS Name in the x509 CSR is not prefixed by the node name (hostname)"
			return
		}

		dnsCtx, dnsCtxCancel := context.WithDeadline(ctx, time.Now().Add(time.Second)) // 1 second timeout for the dns request
		defer dnsCtxCancel()

		var resolvedAddrs []string
		resolvedAddrs, err = r.Resolver.LookupHost(dnsCtx, sanDNSName)

		if err != nil || len(resolvedAddrs) == 0 {
			return false, "The SAN DNS Name could not be resolved, denying the CSR", nil
		}

		var setBuilder netaddr.IPSetBuilder

		for _, a := range resolvedAddrs {
			ipa, err := netaddr.ParseIP(a)
			if err != nil {
				return false, fmt.Sprintf("Error while parsing resolved IP address %s, denying the CSR", ipa), nil
			}

			setBuilder.Add(ipa)
		}

		ipSet, _ := setBuilder.IPSet()

		sanIPAddrs := x509cr.IPAddresses
		for _, ip := range sanIPAddrs {
			ipa, ok := netaddr.FromStdIP(ip)
			if !ok {
				return false, fmt.Sprintf("Error while x509 CR IP address %s, denying the CSR", ip), nil
			}

			if !ipSet.Contains(ipa) {
				return false, fmt.Sprintf("One of the SAN IP addresses, %s, is not contained in the set of resolved IP addresses, denying the CSR.", ipa), nil
			}
		}
	}

	valid = true

	return valid, reason, nil
}
