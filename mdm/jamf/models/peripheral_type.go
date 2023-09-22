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

// PeripheralType peripheral type
//
// swagger:model peripheral_type
type PeripheralType struct {

	// fields
	Fields []*PeripheralTypeFieldsItems0 `json:"fields"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// name
	// Example: Peripheral Type Name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this peripheral type
func (m *PeripheralType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
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

func (m *PeripheralType) validateFields(formats strfmt.Registry) error {
	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	for i := 0; i < len(m.Fields); i++ {
		if swag.IsZero(m.Fields[i]) { // not required
			continue
		}

		if m.Fields[i] != nil {
			if err := m.Fields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PeripheralType) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this peripheral type based on the context it is used
func (m *PeripheralType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeripheralType) contextValidateFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fields); i++ {

		if m.Fields[i] != nil {

			if swag.IsZero(m.Fields[i]) { // not required
				return nil
			}

			if err := m.Fields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeripheralType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeripheralType) UnmarshalBinary(b []byte) error {
	var res PeripheralType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PeripheralTypeFieldsItems0 peripheral type fields items0
//
// swagger:model PeripheralTypeFieldsItems0
type PeripheralTypeFieldsItems0 struct {

	// field
	Field *PeripheralTypeFieldsItems0Field `json:"field,omitempty"`
}

// Validate validates this peripheral type fields items0
func (m *PeripheralTypeFieldsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateField(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeripheralTypeFieldsItems0) validateField(formats strfmt.Registry) error {
	if swag.IsZero(m.Field) { // not required
		return nil
	}

	if m.Field != nil {
		if err := m.Field.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("field")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("field")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this peripheral type fields items0 based on the context it is used
func (m *PeripheralTypeFieldsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateField(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeripheralTypeFieldsItems0) contextValidateField(ctx context.Context, formats strfmt.Registry) error {

	if m.Field != nil {

		if swag.IsZero(m.Field) { // not required
			return nil
		}

		if err := m.Field.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("field")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("field")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeripheralTypeFieldsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeripheralTypeFieldsItems0) UnmarshalBinary(b []byte) error {
	var res PeripheralTypeFieldsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PeripheralTypeFieldsItems0Field peripheral type fields items0 field
//
// swagger:model PeripheralTypeFieldsItems0Field
type PeripheralTypeFieldsItems0Field struct {

	// choices
	Choices []*PeripheralTypeFieldsItems0FieldChoicesItems0 `json:"choices"`

	// name
	// Example: Peripheral Field Name
	Name string `json:"name,omitempty"`

	// order
	// Example: 1
	Order int64 `json:"order,omitempty"`

	// type
	// Enum: [menu text]
	Type string `json:"type,omitempty"`
}

// Validate validates this peripheral type fields items0 field
func (m *PeripheralTypeFieldsItems0Field) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChoices(formats); err != nil {
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

func (m *PeripheralTypeFieldsItems0Field) validateChoices(formats strfmt.Registry) error {
	if swag.IsZero(m.Choices) { // not required
		return nil
	}

	for i := 0; i < len(m.Choices); i++ {
		if swag.IsZero(m.Choices[i]) { // not required
			continue
		}

		if m.Choices[i] != nil {
			if err := m.Choices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("field" + "." + "choices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("field" + "." + "choices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var peripheralTypeFieldsItems0FieldTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["menu","text"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		peripheralTypeFieldsItems0FieldTypeTypePropEnum = append(peripheralTypeFieldsItems0FieldTypeTypePropEnum, v)
	}
}

const (

	// PeripheralTypeFieldsItems0FieldTypeMenu captures enum value "menu"
	PeripheralTypeFieldsItems0FieldTypeMenu string = "menu"

	// PeripheralTypeFieldsItems0FieldTypeText captures enum value "text"
	PeripheralTypeFieldsItems0FieldTypeText string = "text"
)

// prop value enum
func (m *PeripheralTypeFieldsItems0Field) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, peripheralTypeFieldsItems0FieldTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PeripheralTypeFieldsItems0Field) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("field"+"."+"type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this peripheral type fields items0 field based on the context it is used
func (m *PeripheralTypeFieldsItems0Field) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChoices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeripheralTypeFieldsItems0Field) contextValidateChoices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Choices); i++ {

		if m.Choices[i] != nil {

			if swag.IsZero(m.Choices[i]) { // not required
				return nil
			}

			if err := m.Choices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("field" + "." + "choices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("field" + "." + "choices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeripheralTypeFieldsItems0Field) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeripheralTypeFieldsItems0Field) UnmarshalBinary(b []byte) error {
	var res PeripheralTypeFieldsItems0Field
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PeripheralTypeFieldsItems0FieldChoicesItems0 peripheral type fields items0 field choices items0
//
// swagger:model PeripheralTypeFieldsItems0FieldChoicesItems0
type PeripheralTypeFieldsItems0FieldChoicesItems0 struct {

	// choice
	// Example: Value for menu type
	Choice string `json:"choice,omitempty"`
}

// Validate validates this peripheral type fields items0 field choices items0
func (m *PeripheralTypeFieldsItems0FieldChoicesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this peripheral type fields items0 field choices items0 based on context it is used
func (m *PeripheralTypeFieldsItems0FieldChoicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PeripheralTypeFieldsItems0FieldChoicesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeripheralTypeFieldsItems0FieldChoicesItems0) UnmarshalBinary(b []byte) error {
	var res PeripheralTypeFieldsItems0FieldChoicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}