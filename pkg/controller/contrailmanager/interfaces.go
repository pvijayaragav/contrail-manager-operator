package contrailmanager

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// InstanceInterface is the type which all instances to implement
type InstanceInterface interface {
	GetInstanceType(reconcile.Request, client.Client) instanceType
	InstanceExists(reconcile.Request, client.Client) bool
	CreateInstance(reconcile.Request, client.Client) bool
	UpdateInstance(reconcile.Request, client.Client) bool
	DeleteInstance(reconcile.Request, client.Client) bool
	GetStatusOfInstance(reconcile.Request, client.Client) map[string]string
	GetParentOfInstance(reconcile.Request, client.Client)
}
