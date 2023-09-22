// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ComputerPost computer post
//
// swagger:model computer_post
type ComputerPost struct {

	// general
	General *ComputerPostGeneral `json:"general,omitempty"`
}

// Validate validates this computer post
func (m *ComputerPost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeneral(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerPost) validateGeneral(formats strfmt.Registry) error {
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

// ContextValidate validate this computer post based on the context it is used
func (m *ComputerPost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGeneral(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerPost) contextValidateGeneral(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ComputerPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerPost) UnmarshalBinary(b []byte) error {
	var res ComputerPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerPostGeneral computer post general
//
// swagger:model ComputerPostGeneral
type ComputerPostGeneral struct {

	// building
	// Example: East
	Building string `json:"building,omitempty"`

	// department
	// Example: Accounting
	Department string `json:"department,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// mac address
	// Example: E0:AC:CB:97:36:G4
	MacAddress string `json:"mac_address,omitempty"`

	// managed
	Managed bool `json:"managed,omitempty"`

	// model
	// Example: 13-inch MacBook Pro (Mid 2016)
	Model string `json:"model,omitempty"`

	// Name of the computer
	// Example: Lauras MacBook Pro
	Name string `json:"name,omitempty"`

	// report date epoch
	// Example: 1499470624555
	ReportDateEpoch int64 `json:"report_date_epoch,omitempty"`

	// report date utc
	// Example: 2017-07-07T18:37:04.555-0500
	ReportDateUtc string `json:"report_date_utc,omitempty"`

	// serial number
	// Example: C02Q7KHTGFWF
	SerialNumber string `json:"serial_number,omitempty"`

	// udid
	// Example: 55900BDC-347C-58B1-D249-F32244B11D30
	Udid string `json:"udid,omitempty"`

	// username
	// Example: Laura
	Username string `json:"username,omitempty"`
}

// Validate validates this computer post general
func (m *ComputerPostGeneral) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this computer post general based on context it is used
func (m *ComputerPostGeneral) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerPostGeneral) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerPostGeneral) UnmarshalBinary(b []byte) error {
	var res ComputerPostGeneral
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}