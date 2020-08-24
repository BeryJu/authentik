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

// NewProvidersProxyListParams creates a new ProvidersProxyListParams object
// with the default values initialized.
func NewProvidersProxyListParams() *ProvidersProxyListParams {
	var ()
	return &ProvidersProxyListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProvidersProxyListParamsWithTimeout creates a new ProvidersProxyListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProvidersProxyListParamsWithTimeout(timeout time.Duration) *ProvidersProxyListParams {
	var ()
	return &ProvidersProxyListParams{

		timeout: timeout,
	}
}

// NewProvidersProxyListParamsWithContext creates a new ProvidersProxyListParams object
// with the default values initialized, and the ability to set a context for a request
func NewProvidersProxyListParamsWithContext(ctx context.Context) *ProvidersProxyListParams {
	var ()
	return &ProvidersProxyListParams{

		Context: ctx,
	}
}

// NewProvidersProxyListParamsWithHTTPClient creates a new ProvidersProxyListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProvidersProxyListParamsWithHTTPClient(client *http.Client) *ProvidersProxyListParams {
	var ()
	return &ProvidersProxyListParams{
		HTTPClient: client,
	}
}

/*ProvidersProxyListParams contains all the parameters to send to the API endpoint
for the providers proxy list operation typically these are written to a http.Request
*/
type ProvidersProxyListParams struct {

	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Ordering
	  Which field to use when ordering the results.

	*/
	Ordering *string
	/*Search
	  A search term.

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the providers proxy list params
func (o *ProvidersProxyListParams) WithTimeout(timeout time.Duration) *ProvidersProxyListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the providers proxy list params
func (o *ProvidersProxyListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the providers proxy list params
func (o *ProvidersProxyListParams) WithContext(ctx context.Context) *ProvidersProxyListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the providers proxy list params
func (o *ProvidersProxyListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the providers proxy list params
func (o *ProvidersProxyListParams) WithHTTPClient(client *http.Client) *ProvidersProxyListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the providers proxy list params
func (o *ProvidersProxyListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the providers proxy list params
func (o *ProvidersProxyListParams) WithLimit(limit *int64) *ProvidersProxyListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the providers proxy list params
func (o *ProvidersProxyListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the providers proxy list params
func (o *ProvidersProxyListParams) WithOffset(offset *int64) *ProvidersProxyListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the providers proxy list params
func (o *ProvidersProxyListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the providers proxy list params
func (o *ProvidersProxyListParams) WithOrdering(ordering *string) *ProvidersProxyListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the providers proxy list params
func (o *ProvidersProxyListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithSearch adds the search to the providers proxy list params
func (o *ProvidersProxyListParams) WithSearch(search *string) *ProvidersProxyListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the providers proxy list params
func (o *ProvidersProxyListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *ProvidersProxyListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string
		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {
			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}

	}

	if o.Search != nil {

		// query param search
		var qrSearch string
		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {
			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
