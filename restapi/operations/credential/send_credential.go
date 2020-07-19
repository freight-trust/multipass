// Code generated by go-swagger; DO NOT EDIT.

package credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SendCredentialHandlerFunc turns a function with the right signature into a send credential handler
type SendCredentialHandlerFunc func(SendCredentialParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SendCredentialHandlerFunc) Handle(params SendCredentialParams) middleware.Responder {
	return fn(params)
}

// SendCredentialHandler interface for that can handle valid send credential params
type SendCredentialHandler interface {
	Handle(SendCredentialParams) middleware.Responder
}

// NewSendCredential creates a new http.Handler for the send credential operation
func NewSendCredential(ctx *middleware.Context, handler SendCredentialHandler) *SendCredential {
	return &SendCredential{Context: ctx, Handler: handler}
}

/*SendCredential swagger:route POST /credential/send credential sendCredential

send a credential

This service verify the credential exist into blockchain and send the credential by email

*/
type SendCredential struct {
	Context *middleware.Context
	Handler SendCredentialHandler
}

func (o *SendCredential) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSendCredentialParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
