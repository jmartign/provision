package dhcp_leases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/rackn/rocket-skates/models"
)

// DELETEDhcpLeaseHandlerFunc turns a function with the right signature into a d e l e t e dhcp lease handler
type DELETEDhcpLeaseHandlerFunc func(DELETEDhcpLeaseParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn DELETEDhcpLeaseHandlerFunc) Handle(params DELETEDhcpLeaseParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// DELETEDhcpLeaseHandler interface for that can handle valid d e l e t e dhcp lease params
type DELETEDhcpLeaseHandler interface {
	Handle(DELETEDhcpLeaseParams, *models.Principal) middleware.Responder
}

// NewDELETEDhcpLease creates a new http.Handler for the d e l e t e dhcp lease operation
func NewDELETEDhcpLease(ctx *middleware.Context, handler DELETEDhcpLeaseHandler) *DELETEDhcpLease {
	return &DELETEDhcpLease{Context: ctx, Handler: handler}
}

/*DELETEDhcpLease swagger:route DELETE /leases/{id} Dhcp leases dELETEDhcpLease

Delete DHCP Lease

*/
type DELETEDhcpLease struct {
	Context *middleware.Context
	Handler DELETEDhcpLeaseHandler
}

func (o *DELETEDhcpLease) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDELETEDhcpLeaseParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
