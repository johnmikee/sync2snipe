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

// OsxConfigurationProfiles os x configuration profiles
//
// swagger:model os_x_configuration_profiles
type OsxConfigurationProfiles []*OsxConfigurationProfilesItems0

// Validate validates this os x configuration profiles
func (m OsxConfigurationProfiles) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this os x configuration profiles based on the context it is used
func (m OsxConfigurationProfiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// OsxConfigurationProfilesItems0 osx configuration profiles items0
//
// swagger:model OsxConfigurationProfilesItems0
type OsxConfigurationProfilesItems0 struct {

	// os x configuration profile
	OsxConfigurationProfile *IDName `json:"os_x_configuration_profile,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this osx configuration profiles items0
func (m *OsxConfigurationProfilesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOsxConfigurationProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OsxConfigurationProfilesItems0) validateOsxConfigurationProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.OsxConfigurationProfile) { // not required
		return nil
	}

	if m.OsxConfigurationProfile != nil {
		if err := m.OsxConfigurationProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os_x_configuration_profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("os_x_configuration_profile")
			}
			return err
		}
	}

	return nil
}

func (m *OsxConfigurationProfilesItems0) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if err := m.Size.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("size")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("size")
		}
		return err
	}

	return nil
}

// ContextValidate validate this osx configuration profiles items0 based on the context it is used
func (m *OsxConfigurationProfilesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOsxConfigurationProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OsxConfigurationProfilesItems0) contextValidateOsxConfigurationProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.OsxConfigurationProfile != nil {

		if swag.IsZero(m.OsxConfigurationProfile) { // not required
			return nil
		}

		if err := m.OsxConfigurationProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os_x_configuration_profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("os_x_configuration_profile")
			}
			return err
		}
	}

	return nil
}

func (m *OsxConfigurationProfilesItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Size) { // not required
		return nil
	}

	if err := m.Size.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("size")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("size")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OsxConfigurationProfilesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OsxConfigurationProfilesItems0) UnmarshalBinary(b []byte) error {
	var res OsxConfigurationProfilesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}