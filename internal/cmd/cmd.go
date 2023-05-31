// Package cmd - command line initialization
package cmd

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"net/netip"

	"go4.org/netipx"

	clientset "k8s.io/client-go/kubernetes"

	"github.com/peterbourgon/ff/v3"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"

	"github.com/go-logr/logr"
	"github.com/postfinance/kubelet-csr-approver/internal/controller"
	ctrlconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

//nolint:gochecknoglobals //this vars are set on build by goreleaser
var (
	commit = "12345678"
	ref    = "refs/refname"
)

// Run will start the controller with the default settings
func Run() int {
	config := prepareCmdlineConfig()
	logger := controller.InitLogger(config)
	_, mgr, errorCode := CreateControllerManager(config, logger)

	if errorCode != 0 {
		return errorCode
	}

	z := mgr.GetLogger()
	z.V(1).Info("starting controller-runtime manager")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		z.Error(err, "problem running manager")
		return 1
	}

	return 0
}

// CreateControllerManager permits creation/customization of the controller-manager
func CreateControllerManager(config *controller.Config, logger logr.Logger) (
	csrController *controller.CertificateSigningRequestReconciler,
	mgr ctrl.Manager,
	code int,
) {
	logger.V(0).Info("Kubelet-CSR-Approver controller starting.", "commit", commit, "ref", ref)

	if config.K8sConfig == nil { // when testing, this variable is already set
		config.K8sConfig = ctrl.GetConfigOrDie()
	}

	if config.RegexStr == "" {
		logger.V(-5).Info("the provider-spefic regex must be specified, exiting")
		return nil, nil, 10
	}

	csrController = &controller.CertificateSigningRequestReconciler{
		Config: *config,
	}
	csrController.ProviderRegexp = regexp.MustCompile(config.RegexStr).MatchString

	// IP Prefixes parsing and IPSet construction

	var setBuilder netipx.IPSetBuilder

	for _, ipPrefix := range strings.Split(config.IPPrefixesStr, ",") {
		ipPref, err := netip.ParsePrefix(ipPrefix)
		if err != nil {
			logger.V(-5).Info(fmt.Sprintf("Unable to parse IP prefix: %s, exiting", ipPrefix))

			return nil, nil, 10
		}

		setBuilder.AddPrefix(ipPref)
	}

	var err error
	csrController.ProviderIPSet, err = setBuilder.IPSet()

	if err != nil {
		logger.V(-5).Info("Unable to build the Set of valid IP addresses, exiting")

		return nil, nil, 10
	}

	ctrl.SetLogger(logger)

	mgrOptions := ctrl.Options{
		MetricsBindAddress:     config.MetricsAddr,
		HealthProbeBindAddress: config.ProbeAddr,
		LeaderElection:         config.LeaderElection,
		LeaderElectionID:       "kubelet-csr-approver",
	}

	// we need to test whether we are running in-cluster or not. if we are not, we will
	// set a default LeaderElectionNamespace to permit our tests to correctly run.
	// the detection is done as defined in controller-runtime/pkg/leaderelection:
	// https://github.com/kubernetes-sigs/controller-runtime/blob/v0.15.0/pkg/leaderelection/leader_election.go
	// TODO: with https://github.com/golang/go/issues/52600 and Go 1.21, we will be able to use
	// testing.Testing() function to know whether running in test mode or not. adapt the next block accordingly
	const inClusterNamespacePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	if _, err := os.Stat(inClusterNamespacePath); os.IsNotExist(err) {
		mgrOptions.LeaderElectionNamespace = "kube-system"
	}

	mgr, err = ctrl.NewManager(config.K8sConfig, mgrOptions)

	if err != nil {
		logger.Error(err, "unable to start manager")

		return nil, nil, 10
	}

	csrController.ClientSet = clientset.NewForConfigOrDie(config.K8sConfig)
	csrController.Client = mgr.GetClient()
	csrController.Scheme = mgr.GetScheme()

	if err = csrController.SetupWithManager(mgr); err != nil {
		logger.Error(err, "unable to create controller", "controller", "CertificateSigningRequest")

		return nil, nil, 10
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		logger.Error(err, "unable to set up health check")

		return nil, nil, 10
	}

	return csrController, mgr, 0
}

func prepareCmdlineConfig() *controller.Config {
	fs := flag.NewFlagSet("kubelet-csr-approver", flag.ExitOnError)
	ctrlconfig.RegisterFlags(fs)

	var (
		logLevel               = fs.Int("level", 0, "level ranges from -5 (Fatal) to 10 (Verbose)")
		metricsAddr            = fs.String("metrics-bind-address", ":8080", "address the metric endpoint binds to.")
		probeAddr              = fs.String("health-probe-bind-address", ":8081", "address the probe endpoint binds to.")
		leaderElection         = fs.Bool("leader-election", false, "set this parameter to true to enable leader election")
		regexStr               = fs.String("provider-regex", ".*", "provider-specified regex to validate CSR SAN names against. accepts everything unless specified")
		maxSec                 = fs.Int("max-expiration-sec", 367*24*3600, "maximum seconds a CSR can request a cerficate for. defaults to 367 days")
		bypassDNSResolution    = fs.Bool("bypass-dns-resolution", false, "set this parameter to true to bypass DNS resolution checks")
		bypassHostnameCheck    = fs.Bool("bypass-hostname-check", false, "set this parameter to true to ignore mismatching DNS name and hostname")
		ignoreNonSystemNodeCsr = fs.Bool("ignore-non-system-node", false, "set this parameter to true to ignore CSR for subjects different than system:node")
		allowedDNSNames        = fs.Int("allowed-dns-names", 1, "number of DNS SAN names allowed in a certificate request. defaults to 1")
		ipPrefixesStr          = fs.String("provider-ip-prefixes", "0.0.0.0/0,::/0",
			`provider-specified, comma separated ip prefixes that CSR IP addresses shall fall into.
			left unspecified, all IPv4/v6 are allowed. example prefix definition:
			192.168.0.0/16,fc00/7`,
		)
	)

	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVars())
	if err != nil {
		fmt.Printf("unable to parse args/envs, exiting. error message: %v", err)

		os.Exit(2)
	}

	if *maxSec < 0 || *maxSec > 367*24*3600 {
		fmt.Print("the maximum expiration seconds cannot be lower than 0 nor greater than 367 days")

		os.Exit(2)
	}

	if *allowedDNSNames < 1 || *allowedDNSNames > 1000 {
		fmt.Print("the number of allowed DNS names must be at least 1 and no more than 1000")
	}

	config := controller.Config{
		LogLevel:               *logLevel,
		MetricsAddr:            *metricsAddr,
		ProbeAddr:              *probeAddr,
		LeaderElection:         *leaderElection,
		RegexStr:               *regexStr,
		IPPrefixesStr:          *ipPrefixesStr,
		BypassDNSResolution:    *bypassDNSResolution,
		BypassHostnameCheck:    *bypassHostnameCheck,
		IgnoreNonSystemNodeCsr: *ignoreNonSystemNodeCsr,
		MaxExpirationSeconds:   int32(*maxSec),
		AllowedDNSNames:        *allowedDNSNames,
	}

	config.DNSResolver = net.DefaultResolver

	return &config
}
