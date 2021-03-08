package controller

import (
	"context"
	"fmt"
	"image-clone-controller/pkg/utility"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const controllerNamespace = "image-clone-controller"

type ReconcileDeployment struct {
	Client client.Client
}


var _ reconcile.Reconciler = &ReconcileDeployment{}

func (r *ReconcileDeployment) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the Deployment
	rs := &appsv1.Deployment{}
	err := r.Client.Get(ctx, request.NamespacedName, rs)
	if errors.IsNotFound(err) {
		log.Error(nil, "Could not find Deployments")
		return reconcile.Result{}, nil
	}
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not fetch Deployments: %+v", err)
	}

	if rs.Namespace != "kube-system" && rs.Namespace != controllerNamespace && len(rs.Spec.Template.Spec.ImagePullSecrets) == 0 {
		containers := &rs.Spec.Template.Spec.Containers
		err = utility.ModifyImage(containers)
		if err != nil {
			return reconcile.Result{}, err
		}
		err = r.Client.Update(ctx, rs)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

type ReconcileDaemonSet struct {
	Client client.Client
}

var _ reconcile.Reconciler = &ReconcileDaemonSet{}

func (r *ReconcileDaemonSet) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log := log.FromContext(ctx)
	// Fetch the DaemonSet
	rs := &appsv1.DaemonSet{}
	err := r.Client.Get(ctx, request.NamespacedName, rs)

	if errors.IsNotFound(err) {
		log.Error(nil, "Could not find DaemonSet")
		return reconcile.Result{}, nil
	}
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("could not fetch DaemonSet: %+v", err)
	}

	if rs.Namespace != "kube-system" && len(rs.Spec.Template.Spec.ImagePullSecrets) == 0 {
		containers := &rs.Spec.Template.Spec.Containers

		err = utility.ModifyImage(containers)
		if err != nil {
			return reconcile.Result{}, err
		}

		err = r.Client.Update(ctx, rs)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}
