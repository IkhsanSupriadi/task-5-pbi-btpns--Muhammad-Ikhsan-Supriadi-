// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PbCreateLogisticRequest pb create logistic request
//
// swagger:model pb.CreateLogisticRequest
type PbCreateLogisticRequest struct {

	// amount
	Amount int64 `json:"amount,omitempty"`

	// destination name
	DestinationName string `json:"destinationName,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// logistic name
	LogisticName string `json:"logisticName,omitempty"`

	// origin name
	OriginName string `json:"originName,omitempty"`
}

// Validate validates this pb create logistic request
func (m *PbCreateLogisticRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pb create logistic request based on context it is used
func (m *PbCreateLogisticRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PbCreateLogisticRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbCreateLogisticRequest) UnmarshalBinary(b []byte) error {
	var res PbCreateLogisticRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
