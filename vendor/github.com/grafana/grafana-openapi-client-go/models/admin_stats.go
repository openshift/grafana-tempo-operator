// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminStats admin stats
//
// swagger:model AdminStats
type AdminStats struct {

	// active admins
	ActiveAdmins int64 `json:"activeAdmins,omitempty"`

	// active devices
	ActiveDevices int64 `json:"activeDevices,omitempty"`

	// active editors
	ActiveEditors int64 `json:"activeEditors,omitempty"`

	// active sessions
	ActiveSessions int64 `json:"activeSessions,omitempty"`

	// active users
	ActiveUsers int64 `json:"activeUsers,omitempty"`

	// active viewers
	ActiveViewers int64 `json:"activeViewers,omitempty"`

	// admins
	Admins int64 `json:"admins,omitempty"`

	// alerts
	Alerts int64 `json:"alerts,omitempty"`

	// daily active admins
	DailyActiveAdmins int64 `json:"dailyActiveAdmins,omitempty"`

	// daily active editors
	DailyActiveEditors int64 `json:"dailyActiveEditors,omitempty"`

	// daily active sessions
	DailyActiveSessions int64 `json:"dailyActiveSessions,omitempty"`

	// daily active users
	DailyActiveUsers int64 `json:"dailyActiveUsers,omitempty"`

	// daily active viewers
	DailyActiveViewers int64 `json:"dailyActiveViewers,omitempty"`

	// dashboards
	Dashboards int64 `json:"dashboards,omitempty"`

	// datasources
	Datasources int64 `json:"datasources,omitempty"`

	// editors
	Editors int64 `json:"editors,omitempty"`

	// monthly active users
	MonthlyActiveUsers int64 `json:"monthlyActiveUsers,omitempty"`

	// orgs
	Orgs int64 `json:"orgs,omitempty"`

	// playlists
	Playlists int64 `json:"playlists,omitempty"`

	// snapshots
	Snapshots int64 `json:"snapshots,omitempty"`

	// stars
	Stars int64 `json:"stars,omitempty"`

	// tags
	Tags int64 `json:"tags,omitempty"`

	// users
	Users int64 `json:"users,omitempty"`

	// viewers
	Viewers int64 `json:"viewers,omitempty"`
}

// Validate validates this admin stats
func (m *AdminStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this admin stats based on context it is used
func (m *AdminStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminStats) UnmarshalBinary(b []byte) error {
	var res AdminStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
