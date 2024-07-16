// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimeIntervalRange time interval range
//
// swagger:model TimeIntervalRange
type TimeIntervalRange struct {

	// end time
	EndTime string `json:"end_time,omitempty"`

	// start time
	StartTime string `json:"start_time,omitempty"`
}

// Validate validates this time interval range
func (m *TimeIntervalRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this time interval range based on context it is used
func (m *TimeIntervalRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeIntervalRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeIntervalRange) UnmarshalBinary(b []byte) error {
	var res TimeIntervalRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}