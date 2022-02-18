// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Event Event
//
// swagger:model Event
type Event struct {

	// meta
	// Required: true
	Meta Meta `json:"meta"`

	// timestamp
	// Required: true
	Timestamp *string `json:"timestamp"`
}

// Validate validates this event
func (m *Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) validateMeta(formats strfmt.Registry) error {

	if err := validate.Required("meta", "body", m.Meta); err != nil {
		return err
	}

	if err := m.Meta.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("meta")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("meta")
		}
		return err
	}

	return nil
}

func (m *Event) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this event based on the context it is used
func (m *Event) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Event) contextValidateMeta(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Meta.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("meta")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("meta")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Event) UnmarshalBinary(b []byte) error {
	var res Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}