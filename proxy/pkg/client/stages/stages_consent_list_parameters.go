// Code generated by go-swagger; DO NOT EDIT.

package stages

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

// NewStagesConsentListParams creates a new StagesConsentListParams object
// with the default values initialized.
func NewStagesConsentListParams() *StagesConsentListParams {
	var ()
	return &StagesConsentListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStagesConsentListParamsWithTimeout creates a new StagesConsentListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStagesConsentListParamsWithTimeout(timeout time.Duration) *StagesConsentListParams {
	var ()
	return &StagesConsentListParams{

		timeout: timeout,
	}
}

// NewStagesConsentListParamsWithContext creates a new StagesConsentListParams object
// with the default values initialized, and the ability to set a context for a request
func NewStagesConsentListParamsWithContext(ctx context.Context) *StagesConsentListParams {
	var ()
	return &StagesConsentListParams{

		Context: ctx,
	}
}

// NewStagesConsentListParamsWithHTTPClient creates a new StagesConsentListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStagesConsentListParamsWithHTTPClient(client *http.Client) *StagesConsentListParams {
	var ()
	return &StagesConsentListParams{
		HTTPClient: client,
	}
}

/*StagesConsentListParams contains all the parameters to send to the API endpoint
for the stages consent list operation typically these are written to a http.Request
*/
type StagesConsentListParams struct {

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

// WithTimeout adds the timeout to the stages consent list params
func (o *StagesConsentListParams) WithTimeout(timeout time.Duration) *StagesConsentListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stages consent list params
func (o *StagesConsentListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stages consent list params
func (o *StagesConsentListParams) WithContext(ctx context.Context) *StagesConsentListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stages consent list params
func (o *StagesConsentListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stages consent list params
func (o *StagesConsentListParams) WithHTTPClient(client *http.Client) *StagesConsentListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stages consent list params
func (o *StagesConsentListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the stages consent list params
func (o *StagesConsentListParams) WithLimit(limit *int64) *StagesConsentListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the stages consent list params
func (o *StagesConsentListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the stages consent list params
func (o *StagesConsentListParams) WithOffset(offset *int64) *StagesConsentListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the stages consent list params
func (o *StagesConsentListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the stages consent list params
func (o *StagesConsentListParams) WithOrdering(ordering *string) *StagesConsentListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the stages consent list params
func (o *StagesConsentListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithSearch adds the search to the stages consent list params
func (o *StagesConsentListParams) WithSearch(search *string) *StagesConsentListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the stages consent list params
func (o *StagesConsentListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *StagesConsentListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
