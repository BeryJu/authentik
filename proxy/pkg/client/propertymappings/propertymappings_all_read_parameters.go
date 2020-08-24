// Code generated by go-swagger; DO NOT EDIT.

package propertymappings

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

// NewPropertymappingsAllReadParams creates a new PropertymappingsAllReadParams object
// with the default values initialized.
func NewPropertymappingsAllReadParams() *PropertymappingsAllReadParams {
	var ()
	return &PropertymappingsAllReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPropertymappingsAllReadParamsWithTimeout creates a new PropertymappingsAllReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPropertymappingsAllReadParamsWithTimeout(timeout time.Duration) *PropertymappingsAllReadParams {
	var ()
	return &PropertymappingsAllReadParams{

		timeout: timeout,
	}
}

// NewPropertymappingsAllReadParamsWithContext creates a new PropertymappingsAllReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewPropertymappingsAllReadParamsWithContext(ctx context.Context) *PropertymappingsAllReadParams {
	var ()
	return &PropertymappingsAllReadParams{

		Context: ctx,
	}
}

// NewPropertymappingsAllReadParamsWithHTTPClient creates a new PropertymappingsAllReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPropertymappingsAllReadParamsWithHTTPClient(client *http.Client) *PropertymappingsAllReadParams {
	var ()
	return &PropertymappingsAllReadParams{
		HTTPClient: client,
	}
}

/*PropertymappingsAllReadParams contains all the parameters to send to the API endpoint
for the propertymappings all read operation typically these are written to a http.Request
*/
type PropertymappingsAllReadParams struct {

	/*PmUUID
	  A UUID string identifying this Property Mapping.

	*/
	PmUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the propertymappings all read params
func (o *PropertymappingsAllReadParams) WithTimeout(timeout time.Duration) *PropertymappingsAllReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the propertymappings all read params
func (o *PropertymappingsAllReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the propertymappings all read params
func (o *PropertymappingsAllReadParams) WithContext(ctx context.Context) *PropertymappingsAllReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the propertymappings all read params
func (o *PropertymappingsAllReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the propertymappings all read params
func (o *PropertymappingsAllReadParams) WithHTTPClient(client *http.Client) *PropertymappingsAllReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the propertymappings all read params
func (o *PropertymappingsAllReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPmUUID adds the pmUUID to the propertymappings all read params
func (o *PropertymappingsAllReadParams) WithPmUUID(pmUUID strfmt.UUID) *PropertymappingsAllReadParams {
	o.SetPmUUID(pmUUID)
	return o
}

// SetPmUUID adds the pmUuid to the propertymappings all read params
func (o *PropertymappingsAllReadParams) SetPmUUID(pmUUID strfmt.UUID) {
	o.PmUUID = pmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PropertymappingsAllReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pm_uuid
	if err := r.SetPathParam("pm_uuid", o.PmUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
