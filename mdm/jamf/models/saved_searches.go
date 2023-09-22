// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SavedSearches saved searches
//
// swagger:model saved_searches
type SavedSearches []*SavedSearchesItems0

// Validate validates this saved searches
func (m SavedSearches) Validate(formats strfmt.Registry) error {
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

// ContextValidate validate this saved searches based on the context it is used
func (m SavedSearches) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

// SavedSearchesItems0 saved searches items0
//
// swagger:model SavedSearchesItems0
type SavedSearchesItems0 struct {

	// saved search
	SavedSearch *SavedSearchesItems0SavedSearch `json:"saved_search,omitempty"`

	// size
	// Example: 1
	Size int64 `json:"size,omitempty"`
}

// Validate validates this saved searches items0
func (m *SavedSearchesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSavedSearch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SavedSearchesItems0) validateSavedSearch(formats strfmt.Registry) error {
	if swag.IsZero(m.SavedSearch) { // not required
		return nil
	}

	if m.SavedSearch != nil {
		if err := m.SavedSearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saved_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saved_search")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this saved searches items0 based on the context it is used
func (m *SavedSearchesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSavedSearch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SavedSearchesItems0) contextValidateSavedSearch(ctx context.Context, formats strfmt.Registry) error {

	if m.SavedSearch != nil {

		if swag.IsZero(m.SavedSearch) { // not required
			return nil
		}

		if err := m.SavedSearch.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("saved_search")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("saved_search")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SavedSearchesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SavedSearchesItems0) UnmarshalBinary(b []byte) error {
	var res SavedSearchesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SavedSearchesItems0SavedSearch saved searches items0 saved search
//
// swagger:model SavedSearchesItems0SavedSearch
type SavedSearchesItems0SavedSearch struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// name
	// Example: Advanced Computer Search Name
	Name string `json:"name,omitempty"`

	// type
	// Example: Computers
	// Enum: [Computers Mobile Devices Users]
	Type string `json:"type,omitempty"`
}

// Validate validates this saved searches items0 saved search
func (m *SavedSearchesItems0SavedSearch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var savedSearchesItems0SavedSearchTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Computers","Mobile Devices","Users"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		savedSearchesItems0SavedSearchTypeTypePropEnum = append(savedSearchesItems0SavedSearchTypeTypePropEnum, v)
	}
}

const (

	// SavedSearchesItems0SavedSearchTypeComputers captures enum value "Computers"
	SavedSearchesItems0SavedSearchTypeComputers string = "Computers"

	// SavedSearchesItems0SavedSearchTypeMobileDevices captures enum value "Mobile Devices"
	SavedSearchesItems0SavedSearchTypeMobileDevices string = "Mobile Devices"

	// SavedSearchesItems0SavedSearchTypeUsers captures enum value "Users"
	SavedSearchesItems0SavedSearchTypeUsers string = "Users"
)

// prop value enum
func (m *SavedSearchesItems0SavedSearch) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, savedSearchesItems0SavedSearchTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SavedSearchesItems0SavedSearch) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("saved_search"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this saved searches items0 saved search based on context it is used
func (m *SavedSearchesItems0SavedSearch) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SavedSearchesItems0SavedSearch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SavedSearchesItems0SavedSearch) UnmarshalBinary(b []byte) error {
	var res SavedSearchesItems0SavedSearch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}