package cassandra

import (
	"fmt"

	"context"

	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// crds "github.com/operators/contrail-manager-test-1/pkg/configs/cassandra/crds"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Configure creates the deployments and cms
func Configure(crName string, request reconcile.Request,
	r client.Client) bool {
	return configureConfigMap(crName, request, r) && configureDeployment(crName, request, r)
}

func configureDeployment(crName string, request reconcile.Request,
	r client.Client) bool {
	_dpName := crName + "-" + "deployment"
	_deployment := &appsv1.Deployment{}
	// _deployment := Deployment1()
	_deployment.Name = _dpName
	_deployment.Namespace = request.Namespace
	_deployment.Spec.Template.Spec.InitContainers[0].EnvFrom[0].ConfigMapRef.Name = crName + "-" + "env"
	_deployment.Spec.Template.Spec.Containers[0].EnvFrom[0].ConfigMapRef.Name = crName + "-" + "env"
	err := r.Get(context.TODO(), types.NamespacedName{Name: _dpName, Namespace: request.Namespace}, _deployment)
	if err != nil {
		if errors.IsNotFound(err) {
			err = r.Create(context.TODO(), _deployment)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Deployment " + _dpName + " for CRD ContrailCassandra failed")
			} else {
				fmt.Println("Deployment " + _dpName + " for CRD ContrailCassandra Created")
				return true
			}
		}
	} else {
		fmt.Println("Deployment " + _dpName + " for CRD ContrailCassandra already present")
		return true
	}
	return false
}

func configureConfigMap(crName string, request reconcile.Request,
	r client.Client) bool {
	_cmName := crName + "-" + "env"
	_configMap := &corev1.ConfigMap{}
	// _configMap := Configmap1()
	_configMap.Name = _cmName
	_configMap.Namespace = request.Namespace
	err := r.Get(context.TODO(), types.NamespacedName{Name: _cmName, Namespace: request.Namespace}, _configMap)
	if err != nil {
		if errors.IsNotFound(err) {
			err = r.Create(context.TODO(), _configMap)
			if err != nil {
				fmt.Println("Configmap " + _cmName + " for CRD ContrailCassandra failed")
			} else {
				fmt.Println("Configmap " + _cmName + " for CRD ContrailCassandra Created")
				return true
			}
		}
	} else {
		fmt.Println("Configmap " + _cmName + " for CRD ContrailCassandra already present")
		return true
	}
	return false
}
