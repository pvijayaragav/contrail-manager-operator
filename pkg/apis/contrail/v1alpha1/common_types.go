package v1alpha1

// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Platform struct {
	Orchestrator string `json:"orchestrator,omitempty"`
}

type Registry struct {
	Name     string `json:"name,omitempty"`
	Tag      string `json:"tag,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type GlobalConfig struct {
	Registry         Registry          `json:"registry,omitempty"`
	ImagePullSecrets []string          `json:"imagePullSecrets,omitempty"`
	ServiceAccount   string            `json:"serviceAccount,omitempty"`
	NodeSelector     map[string]string `json:"nodeSelector,omitempty"`
	Replicas         int32             `json:"replicas,omitempty"`
	HostNetwork      bool              `json:"hostNetwork,omitempty"`
	Platform         Platform          `json:"platform,omitempty"`
}

type NodeConfig struct {
	NodeIP     string      `json:"nodeIp,omitempty"`
	Hostname   string      `json:"hostname,omitempty"`
	Components []string    `json:"components,omitempty"`
	ConfigMaps []ConfigMap `json:"configMaps,omitempty"`
}

type Component struct {
	Name string `json:"name,omitempty"`
	CRD  string `json:"crd,omitempty"`
}
