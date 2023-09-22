// Code generated by go-swagger; DO NOT EDIT.

package patches

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

// NewCreatePatchSoftwareTitlesByIDParams creates a new CreatePatchSoftwareTitlesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePatchSoftwareTitlesByIDParams() *CreatePatchSoftwareTitlesByIDParams {
	return &CreatePatchSoftwareTitlesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePatchSoftwareTitlesByIDParamsWithTimeout creates a new CreatePatchSoftwareTitlesByIDParams object
// with the ability to set a timeout on a request.
func NewCreatePatchSoftwareTitlesByIDParamsWithTimeout(timeout time.Duration) *CreatePatchSoftwareTitlesByIDParams {
	return &CreatePatchSoftwareTitlesByIDParams{
		timeout: timeout,
	}
}

// NewCreatePatchSoftwareTitlesByIDParamsWithContext creates a new CreatePatchSoftwareTitlesByIDParams object
// with the ability to set a context for a request.
func NewCreatePatchSoftwareTitlesByIDParamsWithContext(ctx context.Context) *CreatePatchSoftwareTitlesByIDParams {
	return &CreatePatchSoftwareTitlesByIDParams{
		Context: ctx,
	}
}

// NewCreatePatchSoftwareTitlesByIDParamsWithHTTPClient creates a new CreatePatchSoftwareTitlesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePatchSoftwareTitlesByIDParamsWithHTTPClient(client *http.Client) *CreatePatchSoftwareTitlesByIDParams {
	return &CreatePatchSoftwareTitlesByIDParams{
		HTTPClient: client,
	}
}

/*
CreatePatchSoftwareTitlesByIDParams contains all the parameters to send to the API endpoint

	for the create patch software titles by Id operation.

	Typically these are written to a http.Request.
*/
type CreatePatchSoftwareTitlesByIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create patch software titles by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePatchSoftwareTitlesByIDParams) WithDefaults() *CreatePatchSoftwareTitlesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create patch software titles by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePatchSoftwareTitlesByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create patch software titles by Id params
func (o *CreatePatchSoftwareTitlesByIDParams) WithTimeout(timeout time.Duration) *CreatePatchSoftwareTitlesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create patch software titles by Id params
func (o *CreatePatchSoftwareTitlesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create patch software titles by Id params
func (o *CreatePatchSoftwareTitlesByIDParams) WithContext(ctx context.Context) *CreatePatchSoftwareTitlesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create patch software titles by Id params
func (o *CreatePatchSoftwareTitlesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create patch software titles by Id params
func (o *CreatePatchSoftwareTitlesByIDParams) WithHTTPClient(client *http.Client) *CreatePatchSoftwareTitlesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create patch software titles by Id params
func (o *CreatePatchSoftwareTitlesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the create patch software titles by Id params
func (o *CreatePatchSoftwareTitlesByIDParams) WithID(id string) *CreatePatchSoftwareTitlesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create patch software titles by Id params
func (o *CreatePatchSoftwareTitlesByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePatchSoftwareTitlesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
