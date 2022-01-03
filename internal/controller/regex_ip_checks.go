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

// DNSCheck is a function checking that the DNS name:
// complies with the provider-specific regex
// is resolvable (this check can be opted out with a parameter)
func (r *CertificateSigningRequestReconciler) DNSCheck(ctx context.Context, csr *certificatesv1.CertificateSigningRequest, x509cr *x509.CertificateRequest) (valid bool, reason string, err error) {
	if valid = (len(x509cr.DNSNames) <= 1); !valid {
		reason = "The x509 Cert Request contains more than 1 DNS name"
		return
	}

	// no DNS name to check, the DNS check is approved
	if len(x509cr.DNSNames) == 0 {
		valid = true
		return valid, reason, nil
	}

	sanDNSName := x509cr.DNSNames[0]
	hostname := strings.TrimPrefix(csr.Spec.Username, "system:node:")

	if valid = strings.HasPrefix(sanDNSName, hostname); !valid {
		reason = "The SAN DNS Name in the x509 CSR is not prefixed by the node name (hostname)"
		return
	}

	if valid = r.ProviderRegexp(sanDNSName); !valid {
		reason = "The SAN DNS name in the x509 CR is not allowed by the Cloud provider regex"
		return
	}

	// bypassing DNS reslution - DNS check is approved
	if r.BypassDNSResolution {
		valid = true
		return valid, reason, nil
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

	resolvedIPSet, _ := setBuilder.IPSet()

	sanIPAddrs := x509cr.IPAddresses
	for _, ip := range sanIPAddrs {
		ipa, ok := netaddr.FromStdIP(ip)
		if !ok {
			return false, fmt.Sprintf("Error while parsing x509 CR IP address %s, denying the CSR", ip), nil
		}

		if !resolvedIPSet.Contains(ipa) {
			return false, fmt.Sprintf("One of the SAN IP addresses, %s, is not contained in the set of resolved IP addresses, denying the CSR.", ipa), nil
		}
	}

	return valid, reason, nil
}
