package contrailmanager

import (
	"github.com/iancoleman/strcase"
	contrailv1alpha1 "github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
)

const group string = "contrail.juniper.net"
const version string = "v1aplpha1"

var contrailServices = [...]string{
	"cassandra",
	"zookeeper",
	"rabbitmq",
	"contrail-config",
	"contrail-control",
	"contrail-analytics",
	"contrail-vrouter",
	"contrail-kubemanager",
	"contrail-webui",
}

type instanceType runtime.Object

var runtimeObjectMap = map[string]instanceType{
	"cassandra": &contrailv1alpha1.ContrailCassandra{},
}

// go:generate go run generate_type_map.go

// ContrailService is the struct containing all service info
type ContrailService struct {
	name                     string
	customResourceName       string
	customResourceSpecName   string
	customResourceStatusName string
	deploymentName           string
	daemonSetName            string
	configMapPrefix          string
	runtimeObject            instanceType
}

// ContrailServicesMap is the map of all services present above
var ContrailServicesMap map[string]ContrailService

func init() {
	for _, service := range contrailServices {
		ContrailServicesMap[service] = ContrailService{
			name:                     service,
			customResourceName:       strcase.ToCamel(service),
			customResourceSpecName:   strcase.ToCamel(service) + "Spec",
			customResourceStatusName: strcase.ToCamel(service) + "Status",
			deploymentName:           strcase.ToKebab(service) + "-deployment",
			configMapPrefix:          strcase.ToKebab(service) + "-cm-",
			runtimeObject:            runtimeObjectMap[service],
		}
	}
}
