// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostableGrafanaRule postable grafana rule
//
// swagger:model PostableGrafanaRule
type PostableGrafanaRule struct {

	// condition
	Condition string `json:"condition,omitempty"`

	// data
	Data []*AlertQuery `json:"data"`

	// exec err state
	// Enum: [OK Alerting Error]
	ExecErrState string `json:"exec_err_state,omitempty"`

	// is paused
	IsPaused bool `json:"is_paused,omitempty"`

	// no data state
	// Enum: [Alerting NoData OK]
	NoDataState string `json:"no_data_state,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this postable grafana rule
func (m *PostableGrafanaRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecErrState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoDataState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableGrafanaRule) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var postableGrafanaRuleTypeExecErrStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OK","Alerting","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postableGrafanaRuleTypeExecErrStatePropEnum = append(postableGrafanaRuleTypeExecErrStatePropEnum, v)
	}
}

const (

	// PostableGrafanaRuleExecErrStateOK captures enum value "OK"
	PostableGrafanaRuleExecErrStateOK string = "OK"

	// PostableGrafanaRuleExecErrStateAlerting captures enum value "Alerting"
	PostableGrafanaRuleExecErrStateAlerting string = "Alerting"

	// PostableGrafanaRuleExecErrStateError captures enum value "Error"
	PostableGrafanaRuleExecErrStateError string = "Error"
)

// prop value enum
func (m *PostableGrafanaRule) validateExecErrStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postableGrafanaRuleTypeExecErrStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostableGrafanaRule) validateExecErrState(formats strfmt.Registry) error {
	if swag.IsZero(m.ExecErrState) { // not required
		return nil
	}

	// value enum
	if err := m.validateExecErrStateEnum("exec_err_state", "body", m.ExecErrState); err != nil {
		return err
	}

	return nil
}

var postableGrafanaRuleTypeNoDataStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Alerting","NoData","OK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postableGrafanaRuleTypeNoDataStatePropEnum = append(postableGrafanaRuleTypeNoDataStatePropEnum, v)
	}
}

const (

	// PostableGrafanaRuleNoDataStateAlerting captures enum value "Alerting"
	PostableGrafanaRuleNoDataStateAlerting string = "Alerting"

	// PostableGrafanaRuleNoDataStateNoData captures enum value "NoData"
	PostableGrafanaRuleNoDataStateNoData string = "NoData"

	// PostableGrafanaRuleNoDataStateOK captures enum value "OK"
	PostableGrafanaRuleNoDataStateOK string = "OK"
)

// prop value enum
func (m *PostableGrafanaRule) validateNoDataStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postableGrafanaRuleTypeNoDataStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostableGrafanaRule) validateNoDataState(formats strfmt.Registry) error {
	if swag.IsZero(m.NoDataState) { // not required
		return nil
	}

	// value enum
	if err := m.validateNoDataStateEnum("no_data_state", "body", m.NoDataState); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this postable grafana rule based on the context it is used
func (m *PostableGrafanaRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostableGrafanaRule) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		if m.Data[i] != nil {

			if swag.IsZero(m.Data[i]) { // not required
				return nil
			}

			if err := m.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostableGrafanaRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostableGrafanaRule) UnmarshalBinary(b []byte) error {
	var res PostableGrafanaRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
