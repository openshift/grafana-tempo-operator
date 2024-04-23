// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateAlertNotificationCommand create alert notification command
//
// swagger:model CreateAlertNotificationCommand
type CreateAlertNotificationCommand struct {

	// disable resolve message
	DisableResolveMessage bool `json:"disableResolveMessage,omitempty"`

	// frequency
	Frequency string `json:"frequency,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// secure settings
	SecureSettings map[string]string `json:"secureSettings,omitempty"`

	// send reminder
	SendReminder bool `json:"sendReminder,omitempty"`

	// settings
	Settings JSON `json:"settings,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this create alert notification command
func (m *CreateAlertNotificationCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create alert notification command based on context it is used
func (m *CreateAlertNotificationCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAlertNotificationCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAlertNotificationCommand) UnmarshalBinary(b []byte) error {
	var res CreateAlertNotificationCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
