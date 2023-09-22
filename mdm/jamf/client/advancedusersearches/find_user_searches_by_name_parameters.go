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
)

// NewFindUserSearchesByNameParams creates a new FindUserSearchesByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindUserSearchesByNameParams() *FindUserSearchesByNameParams {
	return &FindUserSearchesByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindUserSearchesByNameParamsWithTimeout creates a new FindUserSearchesByNameParams object
// with the ability to set a timeout on a request.
func NewFindUserSearchesByNameParamsWithTimeout(timeout time.Duration) *FindUserSearchesByNameParams {
	return &FindUserSearchesByNameParams{
		timeout: timeout,
	}
}

// NewFindUserSearchesByNameParamsWithContext creates a new FindUserSearchesByNameParams object
// with the ability to set a context for a request.
func NewFindUserSearchesByNameParamsWithContext(ctx context.Context) *FindUserSearchesByNameParams {
	return &FindUserSearchesByNameParams{
		Context: ctx,
	}
}

// NewFindUserSearchesByNameParamsWithHTTPClient creates a new FindUserSearchesByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindUserSearchesByNameParamsWithHTTPClient(client *http.Client) *FindUserSearchesByNameParams {
	return &FindUserSearchesByNameParams{
		HTTPClient: client,
	}
}

/*
FindUserSearchesByNameParams contains all the parameters to send to the API endpoint

	for the find user searches by name operation.

	Typically these are written to a http.Request.
*/
type FindUserSearchesByNameParams struct {

	/* Name.

	   Name to filter by
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find user searches by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUserSearchesByNameParams) WithDefaults() *FindUserSearchesByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find user searches by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindUserSearchesByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find user searches by name params
func (o *FindUserSearchesByNameParams) WithTimeout(timeout time.Duration) *FindUserSearchesByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find user searches by name params
func (o *FindUserSearchesByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find user searches by name params
func (o *FindUserSearchesByNameParams) WithContext(ctx context.Context) *FindUserSearchesByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find user searches by name params
func (o *FindUserSearchesByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find user searches by name params
func (o *FindUserSearchesByNameParams) WithHTTPClient(client *http.Client) *FindUserSearchesByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find user searches by name params
func (o *FindUserSearchesByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the find user searches by name params
func (o *FindUserSearchesByNameParams) WithName(name string) *FindUserSearchesByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the find user searches by name params
func (o *FindUserSearchesByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *FindUserSearchesByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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