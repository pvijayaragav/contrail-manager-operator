package v1alpha1
	
import (
	"github.com/ghodss/yaml"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

var yamlDataCassandraContrailV1Alpha1ContrailcassandraCrd = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: contrailcassandras.contrail.juniper.net
spec:
  additionalPrinterColumns:
  - JSONPath: .status.active
    description: Defines whether cassandra is active
    name: Active
    type: string
  - JSONPath: .status.nodes
    description: The nodes where cassandra is installed
    name: Nodes
    type: string
  - JSONPath: .status.ports
    description: The ports used by cassandra
    name: Ports
    type: string
  group: contrail.juniper.net
  names:
    kind: ContrailCassandra
    listKind: ContrailCassandraList
    plural: contrailcassandras
    singular: contrailcassandra
  scope: Namespaced
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
        apiVersion:
          type: string
        kind:
          type: string
        metadata:
          type: object
        spec:
          properties:
            containers:
              items:
                properties:
                  configMaps:
                    items:
                      properties:
                        envs:
                          additionalProperties:
                            type: string
                          type: object
                        name:
                          type: string
                      type: object
                    type: array
                  image:
                    type: string
                  name:
                    type: string
                type: object
              type: array
            replicas:
              format: int32
              type: integer
          type: object
        status:
          properties:
            active:
              type: boolean
            nodes:
              items:
                type: string
              type: array
            ports:
              items:
                type: string
              type: array
          type: object
  version: v1alpha1
  versions:
  - name: v1alpha1
    served: true
    storage: true
`
func CassandraContrailV1Alpha1ContrailcassandraCrd() *apiextensionsv1beta1.CustomResourceDefinition{
	crd := apiextensionsv1beta1.CustomResourceDefinition{}
	err := yaml.Unmarshal([]byte(yamlDataCassandraContrailV1Alpha1ContrailcassandraCrd), &crd)
	if err != nil {
		panic(err)
	}
	jsonData, err := yaml.YAMLToJSON([]byte(yamlDataCassandraContrailV1Alpha1ContrailcassandraCrd))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(jsonData), &crd)
	if err != nil {
		panic(err)
	}
	return &crd
}
		