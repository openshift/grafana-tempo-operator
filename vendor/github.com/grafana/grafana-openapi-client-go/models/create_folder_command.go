// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateFolderCommand CreateFolderCommand captures the information required by the folder service
// to create a folder.
//
// swagger:model CreateFolderCommand
type CreateFolderCommand struct {

	// description
	Description string `json:"description,omitempty"`

	// parent Uid
	ParentUID string `json:"parentUid,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this create folder command
func (m *CreateFolderCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create folder command based on context it is used
func (m *CreateFolderCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateFolderCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateFolderCommand) UnmarshalBinary(b []byte) error {
	var res CreateFolderCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}