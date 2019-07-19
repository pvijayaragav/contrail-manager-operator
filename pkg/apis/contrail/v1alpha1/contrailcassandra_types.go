package v1alpha1

import (
	"context"
	"fmt"

	cassandraconfig "github.com/operators/contrail-manager-test-1/pkg/configs"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

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

// Implementing the contrailTypes interface
func (crdType ContrailCassandra) CrExists(crName string, request reconcile.Request,
	r client.Client) bool {
	_cr := ContrailCassandra{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: crName, Namespace: request.Namespace}, &_cr)
	if err == nil {
		fmt.Println("CR " + crName + " for CRD ContrailCassandra present")
		return true
	}
	return false
}

func (crdType ContrailCassandra) CreateCr(crName string, request reconcile.Request,
	r client.Client) bool {
	_contrailManager := ContrailManager{}
	err := r.Get(context.TODO(), request.NamespacedName, &_contrailManager)
	if err != nil {
		fmt.Println("Manager CR not found...")
	} else {
		fmt.Println("Manager CR found " + request.NamespacedName.Name)
		// update config map based on manager cr spec
		if cassandraconfig.Configure(crName, request, r) {
			fmt.Println("Configure success")
			_cassandraCr := ContrailCassandra{}
			err = r.Get(context.TODO(), types.NamespacedName{Name: crName, Namespace: request.Namespace}, &_cassandraCr)
			if err == nil {
				fmt.Println("new CR " + crName + "found")
				if cassandraconfig.UpdateStatus(crName, request, r) {
					_cassandraCr.Status.Active = true
				} else {
					_cassandraCr.Status.Active = false
				}
			} else {
				fmt.Println("new CR " + crName + " not found trying to create")
				if errors.IsNotFound(err) {
					_cassandraCr.Name = crName
					_cassandraCr.Namespace = request.Namespace
					_cassandraCr.Spec.Replicas = 1
					err = r.Create(context.TODO(), &_cassandraCr)
					if err == nil {
						if cassandraconfig.UpdateStatus(crName, request, r) {
							_cassandraCr.Status.Active = true
						} else {
							_cassandraCr.Status.Active = false
						}
						fmt.Println("Created new CR " + crName)
					} else {
						fmt.Println("failed to create new CR " + crName)
						return false
					}
				} else {
					fmt.Println("did not create...")
				}
			}
			err = r.Status().Update(context.TODO(), &_cassandraCr)
			if err != nil {
				fmt.Println("Failed to update Cassandra status")
			}
			if cassandraconfig.UpdateStatus(crName, request, r) {
				return true
			}
		}
	}
	return false
}

func init() {
	SchemeBuilder.Register(&ContrailCassandra{}, &ContrailCassandraList{})
}
