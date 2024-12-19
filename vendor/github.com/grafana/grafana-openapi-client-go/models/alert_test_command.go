// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlertTestCommand alert test command
//
// swagger:model AlertTestCommand
type AlertTestCommand struct {

	// dashboard
	Dashboard JSON `json:"dashboard,omitempty"`

	// panel Id
	PanelID int64 `json:"panelId,omitempty"`
}

// Validate validates this alert test command
func (m *AlertTestCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alert test command based on context it is used
func (m *AlertTestCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertTestCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertTestCommand) UnmarshalBinary(b []byte) error {
	var res AlertTestCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}