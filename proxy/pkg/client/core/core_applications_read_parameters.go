// Code generated by go-swagger; DO NOT EDIT.

package core

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

// NewCoreApplicationsReadParams creates a new CoreApplicationsReadParams object
// with the default values initialized.
func NewCoreApplicationsReadParams() *CoreApplicationsReadParams {
	var ()
	return &CoreApplicationsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCoreApplicationsReadParamsWithTimeout creates a new CoreApplicationsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCoreApplicationsReadParamsWithTimeout(timeout time.Duration) *CoreApplicationsReadParams {
	var ()
	return &CoreApplicationsReadParams{

		timeout: timeout,
	}
}

// NewCoreApplicationsReadParamsWithContext creates a new CoreApplicationsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewCoreApplicationsReadParamsWithContext(ctx context.Context) *CoreApplicationsReadParams {
	var ()
	return &CoreApplicationsReadParams{

		Context: ctx,
	}
}

// NewCoreApplicationsReadParamsWithHTTPClient creates a new CoreApplicationsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCoreApplicationsReadParamsWithHTTPClient(client *http.Client) *CoreApplicationsReadParams {
	var ()
	return &CoreApplicationsReadParams{
		HTTPClient: client,
	}
}

/*CoreApplicationsReadParams contains all the parameters to send to the API endpoint
for the core applications read operation typically these are written to a http.Request
*/
type CoreApplicationsReadParams struct {

	/*PbmUUID
	  A UUID string identifying this Application.

	*/
	PbmUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the core applications read params
func (o *CoreApplicationsReadParams) WithTimeout(timeout time.Duration) *CoreApplicationsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the core applications read params
func (o *CoreApplicationsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the core applications read params
func (o *CoreApplicationsReadParams) WithContext(ctx context.Context) *CoreApplicationsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the core applications read params
func (o *CoreApplicationsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the core applications read params
func (o *CoreApplicationsReadParams) WithHTTPClient(client *http.Client) *CoreApplicationsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the core applications read params
func (o *CoreApplicationsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPbmUUID adds the pbmUUID to the core applications read params
func (o *CoreApplicationsReadParams) WithPbmUUID(pbmUUID strfmt.UUID) *CoreApplicationsReadParams {
	o.SetPbmUUID(pbmUUID)
	return o
}

// SetPbmUUID adds the pbmUuid to the core applications read params
func (o *CoreApplicationsReadParams) SetPbmUUID(pbmUUID strfmt.UUID) {
	o.PbmUUID = pbmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CoreApplicationsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pbm_uuid
	if err := r.SetPathParam("pbm_uuid", o.PbmUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
