// Code generated by go-swagger; DO NOT EDIT.

package logistic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewMixin3Params creates a new Mixin3Params object
//
// There are no default values defined in the spec.
func NewMixin3Params() Mixin3Params {

	return Mixin3Params{}
}

// Mixin3Params contains all the bound params for the mixin3 operation
// typically these are obtained from a http.Request
//
// swagger:parameters Mixin3
type Mixin3Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Authorzation(Bearer random_value)
	  Required: true
	  In: header
	*/
	Authorization string
	/*test example
	  Required: true
	  In: query
	*/
	DestionationName string
	/*test example
	  Required: true
	  In: query
	*/
	OriginName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewMixin3Params() beforehand.
func (o *Mixin3Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qDestionationName, qhkDestionationName, _ := qs.GetOK("destionation_name")
	if err := o.bindDestionationName(qDestionationName, qhkDestionationName, route.Formats); err != nil {
		res = append(res, err)
	}

	qOriginName, qhkOriginName, _ := qs.GetOK("origin_name")
	if err := o.bindOriginName(qOriginName, qhkOriginName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorization binds and validates parameter Authorization from header.
func (o *Mixin3Params) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Authorization", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Authorization", "header", raw); err != nil {
		return err
	}
	o.Authorization = raw

	return nil
}

// bindDestionationName binds and validates parameter DestionationName from query.
func (o *Mixin3Params) bindDestionationName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("destionation_name", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("destionation_name", "query", raw); err != nil {
		return err
	}
	o.DestionationName = raw

	return nil
}

// bindOriginName binds and validates parameter OriginName from query.
func (o *Mixin3Params) bindOriginName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("origin_name", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("origin_name", "query", raw); err != nil {
		return err
	}
	o.OriginName = raw

	return nil
}