// Code generated by go-swagger; DO NOT EDIT.

package mobiledeviceinvitations

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
	"github.com/go-openapi/swag"
)

// NewDeleteMobileDeviceInvitationByInvitationParams creates a new DeleteMobileDeviceInvitationByInvitationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMobileDeviceInvitationByInvitationParams() *DeleteMobileDeviceInvitationByInvitationParams {
	return &DeleteMobileDeviceInvitationByInvitationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMobileDeviceInvitationByInvitationParamsWithTimeout creates a new DeleteMobileDeviceInvitationByInvitationParams object
// with the ability to set a timeout on a request.
func NewDeleteMobileDeviceInvitationByInvitationParamsWithTimeout(timeout time.Duration) *DeleteMobileDeviceInvitationByInvitationParams {
	return &DeleteMobileDeviceInvitationByInvitationParams{
		timeout: timeout,
	}
}

// NewDeleteMobileDeviceInvitationByInvitationParamsWithContext creates a new DeleteMobileDeviceInvitationByInvitationParams object
// with the ability to set a context for a request.
func NewDeleteMobileDeviceInvitationByInvitationParamsWithContext(ctx context.Context) *DeleteMobileDeviceInvitationByInvitationParams {
	return &DeleteMobileDeviceInvitationByInvitationParams{
		Context: ctx,
	}
}

// NewDeleteMobileDeviceInvitationByInvitationParamsWithHTTPClient creates a new DeleteMobileDeviceInvitationByInvitationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMobileDeviceInvitationByInvitationParamsWithHTTPClient(client *http.Client) *DeleteMobileDeviceInvitationByInvitationParams {
	return &DeleteMobileDeviceInvitationByInvitationParams{
		HTTPClient: client,
	}
}

/*
DeleteMobileDeviceInvitationByInvitationParams contains all the parameters to send to the API endpoint

	for the delete mobile device invitation by invitation operation.

	Typically these are written to a http.Request.
*/
type DeleteMobileDeviceInvitationByInvitationParams struct {

	/* Invitation.

	   Invitation value to filter by
	*/
	Invitation int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete mobile device invitation by invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceInvitationByInvitationParams) WithDefaults() *DeleteMobileDeviceInvitationByInvitationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete mobile device invitation by invitation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMobileDeviceInvitationByInvitationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete mobile device invitation by invitation params
func (o *DeleteMobileDeviceInvitationByInvitationParams) WithTimeout(timeout time.Duration) *DeleteMobileDeviceInvitationByInvitationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete mobile device invitation by invitation params
func (o *DeleteMobileDeviceInvitationByInvitationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete mobile device invitation by invitation params
func (o *DeleteMobileDeviceInvitationByInvitationParams) WithContext(ctx context.Context) *DeleteMobileDeviceInvitationByInvitationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete mobile device invitation by invitation params
func (o *DeleteMobileDeviceInvitationByInvitationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete mobile device invitation by invitation params
func (o *DeleteMobileDeviceInvitationByInvitationParams) WithHTTPClient(client *http.Client) *DeleteMobileDeviceInvitationByInvitationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete mobile device invitation by invitation params
func (o *DeleteMobileDeviceInvitationByInvitationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInvitation adds the invitation to the delete mobile device invitation by invitation params
func (o *DeleteMobileDeviceInvitationByInvitationParams) WithInvitation(invitation int64) *DeleteMobileDeviceInvitationByInvitationParams {
	o.SetInvitation(invitation)
	return o
}

// SetInvitation adds the invitation to the delete mobile device invitation by invitation params
func (o *DeleteMobileDeviceInvitationByInvitationParams) SetInvitation(invitation int64) {
	o.Invitation = invitation
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMobileDeviceInvitationByInvitationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param invitation
	if err := r.SetPathParam("invitation", swag.FormatInt64(o.Invitation)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}