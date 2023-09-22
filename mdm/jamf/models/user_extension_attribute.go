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

// UserExtensionAttribute user extension attribute
//
// swagger:model user_extension_attribute
type UserExtensionAttribute struct {

	// data type
	// Enum: [String Integer Date]
	DataType string `json:"data_type,omitempty"`

	// description
	// Example: Text field for logging custom data
	Description string `json:"description,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// input type
	InputType *UserExtensionAttributeInputType `json:"input_type,omitempty"`

	// Name of the user extension attribute
	// Example: User Attributes
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this user extension attribute
func (m *UserExtensionAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputType(formats); err != nil {
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

var userExtensionAttributeTypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["String","Integer","Date"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userExtensionAttributeTypeDataTypePropEnum = append(userExtensionAttributeTypeDataTypePropEnum, v)
	}
}

const (

	// UserExtensionAttributeDataTypeString captures enum value "String"
	UserExtensionAttributeDataTypeString string = "String"

	// UserExtensionAttributeDataTypeInteger captures enum value "Integer"
	UserExtensionAttributeDataTypeInteger string = "Integer"

	// UserExtensionAttributeDataTypeDate captures enum value "Date"
	UserExtensionAttributeDataTypeDate string = "Date"
)

// prop value enum
func (m *UserExtensionAttribute) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userExtensionAttributeTypeDataTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserExtensionAttribute) validateDataType(formats strfmt.Registry) error {
	if swag.IsZero(m.DataType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataTypeEnum("data_type", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *UserExtensionAttribute) validateInputType(formats strfmt.Registry) error {
	if swag.IsZero(m.InputType) { // not required
		return nil
	}

	if m.InputType != nil {
		if err := m.InputType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input_type")
			}
			return err
		}
	}

	return nil
}

func (m *UserExtensionAttribute) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user extension attribute based on the context it is used
func (m *UserExtensionAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserExtensionAttribute) contextValidateInputType(ctx context.Context, formats strfmt.Registry) error {

	if m.InputType != nil {

		if swag.IsZero(m.InputType) { // not required
			return nil
		}

		if err := m.InputType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("input_type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserExtensionAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserExtensionAttribute) UnmarshalBinary(b []byte) error {
	var res UserExtensionAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UserExtensionAttributeInputType user extension attribute input type
//
// swagger:model UserExtensionAttributeInputType
type UserExtensionAttributeInputType struct {

	// type
	// Enum: [Text Field Pop-up Menu]
	Type *string `json:"type,omitempty"`
}

// Validate validates this user extension attribute input type
func (m *UserExtensionAttributeInputType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userExtensionAttributeInputTypeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Text Field","Pop-up Menu"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userExtensionAttributeInputTypeTypeTypePropEnum = append(userExtensionAttributeInputTypeTypeTypePropEnum, v)
	}
}

const (

	// UserExtensionAttributeInputTypeTypeTextField captures enum value "Text Field"
	UserExtensionAttributeInputTypeTypeTextField string = "Text Field"

	// UserExtensionAttributeInputTypeTypePopDashUpMenu captures enum value "Pop-up Menu"
	UserExtensionAttributeInputTypeTypePopDashUpMenu string = "Pop-up Menu"
)

// prop value enum
func (m *UserExtensionAttributeInputType) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userExtensionAttributeInputTypeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserExtensionAttributeInputType) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("input_type"+"."+"type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user extension attribute input type based on context it is used
func (m *UserExtensionAttributeInputType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserExtensionAttributeInputType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserExtensionAttributeInputType) UnmarshalBinary(b []byte) error {
	var res UserExtensionAttributeInputType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}