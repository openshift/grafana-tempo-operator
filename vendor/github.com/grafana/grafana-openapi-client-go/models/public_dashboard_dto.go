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

// PublicDashboardDTO public dashboard DTO
//
// swagger:model PublicDashboardDTO
type PublicDashboardDTO struct {

	// access token
	AccessToken string `json:"accessToken,omitempty"`

	// annotations enabled
	AnnotationsEnabled bool `json:"annotationsEnabled,omitempty"`

	// is enabled
	IsEnabled bool `json:"isEnabled,omitempty"`

	// share
	Share ShareType `json:"share,omitempty"`

	// time selection enabled
	TimeSelectionEnabled bool `json:"timeSelectionEnabled,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this public dashboard DTO
func (m *PublicDashboardDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShare(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicDashboardDTO) validateShare(formats strfmt.Registry) error {
	if swag.IsZero(m.Share) { // not required
		return nil
	}

	if err := m.Share.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("share")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("share")
		}
		return err
	}

	return nil
}

// ContextValidate validate this public dashboard DTO based on the context it is used
func (m *PublicDashboardDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateShare(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicDashboardDTO) contextValidateShare(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Share) { // not required
		return nil
	}

	if err := m.Share.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("share")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("share")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicDashboardDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicDashboardDTO) UnmarshalBinary(b []byte) error {
	var res PublicDashboardDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
