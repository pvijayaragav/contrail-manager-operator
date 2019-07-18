package contrailmanager

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type instanceType string

const (
	contrailManager instanceType = "managerInstance"
	cassandra       instanceType = "cassandra"
)

type typeInterface interface {
	GetInstanceType(reconcile.Request, client.Client) instanceType
	CustomResourceExists(reconcile.Request, client.Client) bool
	CreateCustomResource(reconcile.Request, client.Client) bool
	UpdateCustomResource(reconcile.Request, client.Client) bool
	DeleteCustomResource(reconcile.Request, client.Client) bool
	GetStatusOfCustomResource(reconcile.Request, client.Client) map[string]string
}
