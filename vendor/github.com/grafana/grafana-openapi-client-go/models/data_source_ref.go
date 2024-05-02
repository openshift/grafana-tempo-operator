// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataSourceRef Ref to a DataSource instance
//
// swagger:model DataSourceRef
type DataSourceRef struct {

	// The plugin type-id
	Type string `json:"type,omitempty"`

	// Specific datasource instance
	UID string `json:"uid,omitempty"`
}

// Validate validates this data source ref
func (m *DataSourceRef) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this data source ref based on context it is used
func (m *DataSourceRef) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataSourceRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSourceRef) UnmarshalBinary(b []byte) error {
	var res DataSourceRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
