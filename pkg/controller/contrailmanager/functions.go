package contrailmanager

import (
	"context"
	"fmt"
	"strings"

	//  "time"

	contrailv1alpha1 "github.com/operators/contrail-manager-test-1/pkg/apis/contrail/v1alpha1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	//  logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

func crdPresent(crdType string,
	request reconcile.Request,
	r client.Client) bool {
	// check if cr present already
	_crdFullName := strings.ToLower(crdType) + "s" + "." + group
	_crd := apiextensionsv1beta1.CustomResourceDefinition{}

	err := r.Get(context.TODO(), types.NamespacedName{Name: _crdFullName}, &_crd)
	if err != nil {
		fmt.Println("CRD not present")
		return false
	}
	fmt.Println("CRD found")
	return true
}

func crPresent(crdType string,
	crName string,
	request reconcile.Request,
	r client.Client) bool {
	// check if cr present already
	if crdPresent(crdType, request, r) {
		switch crdType {
		case "ContrailCassandra":
			return contrailv1alpha1.ContrailCassandra{}.CrExists(crName, request, r)
		}
	}
	return false
}

func createCr(crdType string,
	crName string,
	request reconcile.Request,
	r client.Client) bool {
	switch crdType {
	case "ContrailCassandra":
		return contrailv1alpha1.ContrailCassandra{}.CreateCr(crName, request, r)
	}
	return false
}

// CreateContrailInstance creates an Instance of type Runtime.Object with input from ContrailService
func CreateContrailInstance(service ContrailService, request reconcile.Request) error {
	rto := getRuntimeObject(service)
	objKind := rto.GetObjectKind()
	fmt.Println(objKind)
	return nil
}

func getRuntimeObject(service ContrailService) runtime.Object {
	return service.runtimeObject
}
