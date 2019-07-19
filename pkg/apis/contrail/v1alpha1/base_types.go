package v1alpha1

// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:openapi-gen=true
type Status struct {
	Active bool     `json:"active,omitempty"`
	Nodes  []string `json:"nodes,omitempty"`
	Ports  []string `json:"ports,omitempty"`
}

type ConfigMap struct {
	Name string            `json:"name,omitempty"`
	Envs map[string]string `json:"envs,omitempty"`
}
