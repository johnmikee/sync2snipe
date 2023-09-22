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

// NewFindMobileDeviceEnrollmentProfilesByInvitationParams creates a new FindMobileDeviceEnrollmentProfilesByInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMobileDeviceEnrollmentProfilesByInvitationParams() *FindMobileDeviceEnrollmentProfilesByInvitationParams {
	return &FindMobileDeviceEnrollmentProfilesByInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMobileDeviceEnrollmentProfilesByInvitationParamsWithTimeout creates a new FindMobileDeviceEnrollmentProfilesByInvitationParams object
// with the ability to set a timeout on a request.
func NewFindMobileDeviceEnrollmentProfilesByInvitationParamsWithTimeout(timeout time.Duration) *FindMobileDeviceEnrollmentProfilesByInvitationParams {
	return &FindMobileDeviceEnrollmentProfilesByInvitationParams{
		timeout: timeout,
	}
}

// NewFindMobileDeviceEnrollmentProfilesByInvitationParamsWithContext creates a new FindMobileDeviceEnrollmentProfilesByInvitationParams object
// with the ability to set a context for a request.
func NewFindMobileDeviceEnrollmentProfilesByInvitationParamsWithContext(ctx context.Context) *FindMobileDeviceEnrollmentProfilesByInvitationParams {
	return &FindMobileDeviceEnrollmentProfilesByInvitationParams{
		Context: ctx,
	}
}

// NewFindMobileDeviceEnrollmentProfilesByInvitationParamsWithHTTPClient creates a new FindMobileDeviceEnrollmentProfilesByInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMobileDeviceEnrollmentProfilesByInvitationParamsWithHTTPClient(client *http.Client) *FindMobileDeviceEnrollmentProfilesByInvitationParams {
	return &FindMobileDeviceEnrollmentProfilesByInvitationParams{
		HTTPClient: client,
	}
}

/*
FindMobileDeviceEnrollmentProfilesByInvitationParams contains all the parameters to send to the API endpoint

	for the find mobile device enrollment profiles by invitation operation.

	Typically these are written to a http.Request.
*/
type FindMobileDeviceEnrollmentProfilesByInvitationParams struct {

	/* Invitation.

	   Invitation value to filter by
	*/
	Invitation string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find mobile device enrollment profiles by invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) WithDefaults() *FindMobileDeviceEnrollmentProfilesByInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find mobile device enrollment profiles by invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find mobile device enrollment profiles by invitation params
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) WithTimeout(timeout time.Duration) *FindMobileDeviceEnrollmentProfilesByInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find mobile device enrollment profiles by invitation params
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find mobile device enrollment profiles by invitation params
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) WithContext(ctx context.Context) *FindMobileDeviceEnrollmentProfilesByInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find mobile device enrollment profiles by invitation params
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find mobile device enrollment profiles by invitation params
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) WithHTTPClient(client *http.Client) *FindMobileDeviceEnrollmentProfilesByInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find mobile device enrollment profiles by invitation params
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitation adds the invitation to the find mobile device enrollment profiles by invitation params
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) WithInvitation(invitation string) *FindMobileDeviceEnrollmentProfilesByInvitationParams {
	o.SetInvitation(invitation)
	return o
}

// SetInvitation adds the invitation to the find mobile device enrollment profiles by invitation params
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) SetInvitation(invitation string) {
	o.Invitation = invitation
}

// WriteToRequest writes these params to a swagger request
func (o *FindMobileDeviceEnrollmentProfilesByInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
