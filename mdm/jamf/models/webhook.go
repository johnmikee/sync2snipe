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

// Webhook webhook
//
// swagger:model webhook
type Webhook struct {

	// authentication type
	// Enum: [NONE BASIC]
	AuthenticationType *string `json:"authentication_type,omitempty"`

	// Number of seconds to attempt to connect to the webhooks host server
	ConnectionTimeout *int64 `json:"connection_timeout,omitempty"`

	// content type
	// Enum: [text/xml application/json]
	ContentType *string `json:"content_type,omitempty"`

	// display fields
	DisplayFields []*WebhookDisplayFieldsItems0 `json:"display_fields"`

	// enable display fields for group object
	EnableDisplayFieldsForGroupObject *bool `json:"enable_display_fields_for_group_object,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// event
	// Required: true
	// Enum: [ComputerAdded ComputerCheckIn ComputerInventoryCompleted ComputerPolicyFinished ComputerPushCapabilityChanged JSSShutdown JSSStartup MobileDeviceCheckIn MobileDeviceCommandCompleted MobileDeviceEnrolled MobileDevicePushSent MobileDeviceUnEnrolled PatchSoftwareTitleUpdated PushSent RestAPIOperation SCEPChallenge SmartGroupComputerMembershipChange SmartGroupMobileDeviceMembershipChange]
	Event *string `json:"event"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// name
	// Example: Computer Enrolled Hook
	// Required: true
	Name *string `json:"name"`

	// password
	Password string `json:"password,omitempty"`

	// Number of seconds to wait for a response from the webhooks host server after sending a request
	ReadTimeout *int64 `json:"read_timeout,omitempty"`

	// url
	// Example: https://requestb.in/wsfasfws
	// Required: true
	URL *string `json:"url"`

	// username
	// Example: webhook_admin
	Username string `json:"username,omitempty"`
}

// Validate validates this webhook
func (m *Webhook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var webhookTypeAuthenticationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","BASIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webhookTypeAuthenticationTypePropEnum = append(webhookTypeAuthenticationTypePropEnum, v)
	}
}

const (

	// WebhookAuthenticationTypeNONE captures enum value "NONE"
	WebhookAuthenticationTypeNONE string = "NONE"

	// WebhookAuthenticationTypeBASIC captures enum value "BASIC"
	WebhookAuthenticationTypeBASIC string = "BASIC"
)

// prop value enum
func (m *Webhook) validateAuthenticationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webhookTypeAuthenticationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Webhook) validateAuthenticationType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationTypeEnum("authentication_type", "body", *m.AuthenticationType); err != nil {
		return err
	}

	return nil
}

var webhookTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text/xml","application/json"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webhookTypeContentTypePropEnum = append(webhookTypeContentTypePropEnum, v)
	}
}

const (

	// WebhookContentTypeTextXML captures enum value "text/xml"
	WebhookContentTypeTextXML string = "text/xml"

	// WebhookContentTypeApplicationJSON captures enum value "application/json"
	WebhookContentTypeApplicationJSON string = "application/json"
)

