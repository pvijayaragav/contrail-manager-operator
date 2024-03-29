package controller

import (
	"github.com/operators/contrail-manager-test-1/pkg/controller/contrailcassandra"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, contrailcassandra.Add)
}
