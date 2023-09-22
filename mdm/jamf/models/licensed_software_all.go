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

// LicensedSoftwareAll licensed software all
//
// swagger:model licensed_software_all
type LicensedSoftwareAll []*LicensedSoftwareAllItems0

// Validate validates this licensed software all
func (m LicensedSoftwareAll) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this licensed software all based on the context it is used
func (m LicensedSoftwareAll) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// LicensedSoftwareAllItems0 licensed software all items0
//
// swagger:model LicensedSoftwareAllItems0
type LicensedSoftwareAllItems0 struct {

	// licensed software
	LicensedSoftware *LicensedSoftwareAllItems0LicensedSoftware `json:"licensed_software,omitempty"`
}

// Validate validates this licensed software all items0
func (m *LicensedSoftwareAllItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicensedSoftware(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicensedSoftwareAllItems0) validateLicensedSoftware(formats strfmt.Registry) error {
	if swag.IsZero(m.LicensedSoftware) { // not required
		return nil
	}

	if m.LicensedSoftware != nil {
		if err := m.LicensedSoftware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licensed_software")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licensed_software")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this licensed software all items0 based on the context it is used
func (m *LicensedSoftwareAllItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLicensedSoftware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicensedSoftwareAllItems0) contextValidateLicensedSoftware(ctx context.Context, formats strfmt.Registry) error {

	if m.LicensedSoftware != nil {

		if swag.IsZero(m.LicensedSoftware) { // not required
			return nil
		}

		if err := m.LicensedSoftware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licensed_software")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licensed_software")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicensedSoftwareAllItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicensedSoftwareAllItems0) UnmarshalBinary(b []byte) error {
	var res LicensedSoftwareAllItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LicensedSoftwareAllItems0LicensedSoftware licensed software all items0 licensed software
//
// swagger:model LicensedSoftwareAllItems0LicensedSoftware
type LicensedSoftwareAllItems0LicensedSoftware struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Name of the licensed software
	// Example: Adobe Creative Suite
	Name string `json:"name,omitempty"`
}

// Validate validates this licensed software all items0 licensed software
func (m *LicensedSoftwareAllItems0LicensedSoftware) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this licensed software all items0 licensed software based on context it is used
func (m *LicensedSoftwareAllItems0LicensedSoftware) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LicensedSoftwareAllItems0LicensedSoftware) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicensedSoftwareAllItems0LicensedSoftware) UnmarshalBinary(b []byte) error {
	var res LicensedSoftwareAllItems0LicensedSoftware
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}