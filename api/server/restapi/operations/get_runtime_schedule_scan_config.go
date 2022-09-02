// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRuntimeScheduleScanConfigHandlerFunc turns a function with the right signature into a get runtime schedule scan config handler
type GetRuntimeScheduleScanConfigHandlerFunc func(GetRuntimeScheduleScanConfigParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRuntimeScheduleScanConfigHandlerFunc) Handle(params GetRuntimeScheduleScanConfigParams) middleware.Responder {
	return fn(params)
}

// GetRuntimeScheduleScanConfigHandler interface for that can handle valid get runtime schedule scan config params
type GetRuntimeScheduleScanConfigHandler interface {
	Handle(GetRuntimeScheduleScanConfigParams) middleware.Responder
}

// NewGetRuntimeScheduleScanConfig creates a new http.Handler for the get runtime schedule scan config operation
func NewGetRuntimeScheduleScanConfig(ctx *middleware.Context, handler GetRuntimeScheduleScanConfigHandler) *GetRuntimeScheduleScanConfig {
	return &GetRuntimeScheduleScanConfig{Context: ctx, Handler: handler}
}

/* GetRuntimeScheduleScanConfig swagger:route GET /runtime/scheduleScan/config getRuntimeScheduleScanConfig

Get runtime scheduled scan configuration

*/
type GetRuntimeScheduleScanConfig struct {
	Context *middleware.Context
	Handler GetRuntimeScheduleScanConfigHandler
}

func (o *GetRuntimeScheduleScanConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRuntimeScheduleScanConfigParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}