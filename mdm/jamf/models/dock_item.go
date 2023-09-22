// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DockItem dock item
//
// swagger:model dock_item
type DockItem struct {

	// contents
	Contents string `json:"contents,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Name of the dock item
	// Example: Safari
	// Required: true
	Name *string `json:"name"`

	// path
	// Example: file://localhost/Applications/Safari.app/
	// Required: true
	Path *string `json:"path"`

	// type
	// Required: true
	// Enum: [App File Folder]
	Type *string `json:"type"`
}

// Validate validates this dock item
func (m *DockItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DockItem) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DockItem) validatePath(formats strfmt.Registry) error {

	if err := validate.Required("path", "body", m.Path); err != nil {
		return err
	}

	return nil
}

var dockItemTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["App","File","Folder"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dockItemTypeTypePropEnum = append(dockItemTypeTypePropEnum, v)
	}
}

const (

	// DockItemTypeApp captures enum value "App"
	DockItemTypeApp string = "App"

	// DockItemTypeFile captures enum value "File"
	DockItemTypeFile string = "File"

	// DockItemTypeFolder captures enum value "Folder"
	DockItemTypeFolder string = "Folder"
)

// prop value enum
func (m *DockItem) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dockItemTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DockItem) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dock item based on context it is used
func (m *DockItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DockItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DockItem) UnmarshalBinary(b []byte) error {
	var res DockItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}