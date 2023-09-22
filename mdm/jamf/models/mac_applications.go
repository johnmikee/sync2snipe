// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MacApplications mac applications
//
// swagger:model mac_applications
type MacApplications []*MacApplicationsItems0

// Validate validates this mac applications
func (m MacApplications) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this mac applications based on the context it is used
func (m MacApplications) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if m[i] != nil {

			if swag.IsZero(m[i]) { // not required
				return nil
			}

			if err := m[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MacApplicationsItems0 mac applications items0
//
// swagger:model MacApplicationsItems0
type MacApplicationsItems0 struct {

	// mac application
	MacApplication *MacApplicationsItems0MacApplication `json:"mac_application,omitempty"`
}

// Validate validates this mac applications items0
func (m *MacApplicationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMacApplication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MacApplicationsItems0) validateMacApplication(formats strfmt.Registry) error {
	if swag.IsZero(m.MacApplication) { // not required
		return nil
	}

	if m.MacApplication != nil {
		if err := m.MacApplication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mac_application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mac_application")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mac applications items0 based on the context it is used
func (m *MacApplicationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMacApplication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MacApplicationsItems0) contextValidateMacApplication(ctx context.Context, formats strfmt.Registry) error {

	if m.MacApplication != nil {

		if swag.IsZero(m.MacApplication) { // not required
			return nil
		}

		if err := m.MacApplication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mac_application")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mac_application")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MacApplicationsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MacApplicationsItems0) UnmarshalBinary(b []byte) error {
	var res MacApplicationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MacApplicationsItems0MacApplication mac applications items0 mac application
//
// swagger:model MacApplicationsItems0MacApplication
type MacApplicationsItems0MacApplication struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Name of the application
	// Example: TextWrangler.app
	Name string `json:"name,omitempty"`
}

// Validate validates this mac applications items0 mac application
func (m *MacApplicationsItems0MacApplication) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mac applications items0 mac application based on context it is used
func (m *MacApplicationsItems0MacApplication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MacApplicationsItems0MacApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MacApplicationsItems0MacApplication) UnmarshalBinary(b []byte) error {
	var res MacApplicationsItems0MacApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}