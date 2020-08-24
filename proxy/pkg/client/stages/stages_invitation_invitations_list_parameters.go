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

// NewStagesInvitationInvitationsListParams creates a new StagesInvitationInvitationsListParams object
// with the default values initialized.
func NewStagesInvitationInvitationsListParams() *StagesInvitationInvitationsListParams {
	var ()
	return &StagesInvitationInvitationsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStagesInvitationInvitationsListParamsWithTimeout creates a new StagesInvitationInvitationsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStagesInvitationInvitationsListParamsWithTimeout(timeout time.Duration) *StagesInvitationInvitationsListParams {
	var ()
	return &StagesInvitationInvitationsListParams{

		timeout: timeout,
	}
}

// NewStagesInvitationInvitationsListParamsWithContext creates a new StagesInvitationInvitationsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewStagesInvitationInvitationsListParamsWithContext(ctx context.Context) *StagesInvitationInvitationsListParams {
	var ()
	return &StagesInvitationInvitationsListParams{

		Context: ctx,
	}
}

// NewStagesInvitationInvitationsListParamsWithHTTPClient creates a new StagesInvitationInvitationsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStagesInvitationInvitationsListParamsWithHTTPClient(client *http.Client) *StagesInvitationInvitationsListParams {
	var ()
	return &StagesInvitationInvitationsListParams{
		HTTPClient: client,
	}
}

/*StagesInvitationInvitationsListParams contains all the parameters to send to the API endpoint
for the stages invitation invitations list operation typically these are written to a http.Request
*/
type StagesInvitationInvitationsListParams struct {

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

// WithTimeout adds the timeout to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) WithTimeout(timeout time.Duration) *StagesInvitationInvitationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) WithContext(ctx context.Context) *StagesInvitationInvitationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) WithHTTPClient(client *http.Client) *StagesInvitationInvitationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) WithLimit(limit *int64) *StagesInvitationInvitationsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) WithOffset(offset *int64) *StagesInvitationInvitationsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) WithOrdering(ordering *string) *StagesInvitationInvitationsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithSearch adds the search to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) WithSearch(search *string) *StagesInvitationInvitationsListParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the stages invitation invitations list params
func (o *StagesInvitationInvitationsListParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *StagesInvitationInvitationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
