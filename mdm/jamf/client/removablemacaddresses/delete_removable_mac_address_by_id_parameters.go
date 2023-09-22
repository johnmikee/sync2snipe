// Code generated by go-swagger; DO NOT EDIT.

package removablemacaddresses

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

// NewDeleteRemovableMacAddressByIDParams creates a new DeleteRemovableMacAddressByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRemovableMacAddressByIDParams() *DeleteRemovableMacAddressByIDParams {
	return &DeleteRemovableMacAddressByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRemovableMacAddressByIDParamsWithTimeout creates a new DeleteRemovableMacAddressByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteRemovableMacAddressByIDParamsWithTimeout(timeout time.Duration) *DeleteRemovableMacAddressByIDParams {
	return &DeleteRemovableMacAddressByIDParams{
		timeout: timeout,
	}
}

// NewDeleteRemovableMacAddressByIDParamsWithContext creates a new DeleteRemovableMacAddressByIDParams object
// with the ability to set a context for a request.
func NewDeleteRemovableMacAddressByIDParamsWithContext(ctx context.Context) *DeleteRemovableMacAddressByIDParams {
	return &DeleteRemovableMacAddressByIDParams{
		Context: ctx,
	}
}

// NewDeleteRemovableMacAddressByIDParamsWithHTTPClient creates a new DeleteRemovableMacAddressByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRemovableMacAddressByIDParamsWithHTTPClient(client *http.Client) *DeleteRemovableMacAddressByIDParams {
	return &DeleteRemovableMacAddressByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteRemovableMacAddressByIDParams contains all the parameters to send to the API endpoint

	for the delete removable mac address by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteRemovableMacAddressByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete removable mac address by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRemovableMacAddressByIDParams) WithDefaults() *DeleteRemovableMacAddressByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete removable mac address by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRemovableMacAddressByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete removable mac address by Id params
func (o *DeleteRemovableMacAddressByIDParams) WithTimeout(timeout time.Duration) *DeleteRemovableMacAddressByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete removable mac address by Id params
func (o *DeleteRemovableMacAddressByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete removable mac address by Id params
func (o *DeleteRemovableMacAddressByIDParams) WithContext(ctx context.Context) *DeleteRemovableMacAddressByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete removable mac address by Id params
func (o *DeleteRemovableMacAddressByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete removable mac address by Id params
func (o *DeleteRemovableMacAddressByIDParams) WithHTTPClient(client *http.Client) *DeleteRemovableMacAddressByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete removable mac address by Id params
func (o *DeleteRemovableMacAddressByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete removable mac address by Id params
func (o *DeleteRemovableMacAddressByIDParams) WithID(id int64) *DeleteRemovableMacAddressByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete removable mac address by Id params
func (o *DeleteRemovableMacAddressByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRemovableMacAddressByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
