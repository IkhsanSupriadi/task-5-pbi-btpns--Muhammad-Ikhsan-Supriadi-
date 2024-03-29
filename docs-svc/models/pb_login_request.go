// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PbLoginRequest pb login request
//
// swagger:model pb.LoginRequest
type PbLoginRequest struct {

	// msisdn
	Msisdn string `json:"msisdn,omitempty"`

	// password
	Password string `json:"password,omitempty"`
}

// Validate validates this pb login request
func (m *PbLoginRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pb login request based on context it is used
func (m *PbLoginRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PbLoginRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbLoginRequest) UnmarshalBinary(b []byte) error {
	var res PbLoginRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
