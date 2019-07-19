package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:openapi-gen=true
type ContrailManagerSpec struct {
	GlobalConfig GlobalConfig `json:"globalConfig,omitempty"`
	NodeConfig   NodeConfig   `json:"nodeConfig,omitempty"`
	ConfigMaps   []ConfigMap  `json:"configMaps,omitempty"`
	Components   []Component  `json:"components,omitempty"`
}

// +k8s:openapi-gen=true
type ContrailManagerStatus struct {
	Active          bool     `json:"active,omitempty"`
	Completed       bool     `json:"completed,omitempty"`
	Platform        Platform `json:"platform,omitempty"`
	CassandraStatus bool     `json:"cassandraStatus,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContrailManager is the Schema for the contrailmanagers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ContrailManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContrailManagerSpec   `json:"spec,omitempty"`
	Status ContrailManagerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContrailManagerList contains a list of ContrailManager
type ContrailManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContrailManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContrailManager{}, &ContrailManagerList{})
}
