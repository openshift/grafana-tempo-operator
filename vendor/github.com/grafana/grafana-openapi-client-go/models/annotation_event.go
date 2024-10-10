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

// AnnotationEvent annotation event
//
// swagger:model AnnotationEvent
type AnnotationEvent struct {

	// color
	Color string `json:"color,omitempty"`

	// dashboard Id
	DashboardID int64 `json:"dashboardId,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is region
	IsRegion bool `json:"isRegion,omitempty"`

	// panel Id
	PanelID int64 `json:"panelId,omitempty"`

	// source
	Source *AnnotationQuery `json:"source,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// text
	Text string `json:"text,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// time end
	TimeEnd int64 `json:"timeEnd,omitempty"`
}

// Validate validates this annotation event
func (m *AnnotationEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnnotationEvent) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this annotation event based on the context it is used
func (m *AnnotationEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnnotationEvent) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnnotationEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnnotationEvent) UnmarshalBinary(b []byte) error {
	var res AnnotationEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
