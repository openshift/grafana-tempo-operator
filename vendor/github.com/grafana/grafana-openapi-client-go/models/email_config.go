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

// EmailConfig EmailConfig configures notifications via mail.
//
// swagger:model EmailConfig
type EmailConfig struct {

	// auth identity
	AuthIdentity string `json:"auth_identity,omitempty"`

	// auth password
	AuthPassword Secret `json:"auth_password,omitempty"`

	// auth password file
	AuthPasswordFile string `json:"auth_password_file,omitempty"`

	// auth secret
	AuthSecret Secret `json:"auth_secret,omitempty"`

	// auth username
	AuthUsername string `json:"auth_username,omitempty"`

	// from
	From string `json:"from,omitempty"`

	// headers
	Headers map[string]string `json:"headers,omitempty"`

	// hello
	Hello string `json:"hello,omitempty"`

	// html
	HTML string `json:"html,omitempty"`

	// require tls
	RequireTLS bool `json:"require_tls,omitempty"`

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// smarthost
	Smarthost *HostPort `json:"smarthost,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// tls config
	TLSConfig *TLSConfig `json:"tls_config,omitempty"`

	// Email address to notify.
	To string `json:"to,omitempty"`
}

// Validate validates this email config
func (m *EmailConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthSecret(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmarthost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLSConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailConfig) validateAuthPassword(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthPassword) { // not required
		return nil
	}

	if err := m.AuthPassword.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("auth_password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("auth_password")
		}
		return err
	}

	return nil
}

func (m *EmailConfig) validateAuthSecret(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthSecret) { // not required
		return nil
	}

	if err := m.AuthSecret.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("auth_secret")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("auth_secret")
		}
		return err
	}

	return nil
}

func (m *EmailConfig) validateSmarthost(formats strfmt.Registry) error {
	if swag.IsZero(m.Smarthost) { // not required
		return nil
	}

	if m.Smarthost != nil {
		if err := m.Smarthost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smarthost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smarthost")
			}
			return err
		}
	}

	return nil
}

func (m *EmailConfig) validateTLSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.TLSConfig) { // not required
		return nil
	}

	if m.TLSConfig != nil {
		if err := m.TLSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this email config based on the context it is used
func (m *EmailConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthPassword(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAuthSecret(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSmarthost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailConfig) contextValidateAuthPassword(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AuthPassword) { // not required
		return nil
	}

	if err := m.AuthPassword.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("auth_password")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("auth_password")
		}
		return err
	}

	return nil
}

func (m *EmailConfig) contextValidateAuthSecret(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.AuthSecret) { // not required
		return nil
	}

	if err := m.AuthSecret.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("auth_secret")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("auth_secret")
		}
		return err
	}

	return nil
}

func (m *EmailConfig) contextValidateSmarthost(ctx context.Context, formats strfmt.Registry) error {

	if m.Smarthost != nil {

		if swag.IsZero(m.Smarthost) { // not required
			return nil
		}

		if err := m.Smarthost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smarthost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smarthost")
			}
			return err
		}
	}

	return nil
}

func (m *EmailConfig) contextValidateTLSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.TLSConfig != nil {

		if swag.IsZero(m.TLSConfig) { // not required
			return nil
		}

		if err := m.TLSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmailConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailConfig) UnmarshalBinary(b []byte) error {
	var res EmailConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}