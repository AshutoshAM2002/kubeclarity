// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetVulnerabilitiesVulIDPkgIDParams creates a new GetVulnerabilitiesVulIDPkgIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVulnerabilitiesVulIDPkgIDParams() *GetVulnerabilitiesVulIDPkgIDParams {
	return &GetVulnerabilitiesVulIDPkgIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVulnerabilitiesVulIDPkgIDParamsWithTimeout creates a new GetVulnerabilitiesVulIDPkgIDParams object
// with the ability to set a timeout on a request.
func NewGetVulnerabilitiesVulIDPkgIDParamsWithTimeout(timeout time.Duration) *GetVulnerabilitiesVulIDPkgIDParams {
	return &GetVulnerabilitiesVulIDPkgIDParams{
		timeout: timeout,
	}
}

// NewGetVulnerabilitiesVulIDPkgIDParamsWithContext creates a new GetVulnerabilitiesVulIDPkgIDParams object
// with the ability to set a context for a request.
func NewGetVulnerabilitiesVulIDPkgIDParamsWithContext(ctx context.Context) *GetVulnerabilitiesVulIDPkgIDParams {
	return &GetVulnerabilitiesVulIDPkgIDParams{
		Context: ctx,
	}
}

// NewGetVulnerabilitiesVulIDPkgIDParamsWithHTTPClient creates a new GetVulnerabilitiesVulIDPkgIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVulnerabilitiesVulIDPkgIDParamsWithHTTPClient(client *http.Client) *GetVulnerabilitiesVulIDPkgIDParams {
	return &GetVulnerabilitiesVulIDPkgIDParams{
		HTTPClient: client,
	}
}

/* GetVulnerabilitiesVulIDPkgIDParams contains all the parameters to send to the API endpoint
   for the get vulnerabilities vul ID pkg ID operation.

   Typically these are written to a http.Request.
*/
type GetVulnerabilitiesVulIDPkgIDParams struct {

	// PkgID.
	PkgID string

	// VulID.
	VulID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vulnerabilities vul ID pkg ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVulnerabilitiesVulIDPkgIDParams) WithDefaults() *GetVulnerabilitiesVulIDPkgIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vulnerabilities vul ID pkg ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVulnerabilitiesVulIDPkgIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) WithTimeout(timeout time.Duration) *GetVulnerabilitiesVulIDPkgIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) WithContext(ctx context.Context) *GetVulnerabilitiesVulIDPkgIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) WithHTTPClient(client *http.Client) *GetVulnerabilitiesVulIDPkgIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPkgID adds the pkgID to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) WithPkgID(pkgID string) *GetVulnerabilitiesVulIDPkgIDParams {
	o.SetPkgID(pkgID)
	return o
}

// SetPkgID adds the pkgId to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) SetPkgID(pkgID string) {
	o.PkgID = pkgID
}

// WithVulID adds the vulID to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) WithVulID(vulID string) *GetVulnerabilitiesVulIDPkgIDParams {
	o.SetVulID(vulID)
	return o
}

// SetVulID adds the vulId to the get vulnerabilities vul ID pkg ID params
func (o *GetVulnerabilitiesVulIDPkgIDParams) SetVulID(vulID string) {
	o.VulID = vulID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVulnerabilitiesVulIDPkgIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pkg_id
	if err := r.SetPathParam("pkg_id", o.PkgID); err != nil {
		return err
	}

	// path param vul_id
	if err := r.SetPathParam("vul_id", o.VulID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}