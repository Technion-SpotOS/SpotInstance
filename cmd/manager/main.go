// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	controllers "github.com/Technion-SpotOS/SpotInstance/pkg/controller"
	"github.com/go-logr/logr"
	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	metricsHost                                  = "0.0.0.0"
	metricsPort                            int32 = 8384
	environmentVariableControllerNamespace       = "POD_NAMESPACE"
	environmentVariableWatchNamespace            = "WATCH_NAMESPACE"
	leaderElectionLockName                       = "spot-instance-controller-lock"
)

func printVersion(log logr.Logger) {
	log.Info(fmt.Sprintf("Go Version: %s", runtime.Version()))
	log.Info(fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH))
}

// function to handle defers with exit, see https://stackoverflow.com/a/27629493/553720.
func doMain() int {
	ctrl.SetLogger(zap.New())
	log := ctrl.Log.WithName("cmd")

	printVersion(log)

	leaderElectionNamespace, found := os.LookupEnv(environmentVariableControllerNamespace)
	if !found {
		log.Error(nil, "Failed to get leader-election error")
		return 1
	}

	namespace, found := os.LookupEnv(environmentVariableWatchNamespace)
	if !found {
		log.Error(nil, "Failed to get watch namespace")
		return 1
	}

	mgr, err := createManager(leaderElectionNamespace, namespace, metricsHost, metricsPort)
	if err != nil {
		log.Error(err, "Failed to create manager")
		return 1
	}

	log.Info("Starting the Cmd.")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		log.Error(err, "Manager exited non-zero")
		return 1
	}

	return 0
}

func createManager(leaderElectionNamespace, namespace, metricsHost string, metricsPort int32) (ctrl.Manager, error) {
	options := ctrl.Options{
		Namespace:               namespace,
		MetricsBindAddress:      fmt.Sprintf("%s:%d", metricsHost, metricsPort),
		LeaderElection:          true,
		LeaderElectionNamespace: leaderElectionNamespace,
		LeaderElectionID:        "hub-of-hubs-spec-sync-lock",
	}

	// Add support for MultiNamespace set in WATCH_NAMESPACE (e.g ns1,ns2)
	// Note that this is not intended to be used for excluding namespaces, this is better done via a Predicate
	// Also note that you may face performance issues when using this with a high number of namespaces.
	// More Info: https://godoc.org/github.com/kubernetes-sigs/controller-runtime/pkg/cache#MultiNamespacedCacheBuilder
	if strings.Contains(namespace, ",") {
		options.Namespace = ""
		options.NewCache = cache.MultiNamespacedCacheBuilder(strings.Split(namespace, ","))
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), options)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new manager: %w", err)
	}

	if err := controllers.AddToScheme(mgr.GetScheme()); err != nil {
		return nil, fmt.Errorf("failed to add schemes: %w", err)
	}

	if err := controllers.SetupWithManager(mgr); err != nil {
		return nil, fmt.Errorf("failed to add controllers: %w", err)
	}

	return mgr, nil
}

func main() {
	os.Exit(doMain())
}
