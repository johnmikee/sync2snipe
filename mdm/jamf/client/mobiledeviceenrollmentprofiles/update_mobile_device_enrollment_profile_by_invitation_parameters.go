// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceenrollmentprofiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewUpdateMobileDeviceEnrollmentProfileByInvitationParams creates a new UpdateMobileDeviceEnrollmentProfileByInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMobileDeviceEnrollmentProfileByInvitationParams() *UpdateMobileDeviceEnrollmentProfileByInvitationParams {
	return &UpdateMobileDeviceEnrollmentProfileByInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMobileDeviceEnrollmentProfileByInvitationParamsWithTimeout creates a new UpdateMobileDeviceEnrollmentProfileByInvitationParams object
// with the ability to set a timeout on a request.
func NewUpdateMobileDeviceEnrollmentProfileByInvitationParamsWithTimeout(timeout time.Duration) *UpdateMobileDeviceEnrollmentProfileByInvitationParams {
	return &UpdateMobileDeviceEnrollmentProfileByInvitationParams{
		timeout: timeout,
	}
}

// NewUpdateMobileDeviceEnrollmentProfileByInvitationParamsWithContext creates a new UpdateMobileDeviceEnrollmentProfileByInvitationParams object
// with the ability to set a context for a request.
func NewUpdateMobileDeviceEnrollmentProfileByInvitationParamsWithContext(ctx context.Context) *UpdateMobileDeviceEnrollmentProfileByInvitationParams {
	return &UpdateMobileDeviceEnrollmentProfileByInvitationParams{
		Context: ctx,
	}
}

// NewUpdateMobileDeviceEnrollmentProfileByInvitationParamsWithHTTPClient creates a new UpdateMobileDeviceEnrollmentProfileByInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMobileDeviceEnrollmentProfileByInvitationParamsWithHTTPClient(client *http.Client) *UpdateMobileDeviceEnrollmentProfileByInvitationParams {
	return &UpdateMobileDeviceEnrollmentProfileByInvitationParams{
		HTTPClient: client,
	}
}

/*
UpdateMobileDeviceEnrollmentProfileByInvitationParams contains all the parameters to send to the API endpoint

	for the update mobile device enrollment profile by invitation operation.

	Typically these are written to a http.Request.
*/
type UpdateMobileDeviceEnrollmentProfileByInvitationParams struct {

	/* Invitation.

	   Invitation value to filter by
	*/
	Invitation string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update mobile device enrollment profile by invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) WithDefaults() *UpdateMobileDeviceEnrollmentProfileByInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update mobile device enrollment profile by invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update mobile device enrollment profile by invitation params
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) WithTimeout(timeout time.Duration) *UpdateMobileDeviceEnrollmentProfileByInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update mobile device enrollment profile by invitation params
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update mobile device enrollment profile by invitation params
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) WithContext(ctx context.Context) *UpdateMobileDeviceEnrollmentProfileByInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update mobile device enrollment profile by invitation params
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update mobile device enrollment profile by invitation params
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) WithHTTPClient(client *http.Client) *UpdateMobileDeviceEnrollmentProfileByInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update mobile device enrollment profile by invitation params
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitation adds the invitation to the update mobile device enrollment profile by invitation params
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) WithInvitation(invitation string) *UpdateMobileDeviceEnrollmentProfileByInvitationParams {
	o.SetInvitation(invitation)
	return o
}

// SetInvitation adds the invitation to the update mobile device enrollment profile by invitation params
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) SetInvitation(invitation string) {
	o.Invitation = invitation
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMobileDeviceEnrollmentProfileByInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invitation
	if err := r.SetPathParam("invitation", o.Invitation); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
