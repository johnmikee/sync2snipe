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

// PatchPolicies patch policies
//
// swagger:model patch_policies
type PatchPolicies []*PatchPoliciesItems0

// Validate validates this patch policies
func (m PatchPolicies) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this patch policies based on the context it is used
func (m PatchPolicies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// PatchPoliciesItems0 patch policies items0
//
// swagger:model PatchPoliciesItems0
type PatchPoliciesItems0 struct {

	// patch policy
	PatchPolicy *IDName `json:"patch_policy,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this patch policies items0
func (m *PatchPoliciesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatchPolicy(formats); err != nil {
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

func (m *PatchPoliciesItems0) validatePatchPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.PatchPolicy) { // not required
		return nil
	}

	if m.PatchPolicy != nil {
		if err := m.PatchPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patch_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patch_policy")
			}
			return err
		}
	}

	return nil
}

func (m *PatchPoliciesItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this patch policies items0 based on the context it is used
func (m *PatchPoliciesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePatchPolicy(ctx, formats); err != nil {
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

func (m *PatchPoliciesItems0) contextValidatePatchPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.PatchPolicy != nil {

		if swag.IsZero(m.PatchPolicy) { // not required
			return nil
		}

		if err := m.PatchPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patch_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patch_policy")
			}
			return err
		}
	}

	return nil
}

func (m *PatchPoliciesItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PatchPoliciesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchPoliciesItems0) UnmarshalBinary(b []byte) error {
	var res PatchPoliciesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
