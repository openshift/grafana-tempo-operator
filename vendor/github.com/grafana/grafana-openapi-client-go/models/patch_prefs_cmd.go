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

// PatchPrefsCmd patch prefs cmd
//
// swagger:model PatchPrefsCmd
type PatchPrefsCmd struct {

	// cookies
	Cookies []CookieType `json:"cookies"`

	// The numerical :id of a favorited dashboard
	HomeDashboardID int64 `json:"homeDashboardId,omitempty"`

	// home dashboard UID
	HomeDashboardUID string `json:"homeDashboardUID,omitempty"`

	// language
	Language string `json:"language,omitempty"`

	// query history
	QueryHistory *QueryHistoryPreference `json:"queryHistory,omitempty"`

	// theme
	// Enum: [light dark]
	Theme string `json:"theme,omitempty"`

	// timezone
	// Enum: [utc browser]
	Timezone string `json:"timezone,omitempty"`

	// week start
	WeekStart string `json:"weekStart,omitempty"`
}

// Validate validates this patch prefs cmd
func (m *PatchPrefsCmd) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCookies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchPrefsCmd) validateCookies(formats strfmt.Registry) error {
	if swag.IsZero(m.Cookies) { // not required
		return nil
	}

	for i := 0; i < len(m.Cookies); i++ {

		if err := m.Cookies[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cookies" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cookies" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PatchPrefsCmd) validateQueryHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryHistory) { // not required
		return nil
	}

	if m.QueryHistory != nil {
		if err := m.QueryHistory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryHistory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryHistory")
			}
			return err
		}
	}

	return nil
}

var patchPrefsCmdTypeThemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["light","dark"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchPrefsCmdTypeThemePropEnum = append(patchPrefsCmdTypeThemePropEnum, v)
	}
}

const (

	// PatchPrefsCmdThemeLight captures enum value "light"
	PatchPrefsCmdThemeLight string = "light"

	// PatchPrefsCmdThemeDark captures enum value "dark"
	PatchPrefsCmdThemeDark string = "dark"
)

// prop value enum
func (m *PatchPrefsCmd) validateThemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchPrefsCmdTypeThemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchPrefsCmd) validateTheme(formats strfmt.Registry) error {
	if swag.IsZero(m.Theme) { // not required
		return nil
	}

	// value enum
	if err := m.validateThemeEnum("theme", "body", m.Theme); err != nil {
		return err
	}

	return nil
}

var patchPrefsCmdTypeTimezonePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["utc","browser"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchPrefsCmdTypeTimezonePropEnum = append(patchPrefsCmdTypeTimezonePropEnum, v)
	}
}

const (

	// PatchPrefsCmdTimezoneUtc captures enum value "utc"
	PatchPrefsCmdTimezoneUtc string = "utc"

	// PatchPrefsCmdTimezoneBrowser captures enum value "browser"
	PatchPrefsCmdTimezoneBrowser string = "browser"
)

// prop value enum
func (m *PatchPrefsCmd) validateTimezoneEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchPrefsCmdTypeTimezonePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchPrefsCmd) validateTimezone(formats strfmt.Registry) error {
	if swag.IsZero(m.Timezone) { // not required
		return nil
	}

	// value enum
	if err := m.validateTimezoneEnum("timezone", "body", m.Timezone); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this patch prefs cmd based on the context it is used
func (m *PatchPrefsCmd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCookies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueryHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchPrefsCmd) contextValidateCookies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cookies); i++ {

		if swag.IsZero(m.Cookies[i]) { // not required
			return nil
		}

		if err := m.Cookies[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cookies" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cookies" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PatchPrefsCmd) contextValidateQueryHistory(ctx context.Context, formats strfmt.Registry) error {

	if m.QueryHistory != nil {

		if swag.IsZero(m.QueryHistory) { // not required
			return nil
		}

		if err := m.QueryHistory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryHistory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryHistory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchPrefsCmd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchPrefsCmd) UnmarshalBinary(b []byte) error {
	var res PatchPrefsCmd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
