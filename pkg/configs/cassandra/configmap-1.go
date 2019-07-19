package cassandra

import (
	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
)

var yamlDatacassandra = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: contrail-configdb-conf-env
  namespace: default
data:
  ANALYTICSDB_CQL_PORT: "9045"
  ANALYTICSDB_PORT: "9163"
  CASSANDRA_CLUSTER_NAME: ContrailConfigDB
  CASSANDRA_CQL_PORT: "9044"
  CASSANDRA_JMX_LOCAL_PORT: "7204"
  CASSANDRA_LISTEN_ADDRESS: auto
  CASSANDRA_PORT: "9164"
  CASSANDRA_SSL_STORAGE_PORT: "7014"
  CASSANDRA_START_RPC: "true"
  CASSANDRA_STORAGE_PORT: "7013"
  CONFIGDB_CQL_PORT: "9044"
  CONFIGDB_PORT: "9164"
`

func GetConfigMap() *corev1.ConfigMap {
	fileData := configMap
	yamlData := string(fileData)
	cm := corev1.ConfigMap{}
	err := yaml.Unmarshal([]byte(yamlData), &cm)
	if err != nil {
		panic(err)
	}
	jsonData, err := yaml.YAMLToJSON([]byte(yamlData))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(jsonData), &cm)
	if err != nil {
		panic(err)
	}
	return &cm
}
