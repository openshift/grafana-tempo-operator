// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// ErrorSource ErrorSource type defines the source of the error
//
// swagger:model ErrorSource
type ErrorSource string

// Validate validates this error source
func (m ErrorSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error source based on context it is used
func (m ErrorSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}