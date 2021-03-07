package main

import (
	controller2 "image-clone-controller/pkg/controller"
	appsv1 "k8s.io/api/apps/v1"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/source"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)



func init() {
	log.SetLogger(zap.New())
}

func main(){
	entryLog := log.Log.WithName("entrypoint")

	// Setup a Manager
	entryLog.Info("setting up manager")
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		entryLog.Error(err, "unable to set up overall controller manager")
		os.Exit(1)
	}

	// Setup a new controller to reconcile Deployments
	entryLog.Info("Setting up deployment-controller")
	c, err := controller.New("deployment-controller", mgr, controller.Options{
		Reconciler: &controller2.ReconcileDeployment{Client: mgr.GetClient()},
	})
	if err != nil {
		entryLog.Error(err, "unable to set up deployment-controller")
		os.Exit(1)
	}

	if err := c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForObject{}); err != nil {
		entryLog.Error(err, "unable to watch Deployments")
		os.Exit(1)
	}

	// Setup a new controller to reconcile Daemonset
	d, err := controller.New("daemonset-controller", mgr, controller.Options{
		Reconciler: &controller2.ReconcileDaemonSet{Client: mgr.GetClient()},
	})
	if err != nil {
		entryLog.Error(err, "unable to set up daemonset-controller")
		os.Exit(1)
	}

	if err := d.Watch(&source.Kind{Type: &appsv1.DaemonSet{}}, &handler.EnqueueRequestForObject{}); err != nil {
		entryLog.Error(err, "unable to watch Daemonset")
		os.Exit(1)
	}



	entryLog.Info("starting manager")
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		entryLog.Error(err, "unable to run manager")
		os.Exit(1)
	}


}