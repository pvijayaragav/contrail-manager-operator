package configs

import (
	"fmt"

	"context"

	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	configs "github.com/operators/contrail-manager-test-1/pkg/configs/cassandra"
	appsv1 "k8s.io/api/apps/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// Cassandra configure function
func Configure(crName string, request reconcile.Request,
	r client.Client) bool {
	return configureConfigMap(crName, request, r) && configureDeployment(crName, request, r)
}

func int32ToString(i int32) string {
	var retString string
	retString = fmt.Sprint(i)
	return retString
}

func UpdateStatus(crName string, request reconcile.Request,
	r client.Client) bool {
	_deployment := &appsv1.Deployment{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: crName + "-" + "deployment", Namespace: request.Namespace}, _deployment)
	if err == nil {
		fmt.Println("Deployment " + _deployment.Name + " found ReadyReplicas : " + int32ToString(_deployment.Status.ReadyReplicas) + " spec : " + int32ToString(*_deployment.Spec.Replicas))
		if _deployment.Status.ReadyReplicas <= *_deployment.Spec.Replicas {
			return true
		}
	} else {
		fmt.Println("Deployment " + _deployment.Name + " not found returning false")
	}
	return false
}

func configureDeployment(crName string, request reconcile.Request,
	r client.Client) bool {
	_dpName := crName + "-" + "deployment"
	_deployment := configs.GetDeployment()
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
	_configMap := configs.GetConfigMap()
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
