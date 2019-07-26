package contrailmanager

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// InstanceInterface is the type which all instances to implement
type InstanceInterface interface {
	GetInstanceName(reconcile.Request) string
	CreateInstance(reconcile.Request, client.Client) runtime.Object
	ReadInstance(reconcile.Request, client.Client) error
	UpdateInstance(runtime.Object, reconcile.Request, client.Client) error
	DeleteInstance(reconcile.Request, client.Client) error
	GetStatusOfInstance(reconcile.Request, client.Client) map[string]string
	GetOwnerOfInstance(reconcile.Request, client.Client) error
}
