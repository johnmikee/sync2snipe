// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

// NewCreateWebhookByIDParams creates a new CreateWebhookByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateWebhookByIDParams() *CreateWebhookByIDParams {
	return &CreateWebhookByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWebhookByIDParamsWithTimeout creates a new CreateWebhookByIDParams object
// with the ability to set a timeout on a request.
func NewCreateWebhookByIDParamsWithTimeout(timeout time.Duration) *CreateWebhookByIDParams {
	return &CreateWebhookByIDParams{
		timeout: timeout,
	}
}

// NewCreateWebhookByIDParamsWithContext creates a new CreateWebhookByIDParams object
// with the ability to set a context for a request.
func NewCreateWebhookByIDParamsWithContext(ctx context.Context) *CreateWebhookByIDParams {
	return &CreateWebhookByIDParams{
		Context: ctx,
	}
}

// NewCreateWebhookByIDParamsWithHTTPClient creates a new CreateWebhookByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateWebhookByIDParamsWithHTTPClient(client *http.Client) *CreateWebhookByIDParams {
	return &CreateWebhookByIDParams{
		HTTPClient: client,
	}
}

/*
CreateWebhookByIDParams contains all the parameters to send to the API endpoint

	for the create webhook by Id operation.

	Typically these are written to a http.Request.
*/
type CreateWebhookByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create webhook by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWebhookByIDParams) WithDefaults() *CreateWebhookByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create webhook by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWebhookByIDParams) SetDefaults() {
	var (
		idDefault = int64(0)
	)

	val := CreateWebhookByIDParams{
		ID: idDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create webhook by Id params
func (o *CreateWebhookByIDParams) WithTimeout(timeout time.Duration) *CreateWebhookByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create webhook by Id params
func (o *CreateWebhookByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create webhook by Id params
func (o *CreateWebhookByIDParams) WithContext(ctx context.Context) *CreateWebhookByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create webhook by Id params
func (o *CreateWebhookByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create webhook by Id params
func (o *CreateWebhookByIDParams) WithHTTPClient(client *http.Client) *CreateWebhookByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create webhook by Id params
func (o *CreateWebhookByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create webhook by Id params
func (o *CreateWebhookByIDParams) WithID(id int64) *CreateWebhookByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create webhook by Id params
func (o *CreateWebhookByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWebhookByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
