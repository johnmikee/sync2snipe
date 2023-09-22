// Code generated by go-swagger; DO NOT EDIT.

package healthcarelistener

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

// NewUpdateHealthCareListenerByIDParams creates a new UpdateHealthCareListenerByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateHealthCareListenerByIDParams() *UpdateHealthCareListenerByIDParams {
	return &UpdateHealthCareListenerByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHealthCareListenerByIDParamsWithTimeout creates a new UpdateHealthCareListenerByIDParams object
// with the ability to set a timeout on a request.
func NewUpdateHealthCareListenerByIDParamsWithTimeout(timeout time.Duration) *UpdateHealthCareListenerByIDParams {
	return &UpdateHealthCareListenerByIDParams{
		timeout: timeout,
	}
}

// NewUpdateHealthCareListenerByIDParamsWithContext creates a new UpdateHealthCareListenerByIDParams object
// with the ability to set a context for a request.
func NewUpdateHealthCareListenerByIDParamsWithContext(ctx context.Context) *UpdateHealthCareListenerByIDParams {
	return &UpdateHealthCareListenerByIDParams{
		Context: ctx,
	}
}

// NewUpdateHealthCareListenerByIDParamsWithHTTPClient creates a new UpdateHealthCareListenerByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateHealthCareListenerByIDParamsWithHTTPClient(client *http.Client) *UpdateHealthCareListenerByIDParams {
	return &UpdateHealthCareListenerByIDParams{
		HTTPClient: client,
	}
}

/*
UpdateHealthCareListenerByIDParams contains all the parameters to send to the API endpoint

	for the update health care listener by Id operation.

	Typically these are written to a http.Request.
*/
type UpdateHealthCareListenerByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update health care listener by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHealthCareListenerByIDParams) WithDefaults() *UpdateHealthCareListenerByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update health care listener by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateHealthCareListenerByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update health care listener by Id params
func (o *UpdateHealthCareListenerByIDParams) WithTimeout(timeout time.Duration) *UpdateHealthCareListenerByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update health care listener by Id params
func (o *UpdateHealthCareListenerByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update health care listener by Id params
func (o *UpdateHealthCareListenerByIDParams) WithContext(ctx context.Context) *UpdateHealthCareListenerByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update health care listener by Id params
func (o *UpdateHealthCareListenerByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update health care listener by Id params
func (o *UpdateHealthCareListenerByIDParams) WithHTTPClient(client *http.Client) *UpdateHealthCareListenerByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update health care listener by Id params
func (o *UpdateHealthCareListenerByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update health care listener by Id params
func (o *UpdateHealthCareListenerByIDParams) WithID(id int64) *UpdateHealthCareListenerByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update health care listener by Id params
func (o *UpdateHealthCareListenerByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHealthCareListenerByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
