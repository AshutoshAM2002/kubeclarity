// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Finding finding
//
// swagger:model Finding
type Finding struct {

	// analysis
	Analysis map[string]interface{} `json:"analysis,omitempty"`

	// attribution
	Attribution map[string]interface{} `json:"attribution,omitempty"`

	// component
	Component map[string]interface{} `json:"component,omitempty"`

	// matrix
	Matrix string `json:"matrix,omitempty"`

	// vulnerability
	Vulnerability map[string]interface{} `json:"vulnerability,omitempty"`
}

// Validate validates this finding
func (m *Finding) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this finding based on context it is used
func (m *Finding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Finding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Finding) UnmarshalBinary(b []byte) error {
	var res Finding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}