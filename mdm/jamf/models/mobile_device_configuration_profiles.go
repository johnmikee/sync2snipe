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

// MobileDeviceConfigurationProfiles mobile device configuration profiles
//
// swagger:model mobile_device_configuration_profiles
type MobileDeviceConfigurationProfiles []*MobileDeviceConfigurationProfilesItems0

// Validate validates this mobile device configuration profiles
func (m MobileDeviceConfigurationProfiles) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this mobile device configuration profiles based on the context it is used
func (m MobileDeviceConfigurationProfiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// MobileDeviceConfigurationProfilesItems0 mobile device configuration profiles items0
//
// swagger:model MobileDeviceConfigurationProfilesItems0
type MobileDeviceConfigurationProfilesItems0 struct {

	// configuration profile
	ConfigurationProfile *MobileDeviceConfigurationProfilesItems0ConfigurationProfile `json:"configuration_profile,omitempty"`
}

// Validate validates this mobile device configuration profiles items0
func (m *MobileDeviceConfigurationProfilesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigurationProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MobileDeviceConfigurationProfilesItems0) validateConfigurationProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigurationProfile) { // not required
		return nil
	}

	if m.ConfigurationProfile != nil {
		if err := m.ConfigurationProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration_profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration_profile")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mobile device configuration profiles items0 based on the context it is used
func (m *MobileDeviceConfigurationProfilesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigurationProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MobileDeviceConfigurationProfilesItems0) contextValidateConfigurationProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.ConfigurationProfile != nil {

		if swag.IsZero(m.ConfigurationProfile) { // not required
			return nil
		}

		if err := m.ConfigurationProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration_profile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration_profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MobileDeviceConfigurationProfilesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileDeviceConfigurationProfilesItems0) UnmarshalBinary(b []byte) error {
	var res MobileDeviceConfigurationProfilesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MobileDeviceConfigurationProfilesItems0ConfigurationProfile mobile device configuration profiles items0 configuration profile
//
// swagger:model MobileDeviceConfigurationProfilesItems0ConfigurationProfile
type MobileDeviceConfigurationProfilesItems0ConfigurationProfile struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Name of the configuration profile
	// Example: Corporate Wireless
	Name string `json:"name,omitempty"`
}

// Validate validates this mobile device configuration profiles items0 configuration profile
func (m *MobileDeviceConfigurationProfilesItems0ConfigurationProfile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mobile device configuration profiles items0 configuration profile based on context it is used
func (m *MobileDeviceConfigurationProfilesItems0ConfigurationProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MobileDeviceConfigurationProfilesItems0ConfigurationProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileDeviceConfigurationProfilesItems0ConfigurationProfile) UnmarshalBinary(b []byte) error {
	var res MobileDeviceConfigurationProfilesItems0ConfigurationProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
