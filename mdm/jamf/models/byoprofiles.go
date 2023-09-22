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

// Byoprofiles byoprofiles
//
// swagger:model byoprofiles
type Byoprofiles []*ByoprofilesItems0

// Validate validates this byoprofiles
func (m Byoprofiles) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this byoprofiles based on the context it is used
func (m Byoprofiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// ByoprofilesItems0 byoprofiles items0
//
// swagger:model ByoprofilesItems0
type ByoprofilesItems0 struct {

	// byoprofile
	Byoprofile *IDName `json:"byoprofile,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this byoprofiles items0
func (m *ByoprofilesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateByoprofile(formats); err != nil {
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

func (m *ByoprofilesItems0) validateByoprofile(formats strfmt.Registry) error {
	if swag.IsZero(m.Byoprofile) { // not required
		return nil
	}

	if m.Byoprofile != nil {
		if err := m.Byoprofile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("byoprofile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("byoprofile")
			}
			return err
		}
	}

	return nil
}

func (m *ByoprofilesItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this byoprofiles items0 based on the context it is used
func (m *ByoprofilesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateByoprofile(ctx, formats); err != nil {
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

func (m *ByoprofilesItems0) contextValidateByoprofile(ctx context.Context, formats strfmt.Registry) error {

	if m.Byoprofile != nil {

		if swag.IsZero(m.Byoprofile) { // not required
			return nil
		}

		if err := m.Byoprofile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("byoprofile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("byoprofile")
			}
			return err
		}
	}

	return nil
}

func (m *ByoprofilesItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ByoprofilesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ByoprofilesItems0) UnmarshalBinary(b []byte) error {
	var res ByoprofilesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
