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
)

// NewDeleteRemovableMacAddressByNameParams creates a new DeleteRemovableMacAddressByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRemovableMacAddressByNameParams() *DeleteRemovableMacAddressByNameParams {
	return &DeleteRemovableMacAddressByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRemovableMacAddressByNameParamsWithTimeout creates a new DeleteRemovableMacAddressByNameParams object
// with the ability to set a timeout on a request.
func NewDeleteRemovableMacAddressByNameParamsWithTimeout(timeout time.Duration) *DeleteRemovableMacAddressByNameParams {
	return &DeleteRemovableMacAddressByNameParams{
		timeout: timeout,
	}
}

// NewDeleteRemovableMacAddressByNameParamsWithContext creates a new DeleteRemovableMacAddressByNameParams object
// with the ability to set a context for a request.
func NewDeleteRemovableMacAddressByNameParamsWithContext(ctx context.Context) *DeleteRemovableMacAddressByNameParams {
	return &DeleteRemovableMacAddressByNameParams{
		Context: ctx,
	}
}

// NewDeleteRemovableMacAddressByNameParamsWithHTTPClient creates a new DeleteRemovableMacAddressByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRemovableMacAddressByNameParamsWithHTTPClient(client *http.Client) *DeleteRemovableMacAddressByNameParams {
	return &DeleteRemovableMacAddressByNameParams{
		HTTPClient: client,
	}
}

/*
DeleteRemovableMacAddressByNameParams contains all the parameters to send to the API endpoint

	for the delete removable mac address by name operation.

	Typically these are written to a http.Request.
*/
type DeleteRemovableMacAddressByNameParams struct {

	/* Name.

	   Name value to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete removable mac address by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRemovableMacAddressByNameParams) WithDefaults() *DeleteRemovableMacAddressByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete removable mac address by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRemovableMacAddressByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete removable mac address by name params
func (o *DeleteRemovableMacAddressByNameParams) WithTimeout(timeout time.Duration) *DeleteRemovableMacAddressByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete removable mac address by name params
func (o *DeleteRemovableMacAddressByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete removable mac address by name params
func (o *DeleteRemovableMacAddressByNameParams) WithContext(ctx context.Context) *DeleteRemovableMacAddressByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete removable mac address by name params
func (o *DeleteRemovableMacAddressByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete removable mac address by name params
func (o *DeleteRemovableMacAddressByNameParams) WithHTTPClient(client *http.Client) *DeleteRemovableMacAddressByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete removable mac address by name params
func (o *DeleteRemovableMacAddressByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete removable mac address by name params
func (o *DeleteRemovableMacAddressByNameParams) WithName(name string) *DeleteRemovableMacAddressByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete removable mac address by name params
func (o *DeleteRemovableMacAddressByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRemovableMacAddressByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}