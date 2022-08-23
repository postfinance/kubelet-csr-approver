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
//
//nolint:gocyclo
func (r *CertificateSigningRequestReconciler) DNSCheck(ctx context.Context, csr *certificatesv1.CertificateSigningRequest, x509cr *x509.CertificateRequest) (valid bool, reason string, err error) {
	if valid = (len(x509cr.DNSNames) <= r.AllowedDNSNames); !valid {
		reason = "The x509 Cert Request contains more DNS names than allowed through the config flag"
		return
	}

	// no DNS name to check, the DNS check is approved
	if len(x509cr.DNSNames) == 0 {
		valid = true
		return valid, reason, nil
	}

	// bypassing DNS reslution - DNS check is approved
	if r.BypassDNSResolution {
		valid = true
		return valid, reason, nil
	}

	dnsCtx, dnsCtxCancel := context.WithDeadline(ctx, time.Now().Add(time.Second)) // 1 second timeout for the dns request
	defer dnsCtxCancel()

	var allResolvedAddrs []string

	for _, sanDNSName := range x509cr.DNSNames {
		hostname := strings.TrimPrefix(csr.Spec.Username, "system:node:")

		if valid = strings.HasPrefix(sanDNSName, hostname); !valid && !r.BypassHostnameCheck {
			reason = "The SAN DNS Name in the x509 CSR is not prefixed by the node name (hostname)"
			return
		}

		if valid = r.ProviderRegexp(sanDNSName); !valid {
			reason = "The SAN DNS name in the x509 CR is not allowed by the Cloud provider regex"
			return
		}

		resolvedAddrs, err := r.DNSResolver.LookupHost(dnsCtx, sanDNSName)

		if err != nil || len(resolvedAddrs) == 0 {
			return false, "The SAN DNS Name could not be resolved, denying the CSR", nil
		}

		allResolvedAddrs = append(allResolvedAddrs, resolvedAddrs...)
	}

	var setBuilder netaddr.IPSetBuilder

	for _, a := range allResolvedAddrs {
		ipaddr, err := netaddr.ParseIP(a)
		if err != nil {
			return false, fmt.Sprintf("Error while parsing resolved IP address %s, denying the CSR", ipaddr), nil
		}

		setBuilder.Add(ipaddr)

		if !r.ProviderIPSet.Contains(ipaddr) {
			return false, fmt.Sprintf("One of the resolved IP addresses, %s,"+
				"isn't part of the provider-specified set of whitelisted IP. denying the certificate",
				ipaddr), nil
		}
	}

	resolvedIPSet, _ := setBuilder.IPSet()

	sanIPAddrs := x509cr.IPAddresses
	for _, ip := range sanIPAddrs {
		ipa, ok := netaddr.FromStdIP(ip)
		if !ok {
			return false, fmt.Sprintf("Error while parsing x509 CR IP address %s, denying the CSR", ip), nil
		}

		if !resolvedIPSet.Contains(ipa) {
			return false, fmt.Sprintf("One of the SAN IP addresses, %s, "+
				"is not contained in the set of resolved IP addresses, denying the CSR.", ipa), nil
		}
	}

	return valid, reason, nil
}

// WhitelistedIPCheck verifies that the x509cr SAN IP Addresses are contained in the
// set of ProviderSpecified IP addresses
func (r *CertificateSigningRequestReconciler) WhitelistedIPCheck(csr *certificatesv1.CertificateSigningRequest, x509cr *x509.CertificateRequest) (valid bool, reason string, err error) {
	sanIPAddrs := x509cr.IPAddresses
	for _, ip := range sanIPAddrs {
		ipa, ok := netaddr.FromStdIP(ip)
		if !ok {
			return false, fmt.Sprintf("Error while parsing x509 CR IP address %s, denying the CSR", ip), nil
		}

		if !r.ProviderIPSet.Contains(ipa) {
			return false,
				fmt.Sprintf(
					"One of the SAN IP addresses, %s, is not part"+
						"of the allowed IP Prefixes/Subnets, denying the CSR.", ipa),
				nil
		}
	}

	return true, reason, nil
}
