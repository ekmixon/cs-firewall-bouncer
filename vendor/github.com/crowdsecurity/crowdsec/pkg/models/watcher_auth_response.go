// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WatcherAuthResponse WatcherAuthResponse
//
// the response of a successful authentication
//
// swagger:model WatcherAuthResponse
type WatcherAuthResponse struct {

	// code
	Code int64 `json:"code,omitempty"`

	// expire
	Expire string `json:"expire,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}

// Validate validates this watcher auth response
func (m *WatcherAuthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this watcher auth response based on context it is used
func (m *WatcherAuthResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WatcherAuthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WatcherAuthResponse) UnmarshalBinary(b []byte) error {
	var res WatcherAuthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}