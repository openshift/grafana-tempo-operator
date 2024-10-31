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

// DashboardSnapshotDTO DashboardSnapshotDTO without dashboard map
//
// swagger:model DashboardSnapshotDTO
type DashboardSnapshotDTO struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// expires
	// Format: date-time
	Expires strfmt.DateTime `json:"expires,omitempty"`

	// external
	External bool `json:"external,omitempty"`

	// external Url
	ExternalURL string `json:"externalUrl,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this dashboard snapshot DTO
func (m *DashboardSnapshotDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpires(formats); err != nil {
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

func (m *DashboardSnapshotDTO) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardSnapshotDTO) validateExpires(formats strfmt.Registry) error {
	if swag.IsZero(m.Expires) { // not required
		return nil
	}

	if err := validate.FormatOf("expires", "body", "date-time", m.Expires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardSnapshotDTO) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dashboard snapshot DTO based on context it is used
func (m *DashboardSnapshotDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DashboardSnapshotDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardSnapshotDTO) UnmarshalBinary(b []byte) error {
	var res DashboardSnapshotDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
