// Code generated by go-swagger; DO NOT EDIT.

package providers

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

// NewProvidersProxyReadParams creates a new ProvidersProxyReadParams object
// with the default values initialized.
func NewProvidersProxyReadParams() *ProvidersProxyReadParams {
	var ()
	return &ProvidersProxyReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProvidersProxyReadParamsWithTimeout creates a new ProvidersProxyReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProvidersProxyReadParamsWithTimeout(timeout time.Duration) *ProvidersProxyReadParams {
	var ()
	return &ProvidersProxyReadParams{

		timeout: timeout,
	}
}

// NewProvidersProxyReadParamsWithContext creates a new ProvidersProxyReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewProvidersProxyReadParamsWithContext(ctx context.Context) *ProvidersProxyReadParams {
	var ()
	return &ProvidersProxyReadParams{

		Context: ctx,
	}
}

// NewProvidersProxyReadParamsWithHTTPClient creates a new ProvidersProxyReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProvidersProxyReadParamsWithHTTPClient(client *http.Client) *ProvidersProxyReadParams {
	var ()
	return &ProvidersProxyReadParams{
		HTTPClient: client,
	}
}

/*ProvidersProxyReadParams contains all the parameters to send to the API endpoint
for the providers proxy read operation typically these are written to a http.Request
*/
type ProvidersProxyReadParams struct {

	/*ID
	  A unique integer value identifying this Proxy Provider.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the providers proxy read params
func (o *ProvidersProxyReadParams) WithTimeout(timeout time.Duration) *ProvidersProxyReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the providers proxy read params
func (o *ProvidersProxyReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the providers proxy read params
func (o *ProvidersProxyReadParams) WithContext(ctx context.Context) *ProvidersProxyReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the providers proxy read params
func (o *ProvidersProxyReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the providers proxy read params
func (o *ProvidersProxyReadParams) WithHTTPClient(client *http.Client) *ProvidersProxyReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the providers proxy read params
func (o *ProvidersProxyReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the providers proxy read params
func (o *ProvidersProxyReadParams) WithID(id int64) *ProvidersProxyReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the providers proxy read params
func (o *ProvidersProxyReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProvidersProxyReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
