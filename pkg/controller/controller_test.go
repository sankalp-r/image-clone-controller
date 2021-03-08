package controller

import (
	"context"
	"image-clone-controller/pkg/utility"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"testing"
)

func TestReconcileDeployment_Reconcile(t *testing.T) {
	image := "image:latest"
	dep := &appsv1.Deployment{
		ObjectMeta: v12.ObjectMeta{
			Name:      "dep",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{Containers: []v1.Container{{Image: "docker.io/repo/" + image}}},
			},
		},
	}
	obj := []runtime.Object{dep}

	cl := fake.NewClientBuilder().WithRuntimeObjects(obj...).Build()

	r := ReconcileDeployment{cl}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      "dep",
			Namespace: "default",
		},
	}

	old := utility.CacheFunc
	defer func() { utility.CacheFunc = old }()

	utility.CacheFunc = func(src string, dst string) error {
		return nil
	}

	temp := utility.NewRegistry
	defer func() { utility.NewRegistry = temp }()
	utility.NewRegistry = "quay.io/repo"

	_, err := r.Reconcile(context.TODO(), req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}

	dep1 := &appsv1.Deployment{}
	err = r.Client.Get(context.TODO(), req.NamespacedName, dep1)
	if err != nil {
		t.Fatalf("get deployment: (%v)", err)
	}
	actualImage := dep1.Spec.Template.Spec.Containers[0].Image
	if actualImage != utility.NewRegistry+"/"+image {
		t.Errorf("Expected Image: %s is not the acuaul Image: %s", utility.NewRegistry, actualImage)
	}
}

func TestReconcileDaemonSet_Reconcile(t *testing.T) {
	image := "image:latest"
	daemon := &appsv1.DaemonSet{
		ObjectMeta: v12.ObjectMeta{
			Name:      "daemon",
			Namespace: "default",
		},
		Spec: appsv1.DaemonSetSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{Containers: []v1.Container{{Image: "docker.io/repo/" + image}}},
			},
		},
	}
	obj := []runtime.Object{daemon}

	cl := fake.NewClientBuilder().WithRuntimeObjects(obj...).Build()

	r := ReconcileDaemonSet{cl}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      "daemon",
			Namespace: "default",
		},
	}

	old := utility.CacheFunc
	defer func() { utility.CacheFunc = old }()

	utility.CacheFunc = func(src string, dst string) error {
		return nil
	}

	temp := utility.NewRegistry
	defer func() { utility.NewRegistry = temp }()
	utility.NewRegistry = "quay.io/repo"

	_, err := r.Reconcile(context.TODO(), req)
	if err != nil {
		t.Fatalf("reconcile: (%v)", err)
	}

	daemon1 := &appsv1.DaemonSet{}
	err = r.Client.Get(context.TODO(), req.NamespacedName, daemon1)
	if err != nil {
		t.Fatalf("get deployment: (%v)", err)
	}
	actualImage := daemon1.Spec.Template.Spec.Containers[0].Image
	if actualImage != utility.NewRegistry+"/"+image {
		t.Errorf("Expected Image: %s is not the acuaul Image: %s", utility.NewRegistry, actualImage)
	}
}
