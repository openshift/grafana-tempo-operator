// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FieldTypeConfig FieldTypeConfig has type specific configs, only one should be active at a time
//
// swagger:model FieldTypeConfig
type FieldTypeConfig struct {

	// enum
	Enum *EnumFieldConfig `json:"enum,omitempty"`
}

// Validate validates this field type config
func (m *FieldTypeConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FieldTypeConfig) validateEnum(formats strfmt.Registry) error {
	if swag.IsZero(m.Enum) { // not required
		return nil
	}

	if m.Enum != nil {
		if err := m.Enum.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enum")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enum")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this field type config based on the context it is used
func (m *FieldTypeConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnum(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FieldTypeConfig) contextValidateEnum(ctx context.Context, formats strfmt.Registry) error {

	if m.Enum != nil {

		if swag.IsZero(m.Enum) { // not required
			return nil
		}

		if err := m.Enum.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enum")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enum")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FieldTypeConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FieldTypeConfig) UnmarshalBinary(b []byte) error {
	var res FieldTypeConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
