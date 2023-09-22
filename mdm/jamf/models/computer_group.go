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

// ComputerGroup computer group
//
// swagger:model computer_group
type ComputerGroup struct {

	// computers
	Computers []*ComputerGroupComputersItems0 `json:"computers"`

	// criteria
	Criteria []*ComputerGroupCriteriaItems0 `json:"criteria"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Smart or static group
	IsSmart bool `json:"is_smart,omitempty"`

	// Name of the group
	// Example: Group Name
	Name string `json:"name,omitempty"`

	// site
	Site *SiteObject `json:"site,omitempty"`
}

// Validate validates this computer group
func (m *ComputerGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCriteria(formats); err != nil {
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

func (m *ComputerGroup) validateComputers(formats strfmt.Registry) error {
	if swag.IsZero(m.Computers) { // not required
		return nil
	}

	for i := 0; i < len(m.Computers); i++ {
		if swag.IsZero(m.Computers[i]) { // not required
			continue
		}

		if m.Computers[i] != nil {
			if err := m.Computers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("computers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("computers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputerGroup) validateCriteria(formats strfmt.Registry) error {
	if swag.IsZero(m.Criteria) { // not required
		return nil
	}

	for i := 0; i < len(m.Criteria); i++ {
		if swag.IsZero(m.Criteria[i]) { // not required
			continue
		}

		if m.Criteria[i] != nil {
			if err := m.Criteria[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("criteria" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputerGroup) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this computer group based on the context it is used
func (m *ComputerGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputerGroup) contextValidateComputers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Computers); i++ {

		if m.Computers[i] != nil {

			if swag.IsZero(m.Computers[i]) { // not required
				return nil
			}

			if err := m.Computers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("computers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("computers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputerGroup) contextValidateCriteria(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Criteria); i++ {

		if m.Criteria[i] != nil {

			if swag.IsZero(m.Criteria[i]) { // not required
				return nil
			}

			if err := m.Criteria[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("criteria" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputerGroup) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {

		if swag.IsZero(m.Site) { // not required
			return nil
		}

		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputerGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerGroup) UnmarshalBinary(b []byte) error {
	var res ComputerGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerGroupComputersItems0 computer group computers items0
//
// swagger:model ComputerGroupComputersItems0
type ComputerGroupComputersItems0 struct {

	// computer
	Computer *ComputerGroupComputersItems0Computer `json:"computer,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this computer group computers items0
func (m *ComputerGroupComputersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputer(formats); err != nil {
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

func (m *ComputerGroupComputersItems0) validateComputer(formats strfmt.Registry) error {
	if swag.IsZero(m.Computer) { // not required
		return nil
	}

	if m.Computer != nil {
		if err := m.Computer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computer")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerGroupComputersItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this computer group computers items0 based on the context it is used
func (m *ComputerGroupComputersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputer(ctx, formats); err != nil {
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

func (m *ComputerGroupComputersItems0) contextValidateComputer(ctx context.Context, formats strfmt.Registry) error {

	if m.Computer != nil {

		if swag.IsZero(m.Computer) { // not required
			return nil
		}

		if err := m.Computer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("computer")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerGroupComputersItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ComputerGroupComputersItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerGroupComputersItems0) UnmarshalBinary(b []byte) error {
	var res ComputerGroupComputersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerGroupComputersItems0Computer computer group computers items0 computer
//
// swagger:model ComputerGroupComputersItems0Computer
type ComputerGroupComputersItems0Computer struct {

	// alt mac address
	// Example: E0:AC:CB:67:36:T4
	AltMacAddress string `json:"alt_mac_address,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// mac address
	// Example: E0:AC:CB:97:36:G4
	MacAddress string `json:"mac_address,omitempty"`

	// Name of the computer
	// Example: Joes iMac
	Name string `json:"name,omitempty"`

	// serial number
	// Example: C02Q7KHTGFWF
	SerialNumber string `json:"serial_number,omitempty"`
}

// Validate validates this computer group computers items0 computer
func (m *ComputerGroupComputersItems0Computer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this computer group computers items0 computer based on context it is used
func (m *ComputerGroupComputersItems0Computer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerGroupComputersItems0Computer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerGroupComputersItems0Computer) UnmarshalBinary(b []byte) error {
	var res ComputerGroupComputersItems0Computer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerGroupCriteriaItems0 computer group criteria items0
//
// swagger:model ComputerGroupCriteriaItems0
type ComputerGroupCriteriaItems0 struct {

	// criterion
	Criterion *Criterion `json:"criterion,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this computer group criteria items0
func (m *ComputerGroupCriteriaItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCriterion(formats); err != nil {
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

func (m *ComputerGroupCriteriaItems0) validateCriterion(formats strfmt.Registry) error {
	if swag.IsZero(m.Criterion) { // not required
		return nil
	}

	if m.Criterion != nil {
		if err := m.Criterion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("criterion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("criterion")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerGroupCriteriaItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this computer group criteria items0 based on the context it is used
func (m *ComputerGroupCriteriaItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCriterion(ctx, formats); err != nil {
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

func (m *ComputerGroupCriteriaItems0) contextValidateCriterion(ctx context.Context, formats strfmt.Registry) error {

	if m.Criterion != nil {

		if swag.IsZero(m.Criterion) { // not required
			return nil
		}

		if err := m.Criterion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("criterion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("criterion")
			}
			return err
		}
	}

	return nil
}

func (m *ComputerGroupCriteriaItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ComputerGroupCriteriaItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerGroupCriteriaItems0) UnmarshalBinary(b []byte) error {
	var res ComputerGroupCriteriaItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
