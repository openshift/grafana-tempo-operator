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

// AnnotationQuery TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
//
// swagger:model AnnotationQuery
type AnnotationQuery struct {

	// Set to 1 for the standard annotation query all dashboards have by default.
	BuiltIn float32 `json:"builtIn,omitempty"`

	// datasource
	Datasource *DataSourceRef `json:"datasource,omitempty"`

	// When enabled the annotation query is issued with every dashboard refresh
	Enable bool `json:"enable,omitempty"`

	// filter
	Filter *AnnotationPanelFilter `json:"filter,omitempty"`

	// Annotation queries can be toggled on or off at the top of the dashboard.
	// When hide is true, the toggle is not shown in the dashboard.
	Hide bool `json:"hide,omitempty"`

	// Color to use for the annotation event markers
	IconColor string `json:"iconColor,omitempty"`

	// Name of annotation.
	Name string `json:"name,omitempty"`

	// target
	Target *AnnotationTarget `json:"target,omitempty"`

	// TODO -- this should not exist here, it is based on the --grafana-- datasource
	Type string `json:"type,omitempty"`
}

// Validate validates this annotation query
func (m *AnnotationQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatasource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnnotationQuery) validateDatasource(formats strfmt.Registry) error {
	if swag.IsZero(m.Datasource) { // not required
		return nil
	}

	if m.Datasource != nil {
		if err := m.Datasource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datasource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datasource")
			}
			return err
		}
	}

	return nil
}

func (m *AnnotationQuery) validateFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *AnnotationQuery) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this annotation query based on the context it is used
func (m *AnnotationQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatasource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnnotationQuery) contextValidateDatasource(ctx context.Context, formats strfmt.Registry) error {

	if m.Datasource != nil {

		if swag.IsZero(m.Datasource) { // not required
			return nil
		}

		if err := m.Datasource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datasource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datasource")
			}
			return err
		}
	}

	return nil
}

func (m *AnnotationQuery) contextValidateFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.Filter != nil {

		if swag.IsZero(m.Filter) { // not required
			return nil
		}

		if err := m.Filter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

func (m *AnnotationQuery) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnnotationQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnnotationQuery) UnmarshalBinary(b []byte) error {
	var res AnnotationQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}