package contrailmanager

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

var crdMap = map[string]string{
	"cassandra":            "ContrailCassandra",
	"zookeeper":            "ContrailZookeeper",
	"rabbitmq":             "ContrailRabbitmq",
	"contrail-config":      "ContrailConfig",
	"contrail-control":     "ContrailControl",
	"contrail-analytics":   "ContrailAnalytics",
	"contrail-vrouter":     "ContrailVrouter",
	"contrail-kubemanager": "ContrailKubeManager",
	"contrail-webui":       "ContrailWebui",
}

var contrailStatus = map[string]string{
	"cassandra":            "CassandraStatus",
	"zookeeper":            "ZookeeperStatus",
	"rabbitmq":             "RabbitmqStatus",
	"contrail-config":      "ConfigStatus",
	"contrail-control":     "ControlStatus",
	"contrail-analytics":   "AnalyticsStatus",
	"contrail-vrouter":     "VrouterStatus",
	"contrail-kubemanager": "KubeManagerStatus",
	"contrail-webui":       "WebUIStatus",
}
