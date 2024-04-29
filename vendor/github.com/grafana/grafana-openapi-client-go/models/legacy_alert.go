// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LegacyAlert legacy alert
//
// swagger:model LegacyAlert
type LegacyAlert struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"Created,omitempty"`

	// dashboard ID
	DashboardID int64 `json:"DashboardID,omitempty"`

	// eval data
	EvalData JSON `json:"EvalData,omitempty"`

	// execution error
	ExecutionError string `json:"ExecutionError,omitempty"`

	// for
	For Duration `json:"For,omitempty"`

	// frequency
	Frequency int64 `json:"Frequency,omitempty"`

	// handler
	Handler int64 `json:"Handler,omitempty"`

	// ID
	ID int64 `json:"ID,omitempty"`

	// message
	Message string `json:"Message,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// new state date
	// Format: date-time
	NewStateDate strfmt.DateTime `json:"NewStateDate,omitempty"`

	// org ID
	OrgID int64 `json:"OrgID,omitempty"`

	// panel ID
	PanelID int64 `json:"PanelID,omitempty"`

	// settings
	Settings JSON `json:"Settings,omitempty"`

	// severity
	Severity string `json:"Severity,omitempty"`

	// silenced
	Silenced bool `json:"Silenced,omitempty"`

	// state
	State AlertStateType `json:"State,omitempty"`

	// state changes
	StateChanges int64 `json:"StateChanges,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"Updated,omitempty"`

	// version
	Version int64 `json:"Version,omitempty"`
}

// Validate validates this legacy alert
func (m *LegacyAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNewStateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LegacyAlert) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("Created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyAlert) validateFor(formats strfmt.Registry) error {
	if swag.IsZero(m.For) { // not required
		return nil
	}

	if err := m.For.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("For")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("For")
		}
		return err
	}

	return nil
}

func (m *LegacyAlert) validateNewStateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.NewStateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("NewStateDate", "body", "date-time", m.NewStateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LegacyAlert) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("State")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("State")
		}
		return err
	}

	return nil
}

func (m *LegacyAlert) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("Updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this legacy alert based on the context it is used
func (m *LegacyAlert) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LegacyAlert) contextValidateFor(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.For) { // not required
		return nil
	}

	if err := m.For.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("For")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("For")
		}
		return err
	}

	return nil
}

func (m *LegacyAlert) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("State")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("State")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LegacyAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LegacyAlert) UnmarshalBinary(b []byte) error {
	var res LegacyAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}