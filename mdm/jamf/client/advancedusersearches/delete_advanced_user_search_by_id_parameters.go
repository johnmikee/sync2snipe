// Code generated by go-swagger; DO NOT EDIT.

package advancedusersearches

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

// NewDeleteAdvancedUserSearchByIDParams creates a new DeleteAdvancedUserSearchByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAdvancedUserSearchByIDParams() *DeleteAdvancedUserSearchByIDParams {
	return &DeleteAdvancedUserSearchByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAdvancedUserSearchByIDParamsWithTimeout creates a new DeleteAdvancedUserSearchByIDParams object
// with the ability to set a timeout on a request.
func NewDeleteAdvancedUserSearchByIDParamsWithTimeout(timeout time.Duration) *DeleteAdvancedUserSearchByIDParams {
	return &DeleteAdvancedUserSearchByIDParams{
		timeout: timeout,
	}
}

// NewDeleteAdvancedUserSearchByIDParamsWithContext creates a new DeleteAdvancedUserSearchByIDParams object
// with the ability to set a context for a request.
func NewDeleteAdvancedUserSearchByIDParamsWithContext(ctx context.Context) *DeleteAdvancedUserSearchByIDParams {
	return &DeleteAdvancedUserSearchByIDParams{
		Context: ctx,
	}
}

// NewDeleteAdvancedUserSearchByIDParamsWithHTTPClient creates a new DeleteAdvancedUserSearchByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAdvancedUserSearchByIDParamsWithHTTPClient(client *http.Client) *DeleteAdvancedUserSearchByIDParams {
	return &DeleteAdvancedUserSearchByIDParams{
		HTTPClient: client,
	}
}

/*
DeleteAdvancedUserSearchByIDParams contains all the parameters to send to the API endpoint

	for the delete advanced user search by Id operation.

	Typically these are written to a http.Request.
*/
type DeleteAdvancedUserSearchByIDParams struct {

	/* ID.

	   ID value to filter by
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete advanced user search by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdvancedUserSearchByIDParams) WithDefaults() *DeleteAdvancedUserSearchByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete advanced user search by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAdvancedUserSearchByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete advanced user search by Id params
func (o *DeleteAdvancedUserSearchByIDParams) WithTimeout(timeout time.Duration) *DeleteAdvancedUserSearchByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete advanced user search by Id params
func (o *DeleteAdvancedUserSearchByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete advanced user search by Id params
func (o *DeleteAdvancedUserSearchByIDParams) WithContext(ctx context.Context) *DeleteAdvancedUserSearchByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete advanced user search by Id params
func (o *DeleteAdvancedUserSearchByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete advanced user search by Id params
func (o *DeleteAdvancedUserSearchByIDParams) WithHTTPClient(client *http.Client) *DeleteAdvancedUserSearchByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete advanced user search by Id params
func (o *DeleteAdvancedUserSearchByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete advanced user search by Id params
func (o *DeleteAdvancedUserSearchByIDParams) WithID(id int64) *DeleteAdvancedUserSearchByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete advanced user search by Id params
func (o *DeleteAdvancedUserSearchByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAdvancedUserSearchByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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