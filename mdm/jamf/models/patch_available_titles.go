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

// PatchAvailableTitles patch available titles
//
// swagger:model patch_available_titles
type PatchAvailableTitles struct {

	// available titles
	AvailableTitles []*PatchAvailableTitlesAvailableTitlesItems0 `json:"available_titles"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this patch available titles
func (m *PatchAvailableTitles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableTitles(formats); err != nil {
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

func (m *PatchAvailableTitles) validateAvailableTitles(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailableTitles) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableTitles); i++ {
		if swag.IsZero(m.AvailableTitles[i]) { // not required
			continue
		}

		if m.AvailableTitles[i] != nil {
			if err := m.AvailableTitles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("available_titles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("available_titles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchAvailableTitles) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this patch available titles based on the context it is used
func (m *PatchAvailableTitles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailableTitles(ctx, formats); err != nil {
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

func (m *PatchAvailableTitles) contextValidateAvailableTitles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailableTitles); i++ {

		if m.AvailableTitles[i] != nil {

			if swag.IsZero(m.AvailableTitles[i]) { // not required
				return nil
			}

			if err := m.AvailableTitles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("available_titles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("available_titles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PatchAvailableTitles) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PatchAvailableTitles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchAvailableTitles) UnmarshalBinary(b []byte) error {
	var res PatchAvailableTitles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchAvailableTitlesAvailableTitlesItems0 patch available titles available titles items0
//
// swagger:model PatchAvailableTitlesAvailableTitlesItems0
type PatchAvailableTitlesAvailableTitlesItems0 struct {

	// available title
	AvailableTitle *PatchAvailableTitlesAvailableTitlesItems0AvailableTitle `json:"available_title,omitempty"`
}

// Validate validates this patch available titles available titles items0
func (m *PatchAvailableTitlesAvailableTitlesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchAvailableTitlesAvailableTitlesItems0) validateAvailableTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailableTitle) { // not required
		return nil
	}

	if m.AvailableTitle != nil {
		if err := m.AvailableTitle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("available_title")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("available_title")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch available titles available titles items0 based on the context it is used
func (m *PatchAvailableTitlesAvailableTitlesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailableTitle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchAvailableTitlesAvailableTitlesItems0) contextValidateAvailableTitle(ctx context.Context, formats strfmt.Registry) error {

	if m.AvailableTitle != nil {

		if swag.IsZero(m.AvailableTitle) { // not required
			return nil
		}

		if err := m.AvailableTitle.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("available_title")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("available_title")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchAvailableTitlesAvailableTitlesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchAvailableTitlesAvailableTitlesItems0) UnmarshalBinary(b []byte) error {
	var res PatchAvailableTitlesAvailableTitlesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchAvailableTitlesAvailableTitlesItems0AvailableTitle patch available titles available titles items0 available title
//
// swagger:model PatchAvailableTitlesAvailableTitlesItems0AvailableTitle
type PatchAvailableTitlesAvailableTitlesItems0AvailableTitle struct {

	// app name
	// Example: Google Chrome
	AppName string `json:"app_name,omitempty"`

	// current version
	// Example: 65.0.3325.181
	CurrentVersion string `json:"current_version,omitempty"`

	// last modified
	// Example: 2018-03-20T22:00:09:000Z
	LastModified string `json:"last_modified,omitempty"`

	// name id
	// Example: GoogleChrome
	NameID string `json:"name_id,omitempty"`

	// publisher
	// Example: Google
	Publisher string `json:"publisher,omitempty"`
}

// Validate validates this patch available titles available titles items0 available title
func (m *PatchAvailableTitlesAvailableTitlesItems0AvailableTitle) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch available titles available titles items0 available title based on context it is used
func (m *PatchAvailableTitlesAvailableTitlesItems0AvailableTitle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchAvailableTitlesAvailableTitlesItems0AvailableTitle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchAvailableTitlesAvailableTitlesItems0AvailableTitle) UnmarshalBinary(b []byte) error {
	var res PatchAvailableTitlesAvailableTitlesItems0AvailableTitle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
