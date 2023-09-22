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
	"github.com/go-openapi/validate"
)

// MobileDeviceGroup mobile device group
//
// swagger:model mobile_device_group
type MobileDeviceGroup struct {

	// criteria
	Criteria []*MobileDeviceGroupCriteriaItems0 `json:"criteria"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// is smart
	// Required: true
	IsSmart *bool `json:"is_smart"`

	// mobile devices
	MobileDevices []*MobileDeviceGroupMobileDevicesItems0 `json:"mobile_devices"`

	// name
	// Example: iPhones
	// Required: true
	Name *string `json:"name"`

	// site
	Site *SiteObject `json:"site,omitempty"`
}

// Validate validates this mobile device group
func (m *MobileDeviceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCriteria(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsSmart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobileDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MobileDeviceGroup) validateCriteria(formats strfmt.Registry) error {
	if swag.IsZero(m.Criteria) { // not required
		return nil
	}

	for i := 0; i < len(m.Criteria); i++ {
		if swag.IsZero(m.Criteria[i]) { // not required
			continue
		}

		if m.Criteria[i] != nil {
			if err := m.Criteria[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("criteria" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MobileDeviceGroup) validateIsSmart(formats strfmt.Registry) error {

	if err := validate.Required("is_smart", "body", m.IsSmart); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceGroup) validateMobileDevices(formats strfmt.Registry) error {
	if swag.IsZero(m.MobileDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.MobileDevices); i++ {
		if swag.IsZero(m.MobileDevices[i]) { // not required
			continue
		}

		if m.MobileDevices[i] != nil {
			if err := m.MobileDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mobile_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mobile_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MobileDeviceGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MobileDeviceGroup) validateSite(formats strfmt.Registry) error {
	if swag.IsZero(m.Site) { // not required
		return nil
	}

	if m.Site != nil {
		if err := m.Site.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mobile device group based on the context it is used
func (m *MobileDeviceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCriteria(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMobileDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MobileDeviceGroup) contextValidateCriteria(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Criteria); i++ {

		if m.Criteria[i] != nil {

			if swag.IsZero(m.Criteria[i]) { // not required
				return nil
			}

			if err := m.Criteria[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("criteria" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("criteria" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MobileDeviceGroup) contextValidateMobileDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MobileDevices); i++ {

		if m.MobileDevices[i] != nil {

			if swag.IsZero(m.MobileDevices[i]) { // not required
				return nil
			}

			if err := m.MobileDevices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mobile_devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mobile_devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MobileDeviceGroup) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

	if m.Site != nil {

		if swag.IsZero(m.Site) { // not required
			return nil
		}

		if err := m.Site.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("site")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("site")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MobileDeviceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileDeviceGroup) UnmarshalBinary(b []byte) error {
	var res MobileDeviceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MobileDeviceGroupCriteriaItems0 mobile device group criteria items0
//
// swagger:model MobileDeviceGroupCriteriaItems0
type MobileDeviceGroupCriteriaItems0 struct {

	// criterion
	Criterion *Criterion `json:"criterion,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this mobile device group criteria items0
func (m *MobileDeviceGroupCriteriaItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCriterion(formats); err != nil {
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

func (m *MobileDeviceGroupCriteriaItems0) validateCriterion(formats strfmt.Registry) error {
	if swag.IsZero(m.Criterion) { // not required
		return nil
	}

	if m.Criterion != nil {
		if err := m.Criterion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("criterion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("criterion")
			}
			return err
		}
	}

	return nil
}

func (m *MobileDeviceGroupCriteriaItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this mobile device group criteria items0 based on the context it is used
func (m *MobileDeviceGroupCriteriaItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCriterion(ctx, formats); err != nil {
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

func (m *MobileDeviceGroupCriteriaItems0) contextValidateCriterion(ctx context.Context, formats strfmt.Registry) error {

	if m.Criterion != nil {

		if swag.IsZero(m.Criterion) { // not required
			return nil
		}

		if err := m.Criterion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("criterion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("criterion")
			}
			return err
		}
	}

	return nil
}

func (m *MobileDeviceGroupCriteriaItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *MobileDeviceGroupCriteriaItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileDeviceGroupCriteriaItems0) UnmarshalBinary(b []byte) error {
	var res MobileDeviceGroupCriteriaItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MobileDeviceGroupMobileDevicesItems0 mobile device group mobile devices items0
//
// swagger:model MobileDeviceGroupMobileDevicesItems0
type MobileDeviceGroupMobileDevicesItems0 struct {

	// mobile device
	MobileDevice *MobileDeviceGroupMobileDevicesItems0MobileDevice `json:"mobile_device,omitempty"`
}

// Validate validates this mobile device group mobile devices items0
func (m *MobileDeviceGroupMobileDevicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMobileDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MobileDeviceGroupMobileDevicesItems0) validateMobileDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.MobileDevice) { // not required
		return nil
	}

	if m.MobileDevice != nil {
		if err := m.MobileDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobile_device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mobile_device")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mobile device group mobile devices items0 based on the context it is used
func (m *MobileDeviceGroupMobileDevicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMobileDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MobileDeviceGroupMobileDevicesItems0) contextValidateMobileDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.MobileDevice != nil {

		if swag.IsZero(m.MobileDevice) { // not required
			return nil
		}

		if err := m.MobileDevice.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mobile_device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mobile_device")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MobileDeviceGroupMobileDevicesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileDeviceGroupMobileDevicesItems0) UnmarshalBinary(b []byte) error {
	var res MobileDeviceGroupMobileDevicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MobileDeviceGroupMobileDevicesItems0MobileDevice mobile device group mobile devices items0 mobile device
//
// swagger:model MobileDeviceGroupMobileDevicesItems0MobileDevice
type MobileDeviceGroupMobileDevicesItems0MobileDevice struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// mac address
	// Example: E0:AC:CB:97:36:G4
	MacAddress string `json:"mac_address,omitempty"`

	// Name of the device
	// Example: Shawns iPhone
	Name string `json:"name,omitempty"`

	// serial number
	// Example: C02Q7KHTGFWF
	SerialNumber string `json:"serial_number,omitempty"`

	// udid
	// Example: 55900BDC-347C-58B1-D249-F32244B11D30
	Udid string `json:"udid,omitempty"`

	// wifi mac address
	// Example: E0:AC:CB:97:36:G4
	WifiMacAddress string `json:"wifi_mac_address,omitempty"`
}

// Validate validates this mobile device group mobile devices items0 mobile device
func (m *MobileDeviceGroupMobileDevicesItems0MobileDevice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mobile device group mobile devices items0 mobile device based on context it is used
func (m *MobileDeviceGroupMobileDevicesItems0MobileDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MobileDeviceGroupMobileDevicesItems0MobileDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MobileDeviceGroupMobileDevicesItems0MobileDevice) UnmarshalBinary(b []byte) error {
	var res MobileDeviceGroupMobileDevicesItems0MobileDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}