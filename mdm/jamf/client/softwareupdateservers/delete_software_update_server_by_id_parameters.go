// Code generated by go-swagger; DO NOT EDIT.

package softwareupdateservers

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

// NewDeleteSoftwareUpdateServerByIDParams creates a new DeleteSoftwareUpdateServerByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSoftwareUpdateServerByIDParams() *DeleteSoftwareUpdateServerByIDParams {
	return &DeleteSoftwareUpdateServerByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSoftwareUpdateServerByIDParamsWithTimeout creates a new DeleteSoftwareUpdateServerByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteSoftwareUpdateServerByIDParamsWithTimeout(timeout time.Duration) *DeleteSoftwareUpdateServerByIDParams {
	return &DeleteSoftwareUpdateServerByIDParams{
		timeout: timeout,
	}
}

// NewDeleteSoftwareUpdateServerByIDParamsWithContext creates a new DeleteSoftwareUpdateServerByIDParams object
// with the ability to set a context for a request.
func NewDeleteSoftwareUpdateServerByIDParamsWithContext(ctx context.Context) *DeleteSoftwareUpdateServerByIDParams {
	return &DeleteSoftwareUpdateServerByIDParams{
		Context: ctx,
	}
}

// NewDeleteSoftwareUpdateServerByIDParamsWithHTTPClient creates a new DeleteSoftwareUpdateServerByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSoftwareUpdateServerByIDParamsWithHTTPClient(client *http.Client) *DeleteSoftwareUpdateServerByIDParams {
	return &DeleteSoftwareUpdateServerByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteSoftwareUpdateServerByIDParams contains all the parameters to send to the API endpoint

	for the delete software update server by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteSoftwareUpdateServerByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete software update server by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSoftwareUpdateServerByIDParams) WithDefaults() *DeleteSoftwareUpdateServerByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete software update server by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSoftwareUpdateServerByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete software update server by Id params
func (o *DeleteSoftwareUpdateServerByIDParams) WithTimeout(timeout time.Duration) *DeleteSoftwareUpdateServerByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete software update server by Id params
func (o *DeleteSoftwareUpdateServerByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete software update server by Id params
func (o *DeleteSoftwareUpdateServerByIDParams) WithContext(ctx context.Context) *DeleteSoftwareUpdateServerByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete software update server by Id params
func (o *DeleteSoftwareUpdateServerByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete software update server by Id params
func (o *DeleteSoftwareUpdateServerByIDParams) WithHTTPClient(client *http.Client) *DeleteSoftwareUpdateServerByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete software update server by Id params
func (o *DeleteSoftwareUpdateServerByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete software update server by Id params
func (o *DeleteSoftwareUpdateServerByIDParams) WithID(id int64) *DeleteSoftwareUpdateServerByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete software update server by Id params
func (o *DeleteSoftwareUpdateServerByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSoftwareUpdateServerByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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