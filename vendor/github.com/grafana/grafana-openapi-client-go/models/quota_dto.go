// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuotaDTO quota DTO
//
// swagger:model QuotaDTO
type QuotaDTO struct {

	// limit
	Limit int64 `json:"limit,omitempty"`

	// org id
	OrgID int64 `json:"org_id,omitempty"`

	// target
	Target string `json:"target,omitempty"`

	// used
	Used int64 `json:"used,omitempty"`

	// user id
	UserID int64 `json:"user_id,omitempty"`
}

// Validate validates this quota DTO
func (m *QuotaDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this quota DTO based on context it is used
func (m *QuotaDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaDTO) UnmarshalBinary(b []byte) error {
	var res QuotaDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
