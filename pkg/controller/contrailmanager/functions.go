package contrailmanager

import (
	"context"
	"errors"
	"fmt"
	"strings"

	// "time"

	"github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1"
	contrailv1alpha1 "github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1"
	cassandracontroller "github.com/operators/contrail-manager-test-1/pkg/controller/contrailcassandra"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	operatorManager "sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	//  logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

func getInstanceNameForService(service ContrailService, request reconcile.Request) string {
	return service.name + "-" + request.Name
}

func getRuntimeObject(service ContrailService) runtime.Object {
	return service.runtimeObject
}

func checkCrdPresent(crdType string, request reconcile.Request, r client.Client) bool {
	// check if crd present already
	crdFullName := strings.ToLower(crdType) + "s" + "." + Group
	crd := apiextensionsv1beta1.CustomResourceDefinition{}

	err := r.Get(context.TODO(), types.NamespacedName{Name: crdFullName}, &crd)
	if err != nil {
		fmt.Println("CRD not found")
		return false
	}
	fmt.Println("CRD found")
	return true
}

// CreateContrailService creates the contrail service
func CreateContrailService(service ContrailService, request reconcile.Request,
	client client.Client, managerScheme *runtime.Scheme, managerInstance runtime.Object, operatorManager operatorManager.Manager) error {
	rto := service.runtimeObject
	kind := rto.GetObjectKind()
	gvk := kind.GroupVersionKind()
	objType := gvk.Kind
	gkv := schema.FromAPIVersionAndKind(gvk.Group+"/"+gvk.Version, gvk.Kind)
	newObj, err := scheme.Scheme.New(gkv)
	if err != nil {
		newObj := v1alpha1.ContrailCassandra{}
		newObj.CreateCustomResourceDefinition(client)
		// Add controller to manager
		return cassandracontroller.Add(operatorManager)
	}
	if checkCrdPresent(objType, request, client) {
		return errors.New("crd not found")
	}
	switch objType {
	case "ContrailCassandra":
		instanceObject := newObj.(*v1alpha1.ContrailCassandra)
		instanceName := instanceObject.GetInstanceName(request)
		err := instanceObject.ReadInstance(instanceName, request, client)
		if err == nil {
			fmt.Println("Cassandra instance = " + instanceName + "already present")
			// Update Custom Resource logic goes here
			return nil
		}
		if k8serrors.IsNotFound(err) {
			fmt.Println("Cassandra instance = " + instanceName + "not found creating...")
			// Create the custom resource instance here
			rto := instanceObject.CreateInstance(request, client)
			service.runtimeObject = rto
			rtoObj := rto.(*v1alpha1.ContrailCassandra)
			// Set ContrailCassandra instance as the owner and controller
			if err := controllerutil.SetControllerReference(managerInstance.(*contrailv1alpha1.ContrailManager), rtoObj, managerScheme); err != nil {
				return err
			}
			return rtoObj.UpdateStatus(rtoObj.Name, request, client)
		}
	}
	return nil
}
