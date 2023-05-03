/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	golemv1alpha1 "ithub.com/Technion-SpotOS/SpotInstance/api/v1alpha1"
)

// SpotInstanceReconciler reconciles a SpotInstance object
type SpotInstanceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=golem.spot-os.io,resources=spotinstances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=golem.spot-os.io,resources=spotinstances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=golem.spot-os.io,resources=spotinstances/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *SpotInstanceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	spotInstance := &golemv1alpha1.SpotInstance{}
	err := r.Get(ctx, req.NamespacedName, spotInstance)

	// TODO: create controllers for:
	// 1. Creating a spot instance
	// 2. Deploying the instance as a node
	// 3. Updating CR Status with the node name

	return ctrl.Result{}, nil
}

// setupSpotInstanceController sets up the controller with the Manager.
func (r *SpotInstanceReconciler) setupSpotInstanceController(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&golemv1alpha1.SpotInstance{}).
		Complete(r)
}