package v1alpha1

import (
	"context"
	"fmt"

	cassandraconfig "github.com/operators/contrail-manager-test-1/pkg/configs/cassandra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.Log.WithName("controller_cassandra")

type Container struct {
	Name       string      `json:"name,omitempty"`
	Image      string      `json:"image,omitempty"`
	ConfigMaps []ConfigMap `json:"configMaps,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ContrailCassandraSpec defines the desired state of ContrailCassandra
// +k8s:openapi-gen=true
type ContrailCassandraSpec struct {
	Replicas   int32       `json:"replicas,omitempty"`
	Containers []Container `json:"containers,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContrailCassandra is the Schema for the contrailcassandras API
// +k8s:openapi-gen=true
type ContrailCassandra struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContrailCassandraSpec `json:"spec,omitempty"`
	Status Status                `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContrailCassandraList contains a list of ContrailCassandra
type ContrailCassandraList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContrailCassandra `json:"items"`
}

// GetInstanceName creates a name for the instance
func (crdType *ContrailCassandra) GetInstanceName(request reconcile.Request) string {
	return request.Name + "-cassandra"
}

// CreateInstance creates a new instance
func (crdType *ContrailCassandra) CreateInstance(request reconcile.Request,
	client client.Client) runtime.Object {
	contrailManager := ContrailManager{}
	instanceName := crdType.GetInstanceName(request)
	err := client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: request.Namespace}, &contrailManager)
	if err != nil {
		if errors.IsNotFound(err) {
			if cassandraconfig.Configure(instanceName, request, client) {
				newInstance := ContrailCassandra{}
				err = client.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: request.Namespace}, &newInstance)
				if err == nil {
					return newInstance.DeepCopyObject()
				}
			}
		}
	}
	return nil
}

// ReadInstance is Implementing the contrailTypes interface
func (crdType *ContrailCassandra) ReadInstance(instanceName string, request reconcile.Request,
	r client.Client) error {
	instance := ContrailCassandra{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: instanceName, Namespace: request.Namespace}, &instance)
	if err == nil {
		fmt.Println("Custom Resource Instance " + instanceName + " for CRD ContrailCassandra present")
		return nil
	}
	return err
}

// UpdateInstance updates an instance which is already present with a new object
func (crdType *ContrailCassandra) UpdateInstance(oldObj runtime.Object, newObj runtime.Object, request reconcile.Request, client client.Client) error {
	return nil
}

// UpdateStatus updates status for a instance object
func (crdType *ContrailCassandra) UpdateStatus(instance runtime.Object, request reconcile.Request, client client.Client) error {
	return nil
}

func init() {
	SchemeBuilder.Register(&ContrailCassandra{}, &ContrailCassandraList{})
}
