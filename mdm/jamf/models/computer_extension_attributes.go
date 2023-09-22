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

// ComputerExtensionAttributes computer extension attributes
//
// swagger:model computer_extension_attributes
type ComputerExtensionAttributes []*ComputerExtensionAttributesItems0

// Validate validates this computer extension attributes
func (m ComputerExtensionAttributes) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this computer extension attributes based on the context it is used
func (m ComputerExtensionAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// ComputerExtensionAttributesItems0 computer extension attributes items0
//
// swagger:model ComputerExtensionAttributesItems0
type ComputerExtensionAttributesItems0 struct {

	// computer extension attribute
	ComputerExtensionAttribute *ComputerExtensionAttributesItems0ComputerExtensionAttribute `json:"computer_extension_attribute,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this computer extension attributes items0
func (m *ComputerExtensionAttributesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputerExtensionAttribute(formats); err != nil {
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

func (m *ComputerExtensionAttributesItems0) validateComputerExtensionAttribute(formats strfmt.Registry) error {
	if swag.IsZero(m.ComputerExtensionAttribute) { // not required
		return nil
	}

	if m.ComputerExtensionAttribute != nil {
		if err := m.ComputerExtensionAttribute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer_extension_attribute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computer_extension_attribute")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerExtensionAttributesItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this computer extension attributes items0 based on the context it is used
func (m *ComputerExtensionAttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputerExtensionAttribute(ctx, formats); err != nil {
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

func (m *ComputerExtensionAttributesItems0) contextValidateComputerExtensionAttribute(ctx context.Context, formats strfmt.Registry) error {

	if m.ComputerExtensionAttribute != nil {

		if swag.IsZero(m.ComputerExtensionAttribute) { // not required
			return nil
		}

		if err := m.ComputerExtensionAttribute.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer_extension_attribute")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computer_extension_attribute")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerExtensionAttributesItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ComputerExtensionAttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerExtensionAttributesItems0) UnmarshalBinary(b []byte) error {
	var res ComputerExtensionAttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerExtensionAttributesItems0ComputerExtensionAttribute computer extension attributes items0 computer extension attribute
//
// swagger:model ComputerExtensionAttributesItems0ComputerExtensionAttribute
type ComputerExtensionAttributesItems0ComputerExtensionAttribute struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this computer extension attributes items0 computer extension attribute
func (m *ComputerExtensionAttributesItems0ComputerExtensionAttribute) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this computer extension attributes items0 computer extension attribute based on context it is used
func (m *ComputerExtensionAttributesItems0ComputerExtensionAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerExtensionAttributesItems0ComputerExtensionAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerExtensionAttributesItems0ComputerExtensionAttribute) UnmarshalBinary(b []byte) error {
	var res ComputerExtensionAttributesItems0ComputerExtensionAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
