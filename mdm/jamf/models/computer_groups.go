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

// ComputerGroups computer groups
//
// swagger:model computer_groups
type ComputerGroups []*ComputerGroupsItems0

// Validate validates this computer groups
func (m ComputerGroups) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this computer groups based on the context it is used
func (m ComputerGroups) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// ComputerGroupsItems0 computer groups items0
//
// swagger:model ComputerGroupsItems0
type ComputerGroupsItems0 struct {

	// computer group
	ComputerGroup []*ComputerGroupsItems0ComputerGroupItems0 `json:"computer_group"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this computer groups items0
func (m *ComputerGroupsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputerGroup(formats); err != nil {
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

func (m *ComputerGroupsItems0) validateComputerGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ComputerGroup) { // not required
		return nil
	}

	for i := 0; i < len(m.ComputerGroup); i++ {
		if swag.IsZero(m.ComputerGroup[i]) { // not required
			continue
		}

		if m.ComputerGroup[i] != nil {
			if err := m.ComputerGroup[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("computer_group" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("computer_group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputerGroupsItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this computer groups items0 based on the context it is used
func (m *ComputerGroupsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputerGroup(ctx, formats); err != nil {
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

func (m *ComputerGroupsItems0) contextValidateComputerGroup(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ComputerGroup); i++ {

		if m.ComputerGroup[i] != nil {

			if swag.IsZero(m.ComputerGroup[i]) { // not required
				return nil
			}

			if err := m.ComputerGroup[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("computer_group" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("computer_group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputerGroupsItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *ComputerGroupsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerGroupsItems0) UnmarshalBinary(b []byte) error {
	var res ComputerGroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ComputerGroupsItems0ComputerGroupItems0 computer groups items0 computer group items0
//
// swagger:model ComputerGroupsItems0ComputerGroupItems0
type ComputerGroupsItems0ComputerGroupItems0 struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// is smart
	IsSmart bool `json:"is_smart,omitempty"`

	// Name of the group
	// Example: Group Name
	Name string `json:"name,omitempty"`
}

// Validate validates this computer groups items0 computer group items0
func (m *ComputerGroupsItems0ComputerGroupItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this computer groups items0 computer group items0 based on context it is used
func (m *ComputerGroupsItems0ComputerGroupItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ComputerGroupsItems0ComputerGroupItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputerGroupsItems0ComputerGroupItems0) UnmarshalBinary(b []byte) error {
	var res ComputerGroupsItems0ComputerGroupItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