// prop value enum
func (m *Webhook) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webhookTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Webhook) validateContentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContentTypeEnum("content_type", "body", *m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateDisplayFields(formats strfmt.Registry) error {
	if swag.IsZero(m.DisplayFields) { // not required
		return nil
	}

	for i := 0; i < len(m.DisplayFields); i++ {
		if swag.IsZero(m.DisplayFields[i]) { // not required
			continue
		}

		if m.DisplayFields[i] != nil {
			if err := m.DisplayFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("display_fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("display_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var webhookTypeEventPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ComputerAdded","ComputerCheckIn","ComputerInventoryCompleted","ComputerPolicyFinished","ComputerPushCapabilityChanged","JSSShutdown","JSSStartup","MobileDeviceCheckIn","MobileDeviceCommandCompleted","MobileDeviceEnrolled","MobileDevicePushSent","MobileDeviceUnEnrolled","PatchSoftwareTitleUpdated","PushSent","RestAPIOperation","SCEPChallenge","SmartGroupComputerMembershipChange","SmartGroupMobileDeviceMembershipChange"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webhookTypeEventPropEnum = append(webhookTypeEventPropEnum, v)
	}
}

const (

	// WebhookEventComputerAdded captures enum value "ComputerAdded"
	WebhookEventComputerAdded string = "ComputerAdded"

	// WebhookEventComputerCheckIn captures enum value "ComputerCheckIn"
	WebhookEventComputerCheckIn string = "ComputerCheckIn"

	// WebhookEventComputerInventoryCompleted captures enum value "ComputerInventoryCompleted"
	WebhookEventComputerInventoryCompleted string = "ComputerInventoryCompleted"

	// WebhookEventComputerPolicyFinished captures enum value "ComputerPolicyFinished"
	WebhookEventComputerPolicyFinished string = "ComputerPolicyFinished"

	// WebhookEventComputerPushCapabilityChanged captures enum value "ComputerPushCapabilityChanged"
	WebhookEventComputerPushCapabilityChanged string = "ComputerPushCapabilityChanged"

	// WebhookEventJSSShutdown captures enum value "JSSShutdown"
	WebhookEventJSSShutdown string = "JSSShutdown"

	// WebhookEventJSSStartup captures enum value "JSSStartup"
	WebhookEventJSSStartup string = "JSSStartup"

	// WebhookEventMobileDeviceCheckIn captures enum value "MobileDeviceCheckIn"
	WebhookEventMobileDeviceCheckIn string = "MobileDeviceCheckIn"

	// WebhookEventMobileDeviceCommandCompleted captures enum value "MobileDeviceCommandCompleted"
	WebhookEventMobileDeviceCommandCompleted string = "MobileDeviceCommandCompleted"

	// WebhookEventMobileDeviceEnrolled captures enum value "MobileDeviceEnrolled"
	WebhookEventMobileDeviceEnrolled string = "MobileDeviceEnrolled"

	// WebhookEventMobileDevicePushSent captures enum value "MobileDevicePushSent"
	WebhookEventMobileDevicePushSent string = "MobileDevicePushSent"

	// WebhookEventMobileDeviceUnEnrolled captures enum value "MobileDeviceUnEnrolled"
	WebhookEventMobileDeviceUnEnrolled string = "MobileDeviceUnEnrolled"

	// WebhookEventPatchSoftwareTitleUpdated captures enum value "PatchSoftwareTitleUpdated"
	WebhookEventPatchSoftwareTitleUpdated string = "PatchSoftwareTitleUpdated"

	// WebhookEventPushSent captures enum value "PushSent"
	WebhookEventPushSent string = "PushSent"

	// WebhookEventRestAPIOperation captures enum value "RestAPIOperation"
	WebhookEventRestAPIOperation string = "RestAPIOperation"

	// WebhookEventSCEPChallenge captures enum value "SCEPChallenge"
	WebhookEventSCEPChallenge string = "SCEPChallenge"

	// WebhookEventSmartGroupComputerMembershipChange captures enum value "SmartGroupComputerMembershipChange"
	WebhookEventSmartGroupComputerMembershipChange string = "SmartGroupComputerMembershipChange"

	// WebhookEventSmartGroupMobileDeviceMembershipChange captures enum value "SmartGroupMobileDeviceMembershipChange"
	WebhookEventSmartGroupMobileDeviceMembershipChange string = "SmartGroupMobileDeviceMembershipChange"
)

// prop value enum
func (m *Webhook) validateEventEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webhookTypeEventPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Webhook) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event); err != nil {
		return err
	}

	// value enum
	if err := m.validateEventEnum("event", "body", *m.Event); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Webhook) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this webhook based on the context it is used
func (m *Webhook) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplayFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Webhook) contextValidateDisplayFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DisplayFields); i++ {

		if m.DisplayFields[i] != nil {

			if swag.IsZero(m.DisplayFields[i]) { // not required
				return nil
			}

			if err := m.DisplayFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("display_fields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("display_fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Webhook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Webhook) UnmarshalBinary(b []byte) error {
	var res Webhook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WebhookDisplayFieldsItems0 webhook display fields items0
//
// swagger:model WebhookDisplayFieldsItems0
type WebhookDisplayFieldsItems0 struct {

	// display field
	DisplayField *WebhookDisplayFieldsItems0DisplayField `json:"display_field,omitempty"`

	// size
	Size Size `json:"size,omitempty"`
}

// Validate validates this webhook display fields items0
func (m *WebhookDisplayFieldsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayField(formats); err != nil {
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

func (m *WebhookDisplayFieldsItems0) validateDisplayField(formats strfmt.Registry) error {
	if swag.IsZero(m.DisplayField) { // not required
		return nil
	}

	if m.DisplayField != nil {
		if err := m.DisplayField.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display_field")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display_field")
			}
			return err
		}
	}

	return nil
}

func (m *WebhookDisplayFieldsItems0) validateSize(formats strfmt.Registry) error {
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

// ContextValidate validate this webhook display fields items0 based on the context it is used
func (m *WebhookDisplayFieldsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisplayField(ctx, formats); err != nil {
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

func (m *WebhookDisplayFieldsItems0) contextValidateDisplayField(ctx context.Context, formats strfmt.Registry) error {

	if m.DisplayField != nil {

		if swag.IsZero(m.DisplayField) { // not required
			return nil
		}

		if err := m.DisplayField.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display_field")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("display_field")
			}
			return err
		}
	}

	return nil
}

func (m *WebhookDisplayFieldsItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

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
func (m *WebhookDisplayFieldsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookDisplayFieldsItems0) UnmarshalBinary(b []byte) error {
	var res WebhookDisplayFieldsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// WebhookDisplayFieldsItems0DisplayField webhook display fields items0 display field
//
// swagger:model WebhookDisplayFieldsItems0DisplayField
type WebhookDisplayFieldsItems0DisplayField struct {

	// Name of the display field to include for smart group based webhook events
	// Example: IP Address
	Name string `json:"name,omitempty"`
}

// Validate validates this webhook display fields items0 display field
func (m *WebhookDisplayFieldsItems0DisplayField) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this webhook display fields items0 display field based on context it is used
func (m *WebhookDisplayFieldsItems0DisplayField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookDisplayFieldsItems0DisplayField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookDisplayFieldsItems0DisplayField) UnmarshalBinary(b []byte) error {
	var res WebhookDisplayFieldsItems0DisplayField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
