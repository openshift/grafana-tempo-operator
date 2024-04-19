// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateProviderSettingsParamsBody update provider settings params body
//
// swagger:model updateProviderSettingsParamsBody
type UpdateProviderSettingsParamsBody struct {

	// id
	ID string `json:"id,omitempty"`

	// provider
	Provider string `json:"provider,omitempty"`

	// settings
	Settings interface{} `json:"settings,omitempty"`
}

// Validate validates this update provider settings params body
func (m *UpdateProviderSettingsParamsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update provider settings params body based on context it is used
func (m *UpdateProviderSettingsParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProviderSettingsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProviderSettingsParamsBody) UnmarshalBinary(b []byte) error {
	var res UpdateProviderSettingsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
