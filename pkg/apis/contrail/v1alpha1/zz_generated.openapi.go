// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailCassandra":     schema_pkg_apis_contrail_v1alpha1_ContrailCassandra(ref),
		"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailCassandraSpec": schema_pkg_apis_contrail_v1alpha1_ContrailCassandraSpec(ref),
		"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailManager":       schema_pkg_apis_contrail_v1alpha1_ContrailManager(ref),
		"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailManagerSpec":   schema_pkg_apis_contrail_v1alpha1_ContrailManagerSpec(ref),
		"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailManagerStatus": schema_pkg_apis_contrail_v1alpha1_ContrailManagerStatus(ref),
		"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.Status":                schema_pkg_apis_contrail_v1alpha1_Status(ref),
	}
}

func schema_pkg_apis_contrail_v1alpha1_ContrailCassandra(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContrailCassandra is the Schema for the contrailcassandras API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailCassandraSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.Status"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailCassandraSpec", "github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.Status", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_contrail_v1alpha1_ContrailCassandraSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContrailCassandraSpec defines the desired state of ContrailCassandra",
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"containers": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.Container"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.Container"},
	}
}

func schema_pkg_apis_contrail_v1alpha1_ContrailManager(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ContrailManager is the Schema for the contrailmanagers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailManagerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailManagerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailManagerSpec", "github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ContrailManagerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_contrail_v1alpha1_ContrailManagerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"globalConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.GlobalConfig"),
						},
					},
					"nodeConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.NodeConfig"),
						},
					},
					"configMaps": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ConfigMap"),
									},
								},
							},
						},
					},
					"components": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.Component"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.Component", "github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.ConfigMap", "github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.GlobalConfig", "github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.NodeConfig"},
	}
}

func schema_pkg_apis_contrail_v1alpha1_ContrailManagerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"active": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"completed": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"platform": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.Platform"),
						},
					},
					"cassandraStatus": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1.Platform"},
	}
}

func schema_pkg_apis_contrail_v1alpha1_Status(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Properties: map[string]spec.Schema{
					"active": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"ports": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
