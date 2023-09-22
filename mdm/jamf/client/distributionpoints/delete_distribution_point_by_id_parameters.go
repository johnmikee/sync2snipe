// Code generated by go-swagger; DO NOT EDIT.

package distributionpoints

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

// NewDeleteDistributionPointByIDParams creates a new DeleteDistributionPointByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDistributionPointByIDParams() *DeleteDistributionPointByIDParams {
	return &DeleteDistributionPointByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDistributionPointByIDParamsWithTimeout creates a new DeleteDistributionPointByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDistributionPointByIDParamsWithTimeout(timeout time.Duration) *DeleteDistributionPointByIDParams {
	return &DeleteDistributionPointByIDParams{
		timeout: timeout,
	}
}

// NewDeleteDistributionPointByIDParamsWithContext creates a new DeleteDistributionPointByIDParams object
// with the ability to set a context for a request.
func NewDeleteDistributionPointByIDParamsWithContext(ctx context.Context) *DeleteDistributionPointByIDParams {
	return &DeleteDistributionPointByIDParams{
		Context: ctx,
	}
}

// NewDeleteDistributionPointByIDParamsWithHTTPClient creates a new DeleteDistributionPointByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDistributionPointByIDParamsWithHTTPClient(client *http.Client) *DeleteDistributionPointByIDParams {
	return &DeleteDistributionPointByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteDistributionPointByIDParams contains all the parameters to send to the API endpoint

	for the delete distribution point by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteDistributionPointByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete distribution point by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDistributionPointByIDParams) WithDefaults() *DeleteDistributionPointByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete distribution point by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDistributionPointByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete distribution point by Id params
func (o *DeleteDistributionPointByIDParams) WithTimeout(timeout time.Duration) *DeleteDistributionPointByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete distribution point by Id params
func (o *DeleteDistributionPointByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete distribution point by Id params
func (o *DeleteDistributionPointByIDParams) WithContext(ctx context.Context) *DeleteDistributionPointByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete distribution point by Id params
func (o *DeleteDistributionPointByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete distribution point by Id params
func (o *DeleteDistributionPointByIDParams) WithHTTPClient(client *http.Client) *DeleteDistributionPointByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete distribution point by Id params
func (o *DeleteDistributionPointByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete distribution point by Id params
func (o *DeleteDistributionPointByIDParams) WithID(id int64) *DeleteDistributionPointByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete distribution point by Id params
func (o *DeleteDistributionPointByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDistributionPointByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
