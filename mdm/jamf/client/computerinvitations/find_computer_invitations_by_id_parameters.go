// Code generated by go-swagger; DO NOT EDIT.

package computerinvitations

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

// NewFindComputerInvitationsByIDParams creates a new FindComputerInvitationsByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerInvitationsByIDParams() *FindComputerInvitationsByIDParams {
	return &FindComputerInvitationsByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerInvitationsByIDParamsWithTimeout creates a new FindComputerInvitationsByIDParams object
// with the ability to set a timeout on a request.
func NewFindComputerInvitationsByIDParamsWithTimeout(timeout time.Duration) *FindComputerInvitationsByIDParams {
	return &FindComputerInvitationsByIDParams{
		timeout: timeout,
	}
}

// NewFindComputerInvitationsByIDParamsWithContext creates a new FindComputerInvitationsByIDParams object
// with the ability to set a context for a request.
func NewFindComputerInvitationsByIDParamsWithContext(ctx context.Context) *FindComputerInvitationsByIDParams {
	return &FindComputerInvitationsByIDParams{
		Context: ctx,
	}
}

// NewFindComputerInvitationsByIDParamsWithHTTPClient creates a new FindComputerInvitationsByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerInvitationsByIDParamsWithHTTPClient(client *http.Client) *FindComputerInvitationsByIDParams {
	return &FindComputerInvitationsByIDParams{
		HTTPClient: client,
	}
}

/*
FindComputerInvitationsByIDParams contains all the parameters to send to the API endpoint

	for the find computer invitations by Id operation.

	Typically these are written to a http.Request.
*/
type FindComputerInvitationsByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer invitations by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerInvitationsByIDParams) WithDefaults() *FindComputerInvitationsByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer invitations by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerInvitationsByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer invitations by Id params
func (o *FindComputerInvitationsByIDParams) WithTimeout(timeout time.Duration) *FindComputerInvitationsByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer invitations by Id params
func (o *FindComputerInvitationsByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer invitations by Id params
func (o *FindComputerInvitationsByIDParams) WithContext(ctx context.Context) *FindComputerInvitationsByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer invitations by Id params
func (o *FindComputerInvitationsByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer invitations by Id params
func (o *FindComputerInvitationsByIDParams) WithHTTPClient(client *http.Client) *FindComputerInvitationsByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer invitations by Id params
func (o *FindComputerInvitationsByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find computer invitations by Id params
func (o *FindComputerInvitationsByIDParams) WithID(id int64) *FindComputerInvitationsByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find computer invitations by Id params
func (o *FindComputerInvitationsByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerInvitationsByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
