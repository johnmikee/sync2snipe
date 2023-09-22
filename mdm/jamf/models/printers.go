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

// Printers printers
//
// swagger:model printers
type Printers []*PrintersItems0

// Validate validates this printers
func (m Printers) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this printers based on the context it is used
func (m Printers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// PrintersItems0 printers items0
//
// swagger:model PrintersItems0
type PrintersItems0 struct {

	// printer
	Printer *IDName `json:"printer,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this printers items0
func (m *PrintersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrinter(formats); err != nil {
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

func (m *PrintersItems0) validatePrinter(formats strfmt.Registry) error {
	if swag.IsZero(m.Printer) { // not required
		return nil
	}

	if m.Printer != nil {
		if err := m.Printer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("printer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("printer")
			}
			return err
		}
	}

	return nil
}

func (m *PrintersItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this printers items0 based on the context it is used
func (m *PrintersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrinter(ctx, formats); err != nil {
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

func (m *PrintersItems0) contextValidatePrinter(ctx context.Context, formats strfmt.Registry) error {

	if m.Printer != nil {

		if swag.IsZero(m.Printer) { // not required
			return nil
		}

		if err := m.Printer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("printer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("printer")
			}
			return err
		}
	}

	return nil
}

func (m *PrintersItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PrintersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrintersItems0) UnmarshalBinary(b []byte) error {
	var res PrintersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
