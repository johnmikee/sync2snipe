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

// PatchManagementSoftwareTitle patch management software title
//
// swagger:model patch_management_software_title
type PatchManagementSoftwareTitle struct {

	// category
	Category *CategoryObject `json:"category,omitempty"`

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// name
	// Example: Google Chrome
	// Required: true
	Name *string `json:"name"`

	// name id
	// Example: GoogleChrome
	NameID string `json:"name_id,omitempty"`

	// notifications
	Notifications *PatchManagementSoftwareTitleNotifications `json:"notifications,omitempty"`

	// site
	Site *SiteObject `json:"site,omitempty"`

	// versions
	Versions []*PatchManagementSoftwareTitleVersionsItems0 `json:"versions"`
}

// Validate validates this patch management software title
func (m *PatchManagementSoftwareTitle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSite(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchManagementSoftwareTitle) validateCategory(formats strfmt.Registry) error {
	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if m.Category != nil {
		if err := m.Category.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *PatchManagementSoftwareTitle) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PatchManagementSoftwareTitle) validateNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifications) { // not required
		return nil
	}

	if m.Notifications != nil {
		if err := m.Notifications.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *PatchManagementSoftwareTitle) validateSite(formats strfmt.Registry) error {
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

func (m *PatchManagementSoftwareTitle) validateVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	for i := 0; i < len(m.Versions); i++ {
		if swag.IsZero(m.Versions[i]) { // not required
			continue
		}

		if m.Versions[i] != nil {
			if err := m.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this patch management software title based on the context it is used
func (m *PatchManagementSoftwareTitle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCategory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSite(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchManagementSoftwareTitle) contextValidateCategory(ctx context.Context, formats strfmt.Registry) error {

	if m.Category != nil {

		if swag.IsZero(m.Category) { // not required
			return nil
		}

		if err := m.Category.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("category")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("category")
			}
			return err
		}
	}

	return nil
}

func (m *PatchManagementSoftwareTitle) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	if m.Notifications != nil {

		if swag.IsZero(m.Notifications) { // not required
			return nil
		}

		if err := m.Notifications.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("notifications")
			}
			return err
		}
	}

	return nil
}

func (m *PatchManagementSoftwareTitle) contextValidateSite(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PatchManagementSoftwareTitle) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Versions); i++ {

		if m.Versions[i] != nil {

			if swag.IsZero(m.Versions[i]) { // not required
				return nil
			}

			if err := m.Versions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchManagementSoftwareTitle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchManagementSoftwareTitle) UnmarshalBinary(b []byte) error {
	var res PatchManagementSoftwareTitle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchManagementSoftwareTitleNotifications patch management software title notifications
//
// swagger:model PatchManagementSoftwareTitleNotifications
type PatchManagementSoftwareTitleNotifications struct {

	// email notification
	EmailNotification *bool `json:"email_notification,omitempty"`

	// jss notification
	JssNotification *bool `json:"jss_notification,omitempty"`
}

// Validate validates this patch management software title notifications
func (m *PatchManagementSoftwareTitleNotifications) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch management software title notifications based on context it is used
func (m *PatchManagementSoftwareTitleNotifications) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchManagementSoftwareTitleNotifications) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchManagementSoftwareTitleNotifications) UnmarshalBinary(b []byte) error {
	var res PatchManagementSoftwareTitleNotifications
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchManagementSoftwareTitleVersionsItems0 patch management software title versions items0
//
// swagger:model PatchManagementSoftwareTitleVersionsItems0
type PatchManagementSoftwareTitleVersionsItems0 struct {

	// version
	Version *PatchManagementSoftwareTitleVersionsItems0Version `json:"version,omitempty"`
}

// Validate validates this patch management software title versions items0
func (m *PatchManagementSoftwareTitleVersionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchManagementSoftwareTitleVersionsItems0) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch management software title versions items0 based on the context it is used
func (m *PatchManagementSoftwareTitleVersionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchManagementSoftwareTitleVersionsItems0) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {

		if swag.IsZero(m.Version) { // not required
			return nil
		}

		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchManagementSoftwareTitleVersionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchManagementSoftwareTitleVersionsItems0) UnmarshalBinary(b []byte) error {
	var res PatchManagementSoftwareTitleVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchManagementSoftwareTitleVersionsItems0Version patch management software title versions items0 version
//
// swagger:model PatchManagementSoftwareTitleVersionsItems0Version
type PatchManagementSoftwareTitleVersionsItems0Version struct {

	// package
	Package *PatchManagementSoftwareTitleVersionsItems0VersionPackage `json:"package,omitempty"`

	// software version
	// Example: 64.0.3282.119
	SoftwareVersion string `json:"software_version,omitempty"`
}

// Validate validates this patch management software title versions items0 version
func (m *PatchManagementSoftwareTitleVersionsItems0Version) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchManagementSoftwareTitleVersionsItems0Version) validatePackage(formats strfmt.Registry) error {
	if swag.IsZero(m.Package) { // not required
		return nil
	}

	if m.Package != nil {
		if err := m.Package.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version" + "." + "package")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version" + "." + "package")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch management software title versions items0 version based on the context it is used
func (m *PatchManagementSoftwareTitleVersionsItems0Version) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePackage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchManagementSoftwareTitleVersionsItems0Version) contextValidatePackage(ctx context.Context, formats strfmt.Registry) error {

	if m.Package != nil {

		if swag.IsZero(m.Package) { // not required
			return nil
		}

		if err := m.Package.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version" + "." + "package")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version" + "." + "package")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchManagementSoftwareTitleVersionsItems0Version) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchManagementSoftwareTitleVersionsItems0Version) UnmarshalBinary(b []byte) error {
	var res PatchManagementSoftwareTitleVersionsItems0Version
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchManagementSoftwareTitleVersionsItems0VersionPackage patch management software title versions items0 version package
//
// swagger:model PatchManagementSoftwareTitleVersionsItems0VersionPackage
type PatchManagementSoftwareTitleVersionsItems0VersionPackage struct {

	// id
	// Example: 1
	ID *int64 `json:"id,omitempty"`

	// name
	// Example: Google Chrome.dmg
	Name string `json:"name,omitempty"`
}

// Validate validates this patch management software title versions items0 version package
func (m *PatchManagementSoftwareTitleVersionsItems0VersionPackage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch management software title versions items0 version package based on context it is used
func (m *PatchManagementSoftwareTitleVersionsItems0VersionPackage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchManagementSoftwareTitleVersionsItems0VersionPackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchManagementSoftwareTitleVersionsItems0VersionPackage) UnmarshalBinary(b []byte) error {
	var res PatchManagementSoftwareTitleVersionsItems0VersionPackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
