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
	"github.com/go-openapi/swag"
)

// NewDeleteSoftwareTitlesByIDParams creates a new DeleteSoftwareTitlesByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSoftwareTitlesByIDParams() *DeleteSoftwareTitlesByIDParams {
	return &DeleteSoftwareTitlesByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSoftwareTitlesByIDParamsWithTimeout creates a new DeleteSoftwareTitlesByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteSoftwareTitlesByIDParamsWithTimeout(timeout time.Duration) *DeleteSoftwareTitlesByIDParams {
	return &DeleteSoftwareTitlesByIDParams{
		timeout: timeout,
	}
}

// NewDeleteSoftwareTitlesByIDParamsWithContext creates a new DeleteSoftwareTitlesByIDParams object
// with the ability to set a context for a request.
func NewDeleteSoftwareTitlesByIDParamsWithContext(ctx context.Context) *DeleteSoftwareTitlesByIDParams {
	return &DeleteSoftwareTitlesByIDParams{
		Context: ctx,
	}
}

// NewDeleteSoftwareTitlesByIDParamsWithHTTPClient creates a new DeleteSoftwareTitlesByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSoftwareTitlesByIDParamsWithHTTPClient(client *http.Client) *DeleteSoftwareTitlesByIDParams {
	return &DeleteSoftwareTitlesByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteSoftwareTitlesByIDParams contains all the parameters to send to the API endpoint

	for the delete software titles by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteSoftwareTitlesByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete software titles by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSoftwareTitlesByIDParams) WithDefaults() *DeleteSoftwareTitlesByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete software titles by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSoftwareTitlesByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete software titles by Id params
func (o *DeleteSoftwareTitlesByIDParams) WithTimeout(timeout time.Duration) *DeleteSoftwareTitlesByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete software titles by Id params
func (o *DeleteSoftwareTitlesByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete software titles by Id params
func (o *DeleteSoftwareTitlesByIDParams) WithContext(ctx context.Context) *DeleteSoftwareTitlesByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete software titles by Id params
func (o *DeleteSoftwareTitlesByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete software titles by Id params
func (o *DeleteSoftwareTitlesByIDParams) WithHTTPClient(client *http.Client) *DeleteSoftwareTitlesByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete software titles by Id params
func (o *DeleteSoftwareTitlesByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete software titles by Id params
func (o *DeleteSoftwareTitlesByIDParams) WithID(id int64) *DeleteSoftwareTitlesByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete software titles by Id params
func (o *DeleteSoftwareTitlesByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSoftwareTitlesByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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