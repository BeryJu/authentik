// Code generated by go-swagger; DO NOT EDIT.

package sources

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

// NewSourcesOauthListParams creates a new SourcesOauthListParams object
// with the default values initialized.
func NewSourcesOauthListParams() *SourcesOauthListParams {
	var ()
	return &SourcesOauthListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSourcesOauthListParamsWithTimeout creates a new SourcesOauthListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSourcesOauthListParamsWithTimeout(timeout time.Duration) *SourcesOauthListParams {
	var ()
	return &SourcesOauthListParams{

		timeout: timeout,
	}
}

// NewSourcesOauthListParamsWithContext creates a new SourcesOauthListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSourcesOauthListParamsWithContext(ctx context.Context) *SourcesOauthListParams {
	var ()
	return &SourcesOauthListParams{

		Context: ctx,
	}
}

// NewSourcesOauthListParamsWithHTTPClient creates a new SourcesOauthListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSourcesOauthListParamsWithHTTPClient(client *http.Client) *SourcesOauthListParams {
	var ()
	return &SourcesOauthListParams{
		HTTPClient: client,
	}
}

/*SourcesOauthListParams contains all the parameters to send to the API endpoint
for the sources oauth list operation typically these are written to a http.Request
*/
type SourcesOauthListParams struct {

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

// WithTimeout adds the timeout to the sources oauth list params
func (o *SourcesOauthListParams) WithTimeout(timeout time.Duration) *SourcesOauthListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sources oauth list params
func (o *SourcesOauthListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sources oauth list params
func (o *SourcesOauthListParams) WithContext(ctx context.Context) *SourcesOauthListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sources oauth list params
func (o *SourcesOauthListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sources oauth list params
func (o *SourcesOauthListParams) WithHTTPClient(client *http.Client) *SourcesOauthListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sources oauth list params
func (o *SourcesOauthListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the sources oauth list params
func (o *SourcesOauthListParams) WithLimit(limit *int64) *SourcesOauthListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the sources oauth list params
func (o *SourcesOauthListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the sources oauth list params
func (o *SourcesOauthListParams) WithOffset(offset *int64) *SourcesOauthListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the sources oauth list params
func (o *SourcesOauthListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the sources oauth list params
func (o *SourcesOauthListParams) WithOrdering(ordering *string) *SourcesOauthListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the sources oauth list params
func (o *SourcesOauthListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithSearch adds the search to the sources oauth list params
func (o *SourcesOauthListParams) WithSearch(search *string) *SourcesOauthListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the sources oauth list params
func (o *SourcesOauthListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *SourcesOauthListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
