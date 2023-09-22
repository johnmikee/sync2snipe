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

// SoftwareUpdateServer software update server
//
// swagger:model software_update_server
type SoftwareUpdateServer struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// ip address
	// Example: 10.10.51.248
	// Required: true
	IPAddress *string `json:"ip_address"`

	// Name of the software update server
	// Example: New York SUS
	// Required: true
	Name *string `json:"name"`

	// port
	// Example: 8088
	Port int64 `json:"port,omitempty"`

	// set system wide
	SetSystemWide bool `json:"set_system_wide,omitempty"`
}

// Validate validates this software update server
func (m *SoftwareUpdateServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareUpdateServer) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ip_address", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *SoftwareUpdateServer) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this software update server based on context it is used
func (m *SoftwareUpdateServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareUpdateServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareUpdateServer) UnmarshalBinary(b []byte) error {
	var res SoftwareUpdateServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
