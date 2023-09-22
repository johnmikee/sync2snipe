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

// DiskEncryptionConfiguration disk encryption configuration
//
// swagger:model disk_encryption_configuration
type DiskEncryptionConfiguration struct {

	// file vault enabled users
	// Enum: [Current or Next User Management Account]
	FileVaultEnabledUsers *string `json:"file_vault_enabled_users,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// key type
	// Enum: [Individual Institutional Individual And Institutional]
	KeyType *string `json:"key_type,omitempty"`

	// Name of the disk encryption configuration
	// Example: Corporate Encryption
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this disk encryption configuration
func (m *DiskEncryptionConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileVaultEnabledUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyType(formats); err != nil {
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

var diskEncryptionConfigurationTypeFileVaultEnabledUsersPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Current or Next User","Management Account"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diskEncryptionConfigurationTypeFileVaultEnabledUsersPropEnum = append(diskEncryptionConfigurationTypeFileVaultEnabledUsersPropEnum, v)
	}
}

const (

	// DiskEncryptionConfigurationFileVaultEnabledUsersCurrentOrNextUser captures enum value "Current or Next User"
	DiskEncryptionConfigurationFileVaultEnabledUsersCurrentOrNextUser string = "Current or Next User"

	// DiskEncryptionConfigurationFileVaultEnabledUsersManagementAccount captures enum value "Management Account"
	DiskEncryptionConfigurationFileVaultEnabledUsersManagementAccount string = "Management Account"
)

// prop value enum
func (m *DiskEncryptionConfiguration) validateFileVaultEnabledUsersEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, diskEncryptionConfigurationTypeFileVaultEnabledUsersPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DiskEncryptionConfiguration) validateFileVaultEnabledUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.FileVaultEnabledUsers) { // not required
		return nil
	}

	// value enum
	if err := m.validateFileVaultEnabledUsersEnum("file_vault_enabled_users", "body", *m.FileVaultEnabledUsers); err != nil {
		return err
	}

	return nil
}

var diskEncryptionConfigurationTypeKeyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Individual","Institutional","Individual And Institutional"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		diskEncryptionConfigurationTypeKeyTypePropEnum = append(diskEncryptionConfigurationTypeKeyTypePropEnum, v)
	}
}

const (

	// DiskEncryptionConfigurationKeyTypeIndividual captures enum value "Individual"
	DiskEncryptionConfigurationKeyTypeIndividual string = "Individual"

	// DiskEncryptionConfigurationKeyTypeInstitutional captures enum value "Institutional"
	DiskEncryptionConfigurationKeyTypeInstitutional string = "Institutional"

	// DiskEncryptionConfigurationKeyTypeIndividualAndInstitutional captures enum value "Individual And Institutional"
	DiskEncryptionConfigurationKeyTypeIndividualAndInstitutional string = "Individual And Institutional"
)

// prop value enum
func (m *DiskEncryptionConfiguration) validateKeyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, diskEncryptionConfigurationTypeKeyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DiskEncryptionConfiguration) validateKeyType(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateKeyTypeEnum("key_type", "body", *m.KeyType); err != nil {
		return err
	}

	return nil
}

func (m *DiskEncryptionConfiguration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this disk encryption configuration based on context it is used
func (m *DiskEncryptionConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DiskEncryptionConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiskEncryptionConfiguration) UnmarshalBinary(b []byte) error {
	var res DiskEncryptionConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}