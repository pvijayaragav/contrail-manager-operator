package v1alpha1

import (
	"github.com/ghodss/yaml"
)

var yamlDataCassandraContrailV1Alpha1ContrailcassandraCr = `
apiVersion: contrail.juniper.net/v1alpha1
kind: ContrailCassandra
metadata:
  name: example-contrailcassandra
spec:
  # Add fields here
  size: 3
`

// CassandraContrailV1Alpha1ContrailcassandraCr creates the CR
func CassandraContrailV1Alpha1ContrailcassandraCr() *ContrailCassandra {
	cr := ContrailCassandra{}
	err := yaml.Unmarshal([]byte(yamlDataCassandraContrailV1Alpha1ContrailcassandraCr), &cr)
	if err != nil {
		panic(err)
	}
	jsonData, err := yaml.YAMLToJSON([]byte(yamlDataCassandraContrailV1Alpha1ContrailcassandraCr))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(jsonData), &cr)
	if err != nil {
		panic(err)
	}
	return &cr
}
