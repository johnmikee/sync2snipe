// Code generated by go-swagger; DO NOT EDIT.

package vppinvitations

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

// NewFindInvitationByIDParams creates a new FindInvitationByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindInvitationByIDParams() *FindInvitationByIDParams {
	return &FindInvitationByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindInvitationByIDParamsWithTimeout creates a new FindInvitationByIDParams object
// with the ability to set a timeout on a request.
func NewFindInvitationByIDParamsWithTimeout(timeout time.Duration) *FindInvitationByIDParams {
	return &FindInvitationByIDParams{
		timeout: timeout,
	}
}

// NewFindInvitationByIDParamsWithContext creates a new FindInvitationByIDParams object
// with the ability to set a context for a request.
func NewFindInvitationByIDParamsWithContext(ctx context.Context) *FindInvitationByIDParams {
	return &FindInvitationByIDParams{
		Context: ctx,
	}
}

// NewFindInvitationByIDParamsWithHTTPClient creates a new FindInvitationByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindInvitationByIDParamsWithHTTPClient(client *http.Client) *FindInvitationByIDParams {
	return &FindInvitationByIDParams{
		HTTPClient: client,
	}
}

/*
FindInvitationByIDParams contains all the parameters to send to the API endpoint

	for the find invitation by Id operation.

	Typically these are written to a http.Request.
*/
type FindInvitationByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find invitation by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindInvitationByIDParams) WithDefaults() *FindInvitationByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find invitation by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindInvitationByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find invitation by Id params
func (o *FindInvitationByIDParams) WithTimeout(timeout time.Duration) *FindInvitationByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find invitation by Id params
func (o *FindInvitationByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find invitation by Id params
func (o *FindInvitationByIDParams) WithContext(ctx context.Context) *FindInvitationByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find invitation by Id params
func (o *FindInvitationByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find invitation by Id params
func (o *FindInvitationByIDParams) WithHTTPClient(client *http.Client) *FindInvitationByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find invitation by Id params
func (o *FindInvitationByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find invitation by Id params
func (o *FindInvitationByIDParams) WithID(id int64) *FindInvitationByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find invitation by Id params
func (o *FindInvitationByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindInvitationByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}