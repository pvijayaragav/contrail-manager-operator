package v1alpha1

import (
	"context"
	"errors"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	cassandramanifests "github.com/operators/contrail-manager-test-1/pkg/configs/cassandra"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
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

func init() {
	SchemeBuilder.Register(&ContrailCassandra{}, &ContrailCassandraList{})
}

func int32ToString(i int32) string {
	var retString string
	retString = fmt.Sprint(i)
	return retString
}

// CreateCustomResourceDefinition creates the crd
func (crdType *ContrailCassandra) CreateCustomResourceDefinition(client client.Client) bool {
	crd := CassandraContrailV1Alpha1ContrailcassandraCrd()
	err := client.Create(context.TODO(), crd)
	if err != nil {
		return false
	}
	return true
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
		if k8serrors.IsNotFound(err) {
			if createCR(instanceName, request, client) && configure(instanceName, request, client) {
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

// GetStatusOfInstance updates status for a instance object
func (crdType *ContrailCassandra) GetStatusOfInstance(instance runtime.Object, request reconcile.Request, client client.Client) error {
	// instanceName := crdType.GetInstanceName(request)
	return nil
}

// UpdateStatus updates the status of the instance
func (crdType *ContrailCassandra) UpdateStatus(crName string, request reconcile.Request,
	r client.Client) error {
	_deployment := &appsv1.Deployment{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: crName + "-" + "deployment", Namespace: request.Namespace}, _deployment)
	if err == nil {
		fmt.Println("Deployment " + _deployment.Name + " found ReadyReplicas : " + int32ToString(_deployment.Status.ReadyReplicas) + " spec : " + int32ToString(*_deployment.Spec.Replicas))
		if _deployment.Status.ReadyReplicas <= *_deployment.Spec.Replicas {
			return nil
		}
	} else {
		fmt.Println("Deployment " + _deployment.Name + " not found returning false")
	}
	return errors.New("Error in updating")
}

func createCR(crName string, request reconcile.Request,
	r client.Client) bool {
	cr := CassandraContrailV1Alpha1ContrailcassandraCr()
	cr.Name = crName
	cr.Namespace = request.Namespace
	err := r.Create(context.TODO(), cr)
	if err == nil {
		return true
	}
	return false
}

func configure(instanceName string, request reconcile.Request,
	client client.Client) bool {
	if cassandramanifests.Configure(instanceName, request, client) {
		return true
	}
	return false
}
