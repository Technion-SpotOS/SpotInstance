// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package controllers

import (
	"fmt"

	"github.com/Technion-SpotOS/SpotInstance/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// AddToScheme adds all the resources to be processed to the Scheme.
func AddToScheme(scheme *runtime.Scheme) error {
	for _, schemeBuilder := range getSchemeBuilders() {
		if err := schemeBuilder.AddToScheme(scheme); err != nil {
			return fmt.Errorf("failed to add scheme: %w", err)
		}
	}

	// install golem v1alpha1 scheme
	if err := v1alpha1.AddToScheme(scheme); err != nil {
		return fmt.Errorf("failed to install scheme: %w", err)
	}

	return nil
}

func getSchemeBuilders() []*scheme.Builder {
	return []*scheme.Builder{
		v1alpha1.SchemeBuilder,
	}
}

// AddControllers adds all the controllers to the Manager.
func SetupWithManager(mgr ctrl.Manager) error {
	setupWithManagerFuncs := []func(ctrl.Manager) error{
		setupSpotInstanceController,
	}

	for _, setupWithManagerFunc := range setupWithManagerFuncs {
		if err := setupWithManagerFunc(mgr); err != nil {
			return fmt.Errorf("failed to setup controller: %w", err)
		}
	}

	return nil
}
