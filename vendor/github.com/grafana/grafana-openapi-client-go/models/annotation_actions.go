// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnnotationActions annotation actions
//
// swagger:model AnnotationActions
type AnnotationActions struct {

	// can add
	CanAdd bool `json:"canAdd,omitempty"`

	// can delete
	CanDelete bool `json:"canDelete,omitempty"`

	// can edit
	CanEdit bool `json:"canEdit,omitempty"`
}

// Validate validates this annotation actions
func (m *AnnotationActions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this annotation actions based on context it is used
func (m *AnnotationActions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnnotationActions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnnotationActions) UnmarshalBinary(b []byte) error {
	var res AnnotationActions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
