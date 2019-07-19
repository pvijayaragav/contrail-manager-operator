package contrailmanager

import (
	contrailv1alpha1 "github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type instanceType runtime.Object

var contrailManager instanceType = &contrailv1alpha1.ContrailManager{}
var cassandra instanceType = &contrailv1alpha1.ContrailCassandra{}
var zookeeper instanceType = &contrailv1alpha1.ContrailManager{}
var rabbitmq instanceType = &contrailv1alpha1.ContrailCassandra{}
var contrailConfig instanceType = &contrailv1alpha1.ContrailManager{}
var contrailControl instanceType = &contrailv1alpha1.ContrailCassandra{}

type instanceInterface interface {
	GetInstanceType(reconcile.Request, client.Client) instanceType
	CustomResourceExists(reconcile.Request, client.Client) bool
	CreateCustomResource(reconcile.Request, client.Client) bool
	UpdateCustomResource(reconcile.Request, client.Client) bool
	DeleteCustomResource(reconcile.Request, client.Client) bool
	GetStatusOfCustomResource(reconcile.Request, client.Client) map[string]string
}
