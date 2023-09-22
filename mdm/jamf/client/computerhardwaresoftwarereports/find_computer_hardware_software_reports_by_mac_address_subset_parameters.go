// Code generated by go-swagger; DO NOT EDIT.

package computerhardwaresoftwarereports

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

// NewFindComputerHardwareSoftwareReportsByMacAddressSubsetParams creates a new FindComputerHardwareSoftwareReportsByMacAddressSubsetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindComputerHardwareSoftwareReportsByMacAddressSubsetParams() *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	return &FindComputerHardwareSoftwareReportsByMacAddressSubsetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindComputerHardwareSoftwareReportsByMacAddressSubsetParamsWithTimeout creates a new FindComputerHardwareSoftwareReportsByMacAddressSubsetParams object
// with the ability to set a timeout on a request.
func NewFindComputerHardwareSoftwareReportsByMacAddressSubsetParamsWithTimeout(timeout time.Duration) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	return &FindComputerHardwareSoftwareReportsByMacAddressSubsetParams{
		timeout: timeout,
	}
}

// NewFindComputerHardwareSoftwareReportsByMacAddressSubsetParamsWithContext creates a new FindComputerHardwareSoftwareReportsByMacAddressSubsetParams object
// with the ability to set a context for a request.
func NewFindComputerHardwareSoftwareReportsByMacAddressSubsetParamsWithContext(ctx context.Context) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	return &FindComputerHardwareSoftwareReportsByMacAddressSubsetParams{
		Context: ctx,
	}
}

// NewFindComputerHardwareSoftwareReportsByMacAddressSubsetParamsWithHTTPClient creates a new FindComputerHardwareSoftwareReportsByMacAddressSubsetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindComputerHardwareSoftwareReportsByMacAddressSubsetParamsWithHTTPClient(client *http.Client) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	return &FindComputerHardwareSoftwareReportsByMacAddressSubsetParams{
		HTTPClient: client,
	}
}

/*
FindComputerHardwareSoftwareReportsByMacAddressSubsetParams contains all the parameters to send to the API endpoint

	for the find computer hardware software reports by mac address subset operation.

	Typically these are written to a http.Request.
*/
type FindComputerHardwareSoftwareReportsByMacAddressSubsetParams struct {

	/* EndDate.

	   End date (e.g. yyyy-mm-dd)

	   Format: date
	*/
	EndDate strfmt.Date

	/* Macaddress.

	   MAC address to filter by
	*/
	Macaddress string

	/* StartDate.

	   Start date (e.g. yyyy-mm-dd)

	   Format: date
	*/
	StartDate strfmt.Date

	/* Subset.

	   Subset to filter by
	*/
	Subset string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find computer hardware software reports by mac address subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) WithDefaults() *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find computer hardware software reports by mac address subset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) WithTimeout(timeout time.Duration) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) WithContext(ctx context.Context) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) WithHTTPClient(client *http.Client) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) WithEndDate(endDate strfmt.Date) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) SetEndDate(endDate strfmt.Date) {
	o.EndDate = endDate
}

// WithMacaddress adds the macaddress to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) WithMacaddress(macaddress string) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	o.SetMacaddress(macaddress)
	return o
}

// SetMacaddress adds the macaddress to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) SetMacaddress(macaddress string) {
	o.Macaddress = macaddress
}

// WithStartDate adds the startDate to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) WithStartDate(startDate strfmt.Date) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) SetStartDate(startDate strfmt.Date) {
	o.StartDate = startDate
}

// WithSubset adds the subset to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) WithSubset(subset string) *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams {
	o.SetSubset(subset)
	return o
}

// SetSubset adds the subset to the find computer hardware software reports by mac address subset params
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) SetSubset(subset string) {
	o.Subset = subset
}

// WriteToRequest writes these params to a swagger request
func (o *FindComputerHardwareSoftwareReportsByMacAddressSubsetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param end_date
	if err := r.SetPathParam("end_date", o.EndDate.String()); err != nil {
		return err
	}

	// path param macaddress
	if err := r.SetPathParam("macaddress", o.Macaddress); err != nil {
		return err
	}

	// path param start_date
	if err := r.SetPathParam("start_date", o.StartDate.String()); err != nil {
		return err
	}

	// path param subset
	if err := r.SetPathParam("subset", o.Subset); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}