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

// DataSource data source
//
// swagger:model DataSource
type DataSource struct {

	// access
	Access DsAccess `json:"access,omitempty"`

	// access control
	AccessControl Metadata `json:"accessControl,omitempty"`

	// basic auth
	BasicAuth bool `json:"basicAuth,omitempty"`

	// basic auth user
	BasicAuthUser string `json:"basicAuthUser,omitempty"`

	// database
	Database string `json:"database,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// is default
	IsDefault bool `json:"isDefault,omitempty"`

	// json data
	JSONData JSON `json:"jsonData,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID int64 `json:"orgId,omitempty"`

	// read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// secure Json fields
	SecureJSONFields map[string]bool `json:"secureJsonFields,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// type logo Url
	TypeLogoURL string `json:"typeLogoUrl,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// user
	User string `json:"user,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`

	// with credentials
	WithCredentials bool `json:"withCredentials,omitempty"`
}

// Validate validates this data source
func (m *DataSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessControl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSource) validateAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Access) { // not required
		return nil
	}

	if err := m.Access.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("access")
		}
		return err
	}

	return nil
}

func (m *DataSource) validateAccessControl(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessControl) { // not required
		return nil
	}

	if m.AccessControl != nil {
		if err := m.AccessControl.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessControl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessControl")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this data source based on the context it is used
func (m *DataSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAccessControl(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSource) contextValidateAccess(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Access) { // not required
		return nil
	}

	if err := m.Access.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("access")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("access")
		}
		return err
	}

	return nil
}

func (m *DataSource) contextValidateAccessControl(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AccessControl) { // not required
		return nil
	}

	if err := m.AccessControl.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("accessControl")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("accessControl")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSource) UnmarshalBinary(b []byte) error {
	var res DataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}