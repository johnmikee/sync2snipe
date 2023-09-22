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

// DistributionPoints distribution points
//
// swagger:model distribution_points
type DistributionPoints []*DistributionPointsItems0

// Validate validates this distribution points
func (m DistributionPoints) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this distribution points based on the context it is used
func (m DistributionPoints) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// DistributionPointsItems0 distribution points items0
//
// swagger:model DistributionPointsItems0
type DistributionPointsItems0 struct {

	// distribution point
	DistributionPoint *IDName `json:"distribution_point,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this distribution points items0
func (m *DistributionPointsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistributionPoint(formats); err != nil {
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

func (m *DistributionPointsItems0) validateDistributionPoint(formats strfmt.Registry) error {
	if swag.IsZero(m.DistributionPoint) { // not required
		return nil
	}

	if m.DistributionPoint != nil {
		if err := m.DistributionPoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distribution_point")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("distribution_point")
			}
			return err
		}
	}

	return nil
}

func (m *DistributionPointsItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this distribution points items0 based on the context it is used
func (m *DistributionPointsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDistributionPoint(ctx, formats); err != nil {
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

func (m *DistributionPointsItems0) contextValidateDistributionPoint(ctx context.Context, formats strfmt.Registry) error {

	if m.DistributionPoint != nil {

		if swag.IsZero(m.DistributionPoint) { // not required
			return nil
		}

		if err := m.DistributionPoint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distribution_point")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("distribution_point")
			}
			return err
		}
	}

	return nil
}

func (m *DistributionPointsItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DistributionPointsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistributionPointsItems0) UnmarshalBinary(b []byte) error {
	var res DistributionPointsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}