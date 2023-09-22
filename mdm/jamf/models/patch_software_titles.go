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

// PatchSoftwareTitles patch software titles
//
// swagger:model patch_software_titles
type PatchSoftwareTitles []*PatchSoftwareTitlesItems0

// Validate validates this patch software titles
func (m PatchSoftwareTitles) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this patch software titles based on the context it is used
func (m PatchSoftwareTitles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// PatchSoftwareTitlesItems0 patch software titles items0
//
// swagger:model PatchSoftwareTitlesItems0
type PatchSoftwareTitlesItems0 struct {

	// patch software title
	PatchSoftwareTitle *PatchSoftwareTitlesItems0PatchSoftwareTitle `json:"patch_software_title,omitempty"`

	// size
	// Example: 1
	Size int64 `json:"size,omitempty"`
}

// Validate validates this patch software titles items0
func (m *PatchSoftwareTitlesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePatchSoftwareTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchSoftwareTitlesItems0) validatePatchSoftwareTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.PatchSoftwareTitle) { // not required
		return nil
	}

	if m.PatchSoftwareTitle != nil {
		if err := m.PatchSoftwareTitle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patch_software_title")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patch_software_title")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch software titles items0 based on the context it is used
func (m *PatchSoftwareTitlesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePatchSoftwareTitle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchSoftwareTitlesItems0) contextValidatePatchSoftwareTitle(ctx context.Context, formats strfmt.Registry) error {

	if m.PatchSoftwareTitle != nil {

		if swag.IsZero(m.PatchSoftwareTitle) { // not required
			return nil
		}

		if err := m.PatchSoftwareTitle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patch_software_title")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patch_software_title")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchSoftwareTitlesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchSoftwareTitlesItems0) UnmarshalBinary(b []byte) error {
	var res PatchSoftwareTitlesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchSoftwareTitlesItems0PatchSoftwareTitle patch software titles items0 patch software title
//
// swagger:model PatchSoftwareTitlesItems0PatchSoftwareTitle
type PatchSoftwareTitlesItems0PatchSoftwareTitle struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// name id
	// Example: GoogleChromne
	NameID string `json:"name_id,omitempty"`

	// source id
	// Example: 1
	SourceID int64 `json:"source_id,omitempty"`
}

// Validate validates this patch software titles items0 patch software title
func (m *PatchSoftwareTitlesItems0PatchSoftwareTitle) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch software titles items0 patch software title based on context it is used
func (m *PatchSoftwareTitlesItems0PatchSoftwareTitle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchSoftwareTitlesItems0PatchSoftwareTitle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchSoftwareTitlesItems0PatchSoftwareTitle) UnmarshalBinary(b []byte) error {
	var res PatchSoftwareTitlesItems0PatchSoftwareTitle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}