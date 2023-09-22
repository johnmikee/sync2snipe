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

// LdapServer ldap server
//
// swagger:model ldap_server
type LdapServer struct {

	// connection
	Connection *LdapServerConnection `json:"connection,omitempty"`

	// mappings for users
	MappingsForUsers *LdapServerMappingsForUsers `json:"mappings_for_users,omitempty"`
}

// Validate validates this ldap server
func (m *LdapServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMappingsForUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServer) validateConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *LdapServer) validateMappingsForUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.MappingsForUsers) { // not required
		return nil
	}

	if m.MappingsForUsers != nil {
		if err := m.MappingsForUsers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mappings_for_users")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mappings_for_users")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap server based on the context it is used
func (m *LdapServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMappingsForUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServer) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {

		if swag.IsZero(m.Connection) { // not required
			return nil
		}

		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *LdapServer) contextValidateMappingsForUsers(ctx context.Context, formats strfmt.Registry) error {

	if m.MappingsForUsers != nil {

		if swag.IsZero(m.MappingsForUsers) { // not required
			return nil
		}

		if err := m.MappingsForUsers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mappings_for_users")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mappings_for_users")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServer) UnmarshalBinary(b []byte) error {
	var res LdapServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapServerConnection ldap server connection
//
// swagger:model LdapServerConnection
type LdapServerConnection struct {

	// account
	Account *LdapServerConnectionAccount `json:"account,omitempty"`

	// authentication type
	// Enum: [simple CRAM-MD5 DIGEST-MD5 none]
	AuthenticationType string `json:"authentication_type,omitempty"`

	// Hostname or IP address of the server
	// Example: company.ad.com
	Hostname string `json:"hostname,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// Name of the LDAP server
	// Example: Company Active Directory
	// Required: true
	Name *string `json:"name"`

	// Timeout in seconds
	// Example: 15
	OpenCloseTimeout int64 `json:"open_close_timeout,omitempty"`

	// port
	// Example: 389
	Port int64 `json:"port,omitempty"`

	// referral response
	// Enum: [ignore follow]
	ReferralResponse string `json:"referral_response,omitempty"`

	// Timeout in seconds
	// Example: 60
	SearchTimeout int64 `json:"search_timeout,omitempty"`

	// server type
	// Enum: [Active Directory Open Directory eDirectory Custom]
	ServerType string `json:"server_type,omitempty"`

	// use ssl
	UseSsl *bool `json:"use_ssl,omitempty"`

	// use wildcards
	UseWildcards bool `json:"use_wildcards,omitempty"`
}

// Validate validates this ldap server connection
func (m *LdapServerConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferralResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServerConnection) validateAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection" + "." + "account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection" + "." + "account")
			}
			return err
		}
	}

	return nil
}

var ldapServerConnectionTypeAuthenticationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["simple","CRAM-MD5","DIGEST-MD5","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerConnectionTypeAuthenticationTypePropEnum = append(ldapServerConnectionTypeAuthenticationTypePropEnum, v)
	}
}

const (

	// LdapServerConnectionAuthenticationTypeSimple captures enum value "simple"
	LdapServerConnectionAuthenticationTypeSimple string = "simple"

	// LdapServerConnectionAuthenticationTypeCRAMDashMD5 captures enum value "CRAM-MD5"
	LdapServerConnectionAuthenticationTypeCRAMDashMD5 string = "CRAM-MD5"

	// LdapServerConnectionAuthenticationTypeDIGESTDashMD5 captures enum value "DIGEST-MD5"
	LdapServerConnectionAuthenticationTypeDIGESTDashMD5 string = "DIGEST-MD5"

	// LdapServerConnectionAuthenticationTypeNone captures enum value "none"
	LdapServerConnectionAuthenticationTypeNone string = "none"
)

// prop value enum
func (m *LdapServerConnection) validateAuthenticationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerConnectionTypeAuthenticationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerConnection) validateAuthenticationType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthenticationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationTypeEnum("connection"+"."+"authentication_type", "body", m.AuthenticationType); err != nil {
		return err
	}

	return nil
}

func (m *LdapServerConnection) validateName(formats strfmt.Registry) error {

	if err := validate.Required("connection"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var ldapServerConnectionTypeReferralResponsePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ignore","follow"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerConnectionTypeReferralResponsePropEnum = append(ldapServerConnectionTypeReferralResponsePropEnum, v)
	}
}

const (

	// LdapServerConnectionReferralResponseIgnore captures enum value "ignore"
	LdapServerConnectionReferralResponseIgnore string = "ignore"

	// LdapServerConnectionReferralResponseFollow captures enum value "follow"
	LdapServerConnectionReferralResponseFollow string = "follow"
)

// prop value enum
func (m *LdapServerConnection) validateReferralResponseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerConnectionTypeReferralResponsePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerConnection) validateReferralResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.ReferralResponse) { // not required
		return nil
	}

	// value enum
	if err := m.validateReferralResponseEnum("connection"+"."+"referral_response", "body", m.ReferralResponse); err != nil {
		return err
	}

	return nil
}

var ldapServerConnectionTypeServerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active Directory","Open Directory","eDirectory","Custom"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerConnectionTypeServerTypePropEnum = append(ldapServerConnectionTypeServerTypePropEnum, v)
	}
}

const (

	// LdapServerConnectionServerTypeActiveDirectory captures enum value "Active Directory"
	LdapServerConnectionServerTypeActiveDirectory string = "Active Directory"

	// LdapServerConnectionServerTypeOpenDirectory captures enum value "Open Directory"
	LdapServerConnectionServerTypeOpenDirectory string = "Open Directory"

	// LdapServerConnectionServerTypeEDirectory captures enum value "eDirectory"
	LdapServerConnectionServerTypeEDirectory string = "eDirectory"

	// LdapServerConnectionServerTypeCustom captures enum value "Custom"
	LdapServerConnectionServerTypeCustom string = "Custom"
)

// prop value enum
func (m *LdapServerConnection) validateServerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerConnectionTypeServerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerConnection) validateServerType(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateServerTypeEnum("connection"+"."+"server_type", "body", m.ServerType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ldap server connection based on the context it is used
func (m *LdapServerConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServerConnection) contextValidateAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.Account != nil {

		if swag.IsZero(m.Account) { // not required
			return nil
		}

		if err := m.Account.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection" + "." + "account")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection" + "." + "account")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapServerConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServerConnection) UnmarshalBinary(b []byte) error {
	var res LdapServerConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapServerConnectionAccount ldap server connection account
//
// swagger:model LdapServerConnectionAccount
type LdapServerConnectionAccount struct {

	// distinguished username
	// Example: CN=Administrator,CN=Users,DC=Company,DC=com
	DistinguishedUsername string `json:"distinguished_username,omitempty"`

	// password
	// Example: password
	Password string `json:"password,omitempty"`
}

// Validate validates this ldap server connection account
func (m *LdapServerConnectionAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ldap server connection account based on context it is used
func (m *LdapServerConnectionAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapServerConnectionAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServerConnectionAccount) UnmarshalBinary(b []byte) error {
	var res LdapServerConnectionAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapServerMappingsForUsers ldap server mappings for users
//
// swagger:model LdapServerMappingsForUsers
type LdapServerMappingsForUsers struct {

	// user group mappings
	UserGroupMappings *LdapServerMappingsForUsersUserGroupMappings `json:"user_group_mappings,omitempty"`

	// user group membership mappings
	UserGroupMembershipMappings *LdapServerMappingsForUsersUserGroupMembershipMappings `json:"user_group_membership_mappings,omitempty"`

	// user mappings
	UserMappings *LdapServerMappingsForUsersUserMappings `json:"user_mappings,omitempty"`
}

// Validate validates this ldap server mappings for users
func (m *LdapServerMappingsForUsers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserGroupMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserGroupMembershipMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserMappings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServerMappingsForUsers) validateUserGroupMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.UserGroupMappings) { // not required
		return nil
	}

	if m.UserGroupMappings != nil {
		if err := m.UserGroupMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mappings_for_users" + "." + "user_group_mappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mappings_for_users" + "." + "user_group_mappings")
			}
			return err
		}
	}

	return nil
}

func (m *LdapServerMappingsForUsers) validateUserGroupMembershipMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.UserGroupMembershipMappings) { // not required
		return nil
	}

	if m.UserGroupMembershipMappings != nil {
		if err := m.UserGroupMembershipMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mappings_for_users" + "." + "user_group_membership_mappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mappings_for_users" + "." + "user_group_membership_mappings")
			}
			return err
		}
	}

	return nil
}

func (m *LdapServerMappingsForUsers) validateUserMappings(formats strfmt.Registry) error {
	if swag.IsZero(m.UserMappings) { // not required
		return nil
	}

	if m.UserMappings != nil {
		if err := m.UserMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mappings_for_users" + "." + "user_mappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mappings_for_users" + "." + "user_mappings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ldap server mappings for users based on the context it is used
func (m *LdapServerMappingsForUsers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserGroupMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserGroupMembershipMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LdapServerMappingsForUsers) contextValidateUserGroupMappings(ctx context.Context, formats strfmt.Registry) error {

	if m.UserGroupMappings != nil {

		if swag.IsZero(m.UserGroupMappings) { // not required
			return nil
		}

		if err := m.UserGroupMappings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mappings_for_users" + "." + "user_group_mappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mappings_for_users" + "." + "user_group_mappings")
			}
			return err
		}
	}

	return nil
}

func (m *LdapServerMappingsForUsers) contextValidateUserGroupMembershipMappings(ctx context.Context, formats strfmt.Registry) error {

	if m.UserGroupMembershipMappings != nil {

		if swag.IsZero(m.UserGroupMembershipMappings) { // not required
			return nil
		}

		if err := m.UserGroupMembershipMappings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mappings_for_users" + "." + "user_group_membership_mappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mappings_for_users" + "." + "user_group_membership_mappings")
			}
			return err
		}
	}

	return nil
}

func (m *LdapServerMappingsForUsers) contextValidateUserMappings(ctx context.Context, formats strfmt.Registry) error {

	if m.UserMappings != nil {

		if swag.IsZero(m.UserMappings) { // not required
			return nil
		}

		if err := m.UserMappings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mappings_for_users" + "." + "user_mappings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mappings_for_users" + "." + "user_mappings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LdapServerMappingsForUsers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServerMappingsForUsers) UnmarshalBinary(b []byte) error {
	var res LdapServerMappingsForUsers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapServerMappingsForUsersUserGroupMappings ldap server mappings for users user group mappings
//
// swagger:model LdapServerMappingsForUsersUserGroupMappings
type LdapServerMappingsForUsersUserGroupMappings struct {

	// map group id
	// Example: uSNCreated
	MapGroupID string `json:"map_group_id,omitempty"`

	// map group name
	// Example: name
	MapGroupName string `json:"map_group_name,omitempty"`

	// map group uuid
	// Example: objectGUID
	MapGroupUUID string `json:"map_group_uuid,omitempty"`

	// map object class to any or all
	// Enum: [all any]
	MapObjectClassToAnyOrAll string `json:"map_object_class_to_any_or_all,omitempty"`

	// object classes
	// Example: top, group
	ObjectClasses string `json:"object_classes,omitempty"`

	// search base
	// Example: DC=Company,DC=com
	SearchBase string `json:"search_base,omitempty"`

	// search scope
	// Enum: [All Subtrees First Level Only]
	SearchScope string `json:"search_scope,omitempty"`
}

// Validate validates this ldap server mappings for users user group mappings
func (m *LdapServerMappingsForUsersUserGroupMappings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMapObjectClassToAnyOrAll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ldapServerMappingsForUsersUserGroupMappingsTypeMapObjectClassToAnyOrAllPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerMappingsForUsersUserGroupMappingsTypeMapObjectClassToAnyOrAllPropEnum = append(ldapServerMappingsForUsersUserGroupMappingsTypeMapObjectClassToAnyOrAllPropEnum, v)
	}
}

const (

	// LdapServerMappingsForUsersUserGroupMappingsMapObjectClassToAnyOrAllAll captures enum value "all"
	LdapServerMappingsForUsersUserGroupMappingsMapObjectClassToAnyOrAllAll string = "all"

	// LdapServerMappingsForUsersUserGroupMappingsMapObjectClassToAnyOrAllAny captures enum value "any"
	LdapServerMappingsForUsersUserGroupMappingsMapObjectClassToAnyOrAllAny string = "any"
)

// prop value enum
func (m *LdapServerMappingsForUsersUserGroupMappings) validateMapObjectClassToAnyOrAllEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerMappingsForUsersUserGroupMappingsTypeMapObjectClassToAnyOrAllPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerMappingsForUsersUserGroupMappings) validateMapObjectClassToAnyOrAll(formats strfmt.Registry) error {
	if swag.IsZero(m.MapObjectClassToAnyOrAll) { // not required
		return nil
	}

	// value enum
	if err := m.validateMapObjectClassToAnyOrAllEnum("mappings_for_users"+"."+"user_group_mappings"+"."+"map_object_class_to_any_or_all", "body", m.MapObjectClassToAnyOrAll); err != nil {
		return err
	}

	return nil
}

var ldapServerMappingsForUsersUserGroupMappingsTypeSearchScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["All Subtrees","First Level Only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerMappingsForUsersUserGroupMappingsTypeSearchScopePropEnum = append(ldapServerMappingsForUsersUserGroupMappingsTypeSearchScopePropEnum, v)
	}
}

const (

	// LdapServerMappingsForUsersUserGroupMappingsSearchScopeAllSubtrees captures enum value "All Subtrees"
	LdapServerMappingsForUsersUserGroupMappingsSearchScopeAllSubtrees string = "All Subtrees"

	// LdapServerMappingsForUsersUserGroupMappingsSearchScopeFirstLevelOnly captures enum value "First Level Only"
	LdapServerMappingsForUsersUserGroupMappingsSearchScopeFirstLevelOnly string = "First Level Only"
)

// prop value enum
func (m *LdapServerMappingsForUsersUserGroupMappings) validateSearchScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerMappingsForUsersUserGroupMappingsTypeSearchScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerMappingsForUsersUserGroupMappings) validateSearchScope(formats strfmt.Registry) error {
	if swag.IsZero(m.SearchScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateSearchScopeEnum("mappings_for_users"+"."+"user_group_mappings"+"."+"search_scope", "body", m.SearchScope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ldap server mappings for users user group mappings based on context it is used
func (m *LdapServerMappingsForUsersUserGroupMappings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapServerMappingsForUsersUserGroupMappings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServerMappingsForUsersUserGroupMappings) UnmarshalBinary(b []byte) error {
	var res LdapServerMappingsForUsersUserGroupMappings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapServerMappingsForUsersUserGroupMembershipMappings ldap server mappings for users user group membership mappings
//
// swagger:model LdapServerMappingsForUsersUserGroupMembershipMappings
type LdapServerMappingsForUsersUserGroupMembershipMappings struct {

	// append to username
	// Example: company.com
	AppendToUsername string `json:"append_to_username,omitempty"`

	// group id
	// Example: uSNCreated
	GroupID string `json:"group_id,omitempty"`

	// map group membership to user field
	// Example: memberOf
	MapGroupMembershipToUserField string `json:"map_group_membership_to_user_field,omitempty"`

	// map object class to any or all
	// Enum: [all any]
	MapObjectClassToAnyOrAll string `json:"map_object_class_to_any_or_all,omitempty"`

	// map user membership to group field
	MapUserMembershipToGroupField bool `json:"map_user_membership_to_group_field,omitempty"`

	// map user membership use dn
	MapUserMembershipUseDn bool `json:"map_user_membership_use_dn,omitempty"`

	// object classes
	// Example: group
	ObjectClasses string `json:"object_classes,omitempty"`

	// recursive lookups
	RecursiveLookups bool `json:"recursive_lookups,omitempty"`

	// search base
	// Example: DC=Company,DC=com
	SearchBase string `json:"search_base,omitempty"`

	// search scope
	// Enum: [All Subtrees First Level Only]
	SearchScope string `json:"search_scope,omitempty"`

	// use dn
	UseDn bool `json:"use_dn,omitempty"`

	// user group membership stored in
	// Enum: [user object group object]
	UserGroupMembershipStoredIn string `json:"user_group_membership_stored_in,omitempty"`

	// user group membership use ldap compare
	UserGroupMembershipUseLdapCompare bool `json:"user_group_membership_use_ldap_compare,omitempty"`

	// username
	// Example: sAMAccountName
	Username string `json:"username,omitempty"`
}

// Validate validates this ldap server mappings for users user group membership mappings
func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMapObjectClassToAnyOrAll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserGroupMembershipStoredIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ldapServerMappingsForUsersUserGroupMembershipMappingsTypeMapObjectClassToAnyOrAllPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerMappingsForUsersUserGroupMembershipMappingsTypeMapObjectClassToAnyOrAllPropEnum = append(ldapServerMappingsForUsersUserGroupMembershipMappingsTypeMapObjectClassToAnyOrAllPropEnum, v)
	}
}

const (

	// LdapServerMappingsForUsersUserGroupMembershipMappingsMapObjectClassToAnyOrAllAll captures enum value "all"
	LdapServerMappingsForUsersUserGroupMembershipMappingsMapObjectClassToAnyOrAllAll string = "all"

	// LdapServerMappingsForUsersUserGroupMembershipMappingsMapObjectClassToAnyOrAllAny captures enum value "any"
	LdapServerMappingsForUsersUserGroupMembershipMappingsMapObjectClassToAnyOrAllAny string = "any"
)

// prop value enum
func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) validateMapObjectClassToAnyOrAllEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerMappingsForUsersUserGroupMembershipMappingsTypeMapObjectClassToAnyOrAllPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) validateMapObjectClassToAnyOrAll(formats strfmt.Registry) error {
	if swag.IsZero(m.MapObjectClassToAnyOrAll) { // not required
		return nil
	}

	// value enum
	if err := m.validateMapObjectClassToAnyOrAllEnum("mappings_for_users"+"."+"user_group_membership_mappings"+"."+"map_object_class_to_any_or_all", "body", m.MapObjectClassToAnyOrAll); err != nil {
		return err
	}

	return nil
}

var ldapServerMappingsForUsersUserGroupMembershipMappingsTypeSearchScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["All Subtrees","First Level Only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerMappingsForUsersUserGroupMembershipMappingsTypeSearchScopePropEnum = append(ldapServerMappingsForUsersUserGroupMembershipMappingsTypeSearchScopePropEnum, v)
	}
}

const (

	// LdapServerMappingsForUsersUserGroupMembershipMappingsSearchScopeAllSubtrees captures enum value "All Subtrees"
	LdapServerMappingsForUsersUserGroupMembershipMappingsSearchScopeAllSubtrees string = "All Subtrees"

	// LdapServerMappingsForUsersUserGroupMembershipMappingsSearchScopeFirstLevelOnly captures enum value "First Level Only"
	LdapServerMappingsForUsersUserGroupMembershipMappingsSearchScopeFirstLevelOnly string = "First Level Only"
)

// prop value enum
func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) validateSearchScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerMappingsForUsersUserGroupMembershipMappingsTypeSearchScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) validateSearchScope(formats strfmt.Registry) error {
	if swag.IsZero(m.SearchScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateSearchScopeEnum("mappings_for_users"+"."+"user_group_membership_mappings"+"."+"search_scope", "body", m.SearchScope); err != nil {
		return err
	}

	return nil
}

var ldapServerMappingsForUsersUserGroupMembershipMappingsTypeUserGroupMembershipStoredInPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user object","group object"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerMappingsForUsersUserGroupMembershipMappingsTypeUserGroupMembershipStoredInPropEnum = append(ldapServerMappingsForUsersUserGroupMembershipMappingsTypeUserGroupMembershipStoredInPropEnum, v)
	}
}

const (

	// LdapServerMappingsForUsersUserGroupMembershipMappingsUserGroupMembershipStoredInUserObject captures enum value "user object"
	LdapServerMappingsForUsersUserGroupMembershipMappingsUserGroupMembershipStoredInUserObject string = "user object"

	// LdapServerMappingsForUsersUserGroupMembershipMappingsUserGroupMembershipStoredInGroupObject captures enum value "group object"
	LdapServerMappingsForUsersUserGroupMembershipMappingsUserGroupMembershipStoredInGroupObject string = "group object"
)

// prop value enum
func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) validateUserGroupMembershipStoredInEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerMappingsForUsersUserGroupMembershipMappingsTypeUserGroupMembershipStoredInPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) validateUserGroupMembershipStoredIn(formats strfmt.Registry) error {
	if swag.IsZero(m.UserGroupMembershipStoredIn) { // not required
		return nil
	}

	// value enum
	if err := m.validateUserGroupMembershipStoredInEnum("mappings_for_users"+"."+"user_group_membership_mappings"+"."+"user_group_membership_stored_in", "body", m.UserGroupMembershipStoredIn); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ldap server mappings for users user group membership mappings based on context it is used
func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServerMappingsForUsersUserGroupMembershipMappings) UnmarshalBinary(b []byte) error {
	var res LdapServerMappingsForUsersUserGroupMembershipMappings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LdapServerMappingsForUsersUserMappings ldap server mappings for users user mappings
//
// swagger:model LdapServerMappingsForUsersUserMappings
type LdapServerMappingsForUsersUserMappings struct {

	// append to email results
	// Example: company.com
	AppendToEmailResults string `json:"append_to_email_results,omitempty"`

	// map building
	// Example: streetAddress
	MapBuilding string `json:"map_building,omitempty"`

	// map department
	// Example: department
	MapDepartment string `json:"map_department,omitempty"`

	// map email address
	// Example: mail
	MapEmailAddress string `json:"map_email_address,omitempty"`

	// map object class to any or all
	// Enum: [all any]
	MapObjectClassToAnyOrAll string `json:"map_object_class_to_any_or_all,omitempty"`

	// map position
	// Example: title
	MapPosition string `json:"map_position,omitempty"`

	// map realname
	// Example: displayName
	MapRealname string `json:"map_realname,omitempty"`

	// map room
	// Example: room
	MapRoom string `json:"map_room,omitempty"`

	// map telephone
	// Example: telephoneNumber
	MapTelephone string `json:"map_telephone,omitempty"`

	// map user id
	// Example: uSNCreated
	MapUserID string `json:"map_user_id,omitempty"`

	// map user uuid
	// Example: objectGUID
	MapUserUUID string `json:"map_user_uuid,omitempty"`

	// map username
	// Example: sAMAccountName
	MapUsername string `json:"map_username,omitempty"`

	// object classes
	// Example: organizationalPerson, user
	ObjectClasses string `json:"object_classes,omitempty"`

	// search base
	// Example: DC=Company,DC=com
	SearchBase string `json:"search_base,omitempty"`

	// search scope
	// Enum: [All Subtrees First Level Only]
	SearchScope string `json:"search_scope,omitempty"`
}

// Validate validates this ldap server mappings for users user mappings
func (m *LdapServerMappingsForUsersUserMappings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMapObjectClassToAnyOrAll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ldapServerMappingsForUsersUserMappingsTypeMapObjectClassToAnyOrAllPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerMappingsForUsersUserMappingsTypeMapObjectClassToAnyOrAllPropEnum = append(ldapServerMappingsForUsersUserMappingsTypeMapObjectClassToAnyOrAllPropEnum, v)
	}
}

const (

	// LdapServerMappingsForUsersUserMappingsMapObjectClassToAnyOrAllAll captures enum value "all"
	LdapServerMappingsForUsersUserMappingsMapObjectClassToAnyOrAllAll string = "all"

	// LdapServerMappingsForUsersUserMappingsMapObjectClassToAnyOrAllAny captures enum value "any"
	LdapServerMappingsForUsersUserMappingsMapObjectClassToAnyOrAllAny string = "any"
)

// prop value enum
func (m *LdapServerMappingsForUsersUserMappings) validateMapObjectClassToAnyOrAllEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerMappingsForUsersUserMappingsTypeMapObjectClassToAnyOrAllPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerMappingsForUsersUserMappings) validateMapObjectClassToAnyOrAll(formats strfmt.Registry) error {
	if swag.IsZero(m.MapObjectClassToAnyOrAll) { // not required
		return nil
	}

	// value enum
	if err := m.validateMapObjectClassToAnyOrAllEnum("mappings_for_users"+"."+"user_mappings"+"."+"map_object_class_to_any_or_all", "body", m.MapObjectClassToAnyOrAll); err != nil {
		return err
	}

	return nil
}

var ldapServerMappingsForUsersUserMappingsTypeSearchScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["All Subtrees","First Level Only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ldapServerMappingsForUsersUserMappingsTypeSearchScopePropEnum = append(ldapServerMappingsForUsersUserMappingsTypeSearchScopePropEnum, v)
	}
}

const (

	// LdapServerMappingsForUsersUserMappingsSearchScopeAllSubtrees captures enum value "All Subtrees"
	LdapServerMappingsForUsersUserMappingsSearchScopeAllSubtrees string = "All Subtrees"

	// LdapServerMappingsForUsersUserMappingsSearchScopeFirstLevelOnly captures enum value "First Level Only"
	LdapServerMappingsForUsersUserMappingsSearchScopeFirstLevelOnly string = "First Level Only"
)

// prop value enum
func (m *LdapServerMappingsForUsersUserMappings) validateSearchScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ldapServerMappingsForUsersUserMappingsTypeSearchScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LdapServerMappingsForUsersUserMappings) validateSearchScope(formats strfmt.Registry) error {
	if swag.IsZero(m.SearchScope) { // not required
		return nil
	}

	// value enum
	if err := m.validateSearchScopeEnum("mappings_for_users"+"."+"user_mappings"+"."+"search_scope", "body", m.SearchScope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ldap server mappings for users user mappings based on context it is used
func (m *LdapServerMappingsForUsersUserMappings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LdapServerMappingsForUsersUserMappings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LdapServerMappingsForUsersUserMappings) UnmarshalBinary(b []byte) error {
	var res LdapServerMappingsForUsersUserMappings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}