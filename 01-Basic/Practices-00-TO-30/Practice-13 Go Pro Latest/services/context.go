// A Context will be automatically created for each HTTP request received by the
// server, which means that all the request handling code that processes
// that request can share the same set of services so that,
// for example, a single struct that provides session information can be used throughout
// processing for a given request.

package services

import (
	"context"
	"reflect"
)

const ServiceKey = "services"

type serviceMap map[reflect.Type]reflect.Value

func NewServiceContext(c context.Context) context.Context {
	if c.Value(ServiceKey) == nil {
		return context.WithValue(c, ServiceKey, make(serviceMap))
	} else {
		return c
	}
}
