// Package cmd - command line initialization
package cmd

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"

	"go.uber.org/zap/zapcore"
	clientset "k8s.io/client-go/kubernetes"

	"github.com/go-logr/zapr"
	"github.com/peterbourgon/ff/v3"
	"github.com/postfinance/flash"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"

	"github.com/postfinance/kubelet-csr-approver/internal/controller"
)

// ProviderRegexEnvvarName holds the name of the env variable containing the provider-spefic regex
const ProviderRegexEnvvarName string = "PROVIDER_REGEX"

// MaxExpirationSecEnvVarName holds the name of the env variable defining the maximum seconds a CSR can request
const MaxExpirationSecEnvVarName string = "MAX_EXPIRATION_SEC"

//nolint:gochecknoglobals //this vars are set on build by goreleaser
var (
	commit = "12345678"
	ref    = "refs/refname"
)

// Run encapsulates all settings related to kubelet-csr-approver
func Run() int {
	flashLogger := flash.New()

	fs := flag.NewFlagSet("kubelet-csr-approver", flag.ExitOnError)

	var (
		logLevel    = fs.Int("level", 0, "level ranges from -5 (Fatal) to 10 (Verbose)")
		metricsAddr = fs.String("metrics-bind-address", ":8080", "address the metric endpoint binds to.")
		probeAddr   = fs.String("health-probe-bind-address", ":8081", "address the probe endpoint binds to.")
		regexStr    = fs.String("provider-regex", "", "provider-specified regex to validate CSR SAN names against")
		maxSec      = fs.Int("max-expiration-sec", 367*24*3600, "maximum seconds a CSR can request a cerficate for. defaults to 367 days")
	)

	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())
	if err != nil {
		fmt.Printf("unable to parse args/envs, exiting. error message: %v", err)
		return 2
	}

	// logger initialization
	if *logLevel < -5 || *logLevel > 10 {
		flashLogger.Fatal(fmt.Errorf("log level should be between -5 and 10 (included)"))
	}

	*logLevel *= -1 // we inverse the level for the logging behavior between zap and logr.Logger to match
	flashLogger.SetLevel(zapcore.Level(*logLevel))
	z := zapr.NewLogger(flashLogger.Desugar())

	z.V(0).Info("Kubelet-CSR-Approver controller starting.", "commit", commit, "ref", ref)

	if *regexStr == "" {
		z.V(-5).Info("the provider-spefic regex must be specified, exiting")
		return 10
	}

	providerRegexp := regexp.MustCompile(*regexStr)

	if *maxSec < 0 || *maxSec > 367*24*3600 {
		err := fmt.Errorf("the maximum expiration seconds env variable cannot be lower than 0 nor greater than 367 days")
		z.Error(err, "reduce the maxExpirationSec value")

		return 10
	}

	ctrl.SetLogger(z)
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		MetricsBindAddress:     *metricsAddr,
		HealthProbeBindAddress: *probeAddr,
	})

	if err != nil {
		z.Error(err, "unable to start manager")
		return 1
	}

	csrController := controller.CertificateSigningRequestReconciler{
		ClientSet:            clientset.NewForConfigOrDie(mgr.GetConfig()),
		Client:               mgr.GetClient(),
		Scheme:               mgr.GetScheme(),
		ProviderRegexp:       providerRegexp.MatchString,
		MaxExpirationSeconds: int32(*maxSec),
		Resolver:             net.DefaultResolver,
	}

	if err = csrController.SetupWithManager(mgr); err != nil {
		z.Error(err, "unable to create controller", "controller", "CertificateSigningRequest")
		return 1
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		z.Error(err, "unable to set up health check")
		return 1
	}

	z.V(1).Info("starting controller-runtime manager")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		z.Error(err, "problem running manager")
		return 1
	}

	return 0
}
