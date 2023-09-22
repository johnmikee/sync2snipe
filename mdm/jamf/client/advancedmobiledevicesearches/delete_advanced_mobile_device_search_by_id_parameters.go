// Code generated by go-swagger; DO NOT EDIT.

package advancedmobiledevicesearches

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

// NewDeleteAdvancedMobileDeviceSearchByIDParams creates a new DeleteAdvancedMobileDeviceSearchByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAdvancedMobileDeviceSearchByIDParams() *DeleteAdvancedMobileDeviceSearchByIDParams {
	return &DeleteAdvancedMobileDeviceSearchByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAdvancedMobileDeviceSearchByIDParamsWithTimeout creates a new DeleteAdvancedMobileDeviceSearchByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAdvancedMobileDeviceSearchByIDParamsWithTimeout(timeout time.Duration) *DeleteAdvancedMobileDeviceSearchByIDParams {
	return &DeleteAdvancedMobileDeviceSearchByIDParams{
		timeout: timeout,
	}
}

// NewDeleteAdvancedMobileDeviceSearchByIDParamsWithContext creates a new DeleteAdvancedMobileDeviceSearchByIDParams object
// with the ability to set a context for a request.
func NewDeleteAdvancedMobileDeviceSearchByIDParamsWithContext(ctx context.Context) *DeleteAdvancedMobileDeviceSearchByIDParams {
	return &DeleteAdvancedMobileDeviceSearchByIDParams{
		Context: ctx,
	}
}

// NewDeleteAdvancedMobileDeviceSearchByIDParamsWithHTTPClient creates a new DeleteAdvancedMobileDeviceSearchByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAdvancedMobileDeviceSearchByIDParamsWithHTTPClient(client *http.Client) *DeleteAdvancedMobileDeviceSearchByIDParams {
	return &DeleteAdvancedMobileDeviceSearchByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteAdvancedMobileDeviceSearchByIDParams contains all the parameters to send to the API endpoint

	for the delete advanced mobile device search by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteAdvancedMobileDeviceSearchByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete advanced mobile device search by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) WithDefaults() *DeleteAdvancedMobileDeviceSearchByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete advanced mobile device search by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete advanced mobile device search by Id params
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) WithTimeout(timeout time.Duration) *DeleteAdvancedMobileDeviceSearchByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete advanced mobile device search by Id params
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete advanced mobile device search by Id params
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) WithContext(ctx context.Context) *DeleteAdvancedMobileDeviceSearchByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete advanced mobile device search by Id params
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete advanced mobile device search by Id params
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) WithHTTPClient(client *http.Client) *DeleteAdvancedMobileDeviceSearchByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete advanced mobile device search by Id params
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete advanced mobile device search by Id params
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) WithID(id int64) *DeleteAdvancedMobileDeviceSearchByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete advanced mobile device search by Id params
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAdvancedMobileDeviceSearchByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
