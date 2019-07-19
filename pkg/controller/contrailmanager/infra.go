package contrailmanager

// const group string = "contrail.juniper.net"
// const version string = "v1aplpha1"

// var contrailServices = [...]string{
// 	"cassandra",
// 	"zookeeper",
// 	"rabbitmq",
// 	"contrail-config",
// 	"contrail-control",
// 	"contrail-analytics",
// 	"contrail-vrouter",
// 	"contrail-kubemanager",
// 	"contrail-webui",
// }

// ContrailService is the struct containing all service info
type ContrailService struct {
	name                     string
	customResourceName       string
	customResourceSpecName   string
	customResourceStatusName string
	deploymentName           string
	daemonSetName            string
	configMapPrefix          string
}

// ContrailServicesMap is the map of all services present above
var ContrailServicesMap map[string]ContrailService

// func init() {
// 	for _, service := range contrailServices {
// 		ContrailServicesMap[service] = ContrailService{
// 			name:                     service,
// 			customResourceName:       strcase.ToCamel(service),
// 			customResourceSpecName:   strcase.ToCamel(service) + "Spec",
// 			customResourceStatusName: strcase.ToCamel(service) + "Status",
// 			deploymentName:           strcase.ToKebab(service) + "-deployment",
// 			configMapPrefix:          strcase.ToKebab(service) + "-cm-",
// 		}
// 	}
// }
