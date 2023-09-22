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

// Byoprofile byoprofile
//
// swagger:model byoprofile
type Byoprofile struct {

	// general
	General *ByoprofileGeneral `json:"general,omitempty"`
}

// Validate validates this byoprofile
func (m *Byoprofile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeneral(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Byoprofile) validateGeneral(formats strfmt.Registry) error {
	if swag.IsZero(m.General) { // not required
		return nil
	}

	if m.General != nil {
		if err := m.General.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("general")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("general")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this byoprofile based on the context it is used
func (m *Byoprofile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGeneral(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Byoprofile) contextValidateGeneral(ctx context.Context, formats strfmt.Registry) error {

	if m.General != nil {

		if swag.IsZero(m.General) { // not required
			return nil
		}

		if err := m.General.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("general")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("general")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Byoprofile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Byoprofile) UnmarshalBinary(b []byte) error {
	var res Byoprofile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ByoprofileGeneral byoprofile general
//
// swagger:model ByoprofileGeneral
type ByoprofileGeneral struct {

	// description
	// Example: Used for Android or iOS BYO device enrollments
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Name of the BYO profile
	// Example: Personal Device Profile
	// Required: true
	Name *string `json:"name"`

	// site
	Site *SiteObject `json:"site,omitempty"`
}

// Validate validates this byoprofile general
func (m *ByoprofileGeneral) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ByoprofileGeneral) validateName(formats strfmt.Registry) error {

	if err := validate.Required("general"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ByoprofileGeneral) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("general" + "." + "site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("general" + "." + "site")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this byoprofile general based on the context it is used
func (m *ByoprofileGeneral) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ByoprofileGeneral) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {

		if swag.IsZero(m.Site) { // not required
			return nil
		}

		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("general" + "." + "site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("general" + "." + "site")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ByoprofileGeneral) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ByoprofileGeneral) UnmarshalBinary(b []byte) error {
	var res ByoprofileGeneral
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}