/*
Copyright 2022.

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
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	crdtryv1 "kubebuilder.test/crdtry/api/v1"
)

const StatusPending string = "PENDING"
const StatusRunning string = "RUNNING"

// LpxpodReconciler reconciles a Lpxpod object
type LpxpodReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=crdtry.kubebuilder.test,resources=lpxpods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=crdtry.kubebuilder.test,resources=lpxpods/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=crdtry.kubebuilder.test,resources=lpxpods/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Lpxpod object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *LpxpodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	reqLogger := log.Log.WithValues("namespace", req.Namespace, "lpxpod", req.Name)
	reqLogger.Info("==== Call Reconcile ====")
	instance := &crdtryv1.Lpxpod{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, instance)

	if err != nil {
		if errors.IsNotFound(err) {
			// Resource has been deleted, ignore
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if instance.Status.Phase != StatusPending && instance.Status.Phase != StatusRunning {
		instance.Status.Phase = StatusPending
	}

	switch instance.Status.Phase {
	case StatusPending:
		loc, _ := time.LoadLocation("Local")
		schedule_time, err := time.ParseInLocation("2006-01-02 15:04:05", instance.Spec.Schedule, loc)
		if err != nil {
			return ctrl.Result{}, err
		}
		d := schedule_time.Sub(time.Now().Local())
		if d > 0 {
			reqLogger.Info("continue sleep", fmt.Sprintf("%v", d), "Seconds")
			return ctrl.Result{RequeueAfter: d}, nil
		}
		instance.Status.Phase = StatusRunning
	case StatusRunning:
		dep := &appsv1.Deployment{}
		find := types.NamespacedName{
			Name:      instance.Spec.Deploymentname,
			Namespace: instance.Namespace,
		}
		dep = getDeployment(instance)
		err = controllerutil.SetControllerReference(instance, dep, r.Scheme)
		if err != nil {
			return ctrl.Result{}, err
		}
		err := r.Client.Get(context.TODO(), find, dep)
		if err != nil && errors.IsNotFound(err) {
			err = r.Create(context.TODO(), dep)
			if err != nil {
				return ctrl.Result{}, err
			}
			reqLogger.Info("Create Deployment", "Name", dep.Name)
		} else if err != nil {
			return ctrl.Result{}, err
		} else {
			err = controllerutil.SetControllerReference(instance, dep, r.Scheme)
			if err != nil {
				return ctrl.Result{}, err
			}
		}

	}

	err = r.Status().Update(context.TODO(), instance)
	if err != nil {
		return ctrl.Result{}, nil
	}
	return ctrl.Result{}, nil
}

func getDeployment(instance *crdtryv1.Lpxpod) *appsv1.Deployment {
	labels := map[string]string{
		"app":    instance.Name + "app",
		"lpxpod": instance.Name,
	}
	size := int32(1)
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      instance.Spec.Deploymentname,
			Namespace: instance.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &size,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image: "lpx_test_springboot:v1",
						Name:  instance.Name,
						Ports: []corev1.ContainerPort{{
							ContainerPort: instance.Spec.Podport,
							HostPort:      instance.Spec.Podport,
							Name:          instance.Name,
						}},
						Command: []string{"/usr/bin/java"},
						Args:    []string{"-jar", "/test.jar"},
					}},
				},
			},
		},
	}
	return dep
}

// SetupWithManager sets up the controller with the Manager.
func (r *LpxpodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&crdtryv1.Lpxpod{}).
		Complete(r)
}
