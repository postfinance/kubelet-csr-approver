package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"

	"go.uber.org/zap/zapcore"
	clientset "k8s.io/client-go/kubernetes"

	_ "k8s.io/client-go/plugin/pkg/client/auth" //TODO: remove when used in-cluster

	"github.com/go-logr/zapr"
	"github.com/postfinance/flash"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"

	"github.com/postfinance/kubelet-csr-approver/controller"
)

// ProviderRegexEnvvarName holds the name of the env variable containing the provider-spefic regex
const ProviderRegexEnvvarName string = "PROVIDER_REGEX"

//nolint:gochecknoglobals //this vars are set on build by goreleaser
var (
	commit = "12345678"
)

func main() {
	flashLogger := flash.New()

	var metricsAddr, probeAddr string

	var logLevel int

	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.IntVar(&logLevel, "level", 0, "level ranges from -5 (Fatal) to 10 (Verbose)")
	flag.Parse()

	if logLevel < -5 || logLevel > 10 {
		flashLogger.Fatal(fmt.Errorf("log level should be between -5 and 10 (included)"))
	}

	logLevel *= -1 // we inverse the level for the logging behavior between zap and logr.Logger to match
	flashLogger.SetLevel(zapcore.Level(logLevel))
	z := zapr.NewLogger(flashLogger.Desugar())

	z.V(0).Info("Kubelet-CSR-Approver controller starting.", "tag+commit", commit)

	var regexEnvVar string
	if regexEnvVar = os.Getenv(ProviderRegexEnvvarName); regexEnvVar == "" {
		err := fmt.Errorf("the provider-spefic regex must be specified in the %s env variable", ProviderRegexEnvvarName)
		z.Error(err, ProviderRegexEnvvarName+" not set")
		os.Exit(1)
	}

	providerRegexp := regexp.MustCompile(regexEnvVar)

	ctrl.SetLogger(z)
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		MetricsBindAddress:     metricsAddr,
		HealthProbeBindAddress: probeAddr,
	})

	if err != nil {
		z.Error(err, "unable to start manager")
		os.Exit(1)
	}

	csrController := controller.CertificateSigningRequestReconciler{
		ClientSet:      clientset.NewForConfigOrDie(mgr.GetConfig()),
		Client:         mgr.GetClient(),
		Scheme:         mgr.GetScheme(),
		ProviderRegexp: providerRegexp.MatchString,
		Resolver:       net.DefaultResolver,
	}

	if err = csrController.SetupWithManager(mgr); err != nil {
		z.Error(err, "unable to create controller", "controller", "CertificateSigningRequest")
		os.Exit(1)
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		z.Error(err, "unable to set up health check")
		os.Exit(1)
	}

	z.V(1).Info("starting controller-runtime manager")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		z.Error(err, "problem running manager")
		os.Exit(1)
	}
}
