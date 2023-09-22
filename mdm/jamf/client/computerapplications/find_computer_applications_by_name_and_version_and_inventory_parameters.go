// Code generated by go-swagger; DO NOT EDIT.

package computerapplications

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

// NewFindComputerApplicationsByNameAndVersionAndInventoryParams creates a new FindComputerApplicationsByNameAndVersionAndInventoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerApplicationsByNameAndVersionAndInventoryParams() *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	return &FindComputerApplicationsByNameAndVersionAndInventoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerApplicationsByNameAndVersionAndInventoryParamsWithTimeout creates a new FindComputerApplicationsByNameAndVersionAndInventoryParams object
// with the ability to set a timeout on a request.
func NewFindComputerApplicationsByNameAndVersionAndInventoryParamsWithTimeout(timeout time.Duration) *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	return &FindComputerApplicationsByNameAndVersionAndInventoryParams{
		timeout: timeout,
	}
}

// NewFindComputerApplicationsByNameAndVersionAndInventoryParamsWithContext creates a new FindComputerApplicationsByNameAndVersionAndInventoryParams object
// with the ability to set a context for a request.
func NewFindComputerApplicationsByNameAndVersionAndInventoryParamsWithContext(ctx context.Context) *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	return &FindComputerApplicationsByNameAndVersionAndInventoryParams{
		Context: ctx,
	}
}

// NewFindComputerApplicationsByNameAndVersionAndInventoryParamsWithHTTPClient creates a new FindComputerApplicationsByNameAndVersionAndInventoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerApplicationsByNameAndVersionAndInventoryParamsWithHTTPClient(client *http.Client) *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	return &FindComputerApplicationsByNameAndVersionAndInventoryParams{
		HTTPClient: client,
	}
}

/*
FindComputerApplicationsByNameAndVersionAndInventoryParams contains all the parameters to send to the API endpoint

	for the find computer applications by name and version and inventory operation.

	Typically these are written to a http.Request.
*/
type FindComputerApplicationsByNameAndVersionAndInventoryParams struct {

	/* Application.

	   Application name to filter by
	*/
	Application string

	/* Inventory.

	   Inventory options
	*/
	Inventory string

	/* Version.

	   Version to filter by
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer applications by name and version and inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) WithDefaults() *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer applications by name and version and inventory params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) WithTimeout(timeout time.Duration) *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) WithContext(ctx context.Context) *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) WithHTTPClient(client *http.Client) *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplication adds the application to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) WithApplication(application string) *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	o.SetApplication(application)
	return o
}

// SetApplication adds the application to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) SetApplication(application string) {
	o.Application = application
}

// WithInventory adds the inventory to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) WithInventory(inventory string) *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	o.SetInventory(inventory)
	return o
}

// SetInventory adds the inventory to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) SetInventory(inventory string) {
	o.Inventory = inventory
}

// WithVersion adds the version to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) WithVersion(version string) *FindComputerApplicationsByNameAndVersionAndInventoryParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the find computer applications by name and version and inventory params
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerApplicationsByNameAndVersionAndInventoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application
	if err := r.SetPathParam("application", o.Application); err != nil {
		return err
	}

	// path param inventory
	if err := r.SetPathParam("inventory", o.Inventory); err != nil {
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
