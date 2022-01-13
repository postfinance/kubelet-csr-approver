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
	"k8s.io/client-go/rest"

	"github.com/go-logr/zapr"
	"github.com/peterbourgon/ff/v3"
	"github.com/postfinance/flash"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"

	"github.com/postfinance/kubelet-csr-approver/internal/controller"
)

//nolint:gochecknoglobals //this vars are set on build by goreleaser
var (
	commit = "12345678"
	ref    = "refs/refname"
)

// Config stores all parameters needed to configure a controller-manager
type Config struct {
	logLevel            int
	metricsAddr         string
	probeAddr           string
	RegexStr            string
	MaxSec              int
	K8sConfig           *rest.Config
	DNSResolver         controller.HostResolver
	BypassDNSResolution bool
}

// Run will start the controller with the default settings
func Run() int {
	config := prepareCmdlineConfig()
	_, mgr, errorCode := CreateControllerManager(config)

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
func CreateControllerManager(config *Config) (
	csrController *controller.CertificateSigningRequestReconciler,
	mgr ctrl.Manager,
	code int,
) {
	// logger initialization
	flashLogger := flash.New()
	if config.logLevel < -5 || config.logLevel > 10 {
		flashLogger.Fatal(fmt.Errorf("log level should be between -5 and 10 (included)"))
	}

	config.logLevel *= -1 // we inverse the level for the logging behavior between zap and logr.Logger to match
	flashLogger.SetLevel(zapcore.Level(config.logLevel))
	z := zapr.NewLogger(flashLogger.Desugar())

	z.V(0).Info("Kubelet-CSR-Approver controller starting.", "commit", commit, "ref", ref)

	if config.RegexStr == "" {
		z.V(-5).Info("the provider-spefic regex must be specified, exiting")

		return nil, nil, 10
	}

	providerRegexp := regexp.MustCompile(config.RegexStr)

	if config.MaxSec < 0 || config.MaxSec > 367*24*3600 {
		err := fmt.Errorf("the maximum expiration seconds env variable cannot be lower than 0 nor greater than 367 days")
		z.Error(err, "reduce the maxExpirationSec value")

		return nil, nil, 10
	}

	ctrl.SetLogger(z)
	mgr, err := ctrl.NewManager(config.K8sConfig, ctrl.Options{
		MetricsBindAddress:     config.metricsAddr,
		HealthProbeBindAddress: config.probeAddr,
	})

	if err != nil {
		z.Error(err, "unable to start manager")

		return nil, nil, 10
	}

	csrController = &controller.CertificateSigningRequestReconciler{
		ClientSet:            clientset.NewForConfigOrDie(config.K8sConfig),
		Client:               mgr.GetClient(),
		Scheme:               mgr.GetScheme(),
		ProviderRegexp:       providerRegexp.MatchString,
		MaxExpirationSeconds: int32(config.MaxSec),
		Resolver:             config.DNSResolver,
		BypassDNSResolution:  config.BypassDNSResolution,
	}

	if err = csrController.SetupWithManager(mgr); err != nil {
		z.Error(err, "unable to create controller", "controller", "CertificateSigningRequest")

		return nil, nil, 10
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		z.Error(err, "unable to set up health check")

		return nil, nil, 10
	}

	return csrController, mgr, 0
}

func prepareCmdlineConfig() *Config {
	fs := flag.NewFlagSet("kubelet-csr-approver", flag.ExitOnError)

	var (
		logLevel            = fs.Int("level", 0, "level ranges from -5 (Fatal) to 10 (Verbose)")
		metricsAddr         = fs.String("metrics-bind-address", ":8080", "address the metric endpoint binds to.")
		probeAddr           = fs.String("health-probe-bind-address", ":8081", "address the probe endpoint binds to.")
		regexStr            = fs.String("provider-regex", "", "provider-specified regex to validate CSR SAN names against")
		maxSec              = fs.Int("max-expiration-sec", 367*24*3600, "maximum seconds a CSR can request a cerficate for. defaults to 367 days")
		bypassDNSResolution = fs.Bool("bypass-dns-resolution", false, "set this parameter to true to bypass DNS resolution checks")
	)

	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())
	if err != nil {
		fmt.Printf("unable to parse args/envs, exiting. error message: %v", err)

		os.Exit(2)
	}

	config := Config{
		logLevel:            *logLevel,
		metricsAddr:         *metricsAddr,
		probeAddr:           *probeAddr,
		RegexStr:            *regexStr,
		BypassDNSResolution: *bypassDNSResolution,
		MaxSec:              *maxSec,
	}

	config.DNSResolver = net.DefaultResolver
	config.K8sConfig = ctrl.GetConfigOrDie()

	return &config
}
