// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// Mixin0HandlerFunc turns a function with the right signature into a mixin0 handler
type Mixin0HandlerFunc func(Mixin0Params) middleware.Responder

// Handle executing the request and returning a response
func (fn Mixin0HandlerFunc) Handle(params Mixin0Params) middleware.Responder {
	return fn(params)
}

// Mixin0Handler interface for that can handle valid mixin0 params
type Mixin0Handler interface {
	Handle(Mixin0Params) middleware.Responder
}

// NewMixin0 creates a new http.Handler for the mixin0 operation
func NewMixin0(ctx *middleware.Context, handler Mixin0Handler) *Mixin0 {
	return &Mixin0{Context: ctx, Handler: handler}
}

/*
	Mixin0 swagger:route POST /auth/register Auth mixin0

# Register user

Save user data in DB
*/
type Mixin0 struct {
	Context *middleware.Context
	Handler Mixin0Handler
}

func (o *Mixin0) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewMixin0Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}