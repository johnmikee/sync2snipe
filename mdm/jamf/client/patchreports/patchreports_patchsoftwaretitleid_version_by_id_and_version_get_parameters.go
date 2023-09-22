// Code generated by go-swagger; DO NOT EDIT.

package patchreports

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

// NewPatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams creates a new PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams() *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	return &PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParamsWithTimeout creates a new PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams object
// with the ability to set a timeout on a request.
func NewPatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParamsWithTimeout(timeout time.Duration) *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	return &PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams{
		timeout: timeout,
	}
}

// NewPatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParamsWithContext creates a new PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams object
// with the ability to set a context for a request.
func NewPatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParamsWithContext(ctx context.Context) *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	return &PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams{
		Context: ctx,
	}
}

// NewPatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParamsWithHTTPClient creates a new PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParamsWithHTTPClient(client *http.Client) *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	return &PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams{
		HTTPClient: client,
	}
}

/*
PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams contains all the parameters to send to the API endpoint

	for the patchreports patchsoftwaretitleid version by Id and version get operation.

	Typically these are written to a http.Request.
*/
type PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams struct {

	/* ID.

	   Patch software title ID to filter by
	*/
	ID string

	/* Version.

	   Version number to filter by
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patchreports patchsoftwaretitleid version by Id and version get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) WithDefaults() *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patchreports patchsoftwaretitleid version by Id and version get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) WithTimeout(timeout time.Duration) *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) WithContext(ctx context.Context) *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) WithHTTPClient(client *http.Client) *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) WithID(id string) *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) WithVersion(version string) *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patchreports patchsoftwaretitleid version by Id and version get params
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchreportsPatchsoftwaretitleidVersionByIDAndVersionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
