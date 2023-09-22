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

// RemovableMacAddresses removable mac addresses
//
// swagger:model removable_mac_addresses
type RemovableMacAddresses []*RemovableMacAddressesItems0

// Validate validates this removable mac addresses
func (m RemovableMacAddresses) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this removable mac addresses based on the context it is used
func (m RemovableMacAddresses) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// RemovableMacAddressesItems0 removable mac addresses items0
//
// swagger:model RemovableMacAddressesItems0
type RemovableMacAddressesItems0 struct {

	// removable mac address
	RemovableMacAddress *RemovableMacAddress `json:"removable_mac_address,omitempty"`
}

// Validate validates this removable mac addresses items0
func (m *RemovableMacAddressesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemovableMacAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemovableMacAddressesItems0) validateRemovableMacAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.RemovableMacAddress) { // not required
		return nil
	}

	if m.RemovableMacAddress != nil {
		if err := m.RemovableMacAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("removable_mac_address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("removable_mac_address")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this removable mac addresses items0 based on the context it is used
func (m *RemovableMacAddressesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRemovableMacAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemovableMacAddressesItems0) contextValidateRemovableMacAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.RemovableMacAddress != nil {

		if swag.IsZero(m.RemovableMacAddress) { // not required
			return nil
		}

		if err := m.RemovableMacAddress.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("removable_mac_address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("removable_mac_address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemovableMacAddressesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemovableMacAddressesItems0) UnmarshalBinary(b []byte) error {
	var res RemovableMacAddressesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
