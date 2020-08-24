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
)

// NewStagesInvitationReadParams creates a new StagesInvitationReadParams object
// with the default values initialized.
func NewStagesInvitationReadParams() *StagesInvitationReadParams {
	var ()
	return &StagesInvitationReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStagesInvitationReadParamsWithTimeout creates a new StagesInvitationReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStagesInvitationReadParamsWithTimeout(timeout time.Duration) *StagesInvitationReadParams {
	var ()
	return &StagesInvitationReadParams{

		timeout: timeout,
	}
}

// NewStagesInvitationReadParamsWithContext creates a new StagesInvitationReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewStagesInvitationReadParamsWithContext(ctx context.Context) *StagesInvitationReadParams {
	var ()
	return &StagesInvitationReadParams{

		Context: ctx,
	}
}

// NewStagesInvitationReadParamsWithHTTPClient creates a new StagesInvitationReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStagesInvitationReadParamsWithHTTPClient(client *http.Client) *StagesInvitationReadParams {
	var ()
	return &StagesInvitationReadParams{
		HTTPClient: client,
	}
}

/*StagesInvitationReadParams contains all the parameters to send to the API endpoint
for the stages invitation read operation typically these are written to a http.Request
*/
type StagesInvitationReadParams struct {

	/*StageUUID
	  A UUID string identifying this Invitation Stage.

	*/
	StageUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stages invitation read params
func (o *StagesInvitationReadParams) WithTimeout(timeout time.Duration) *StagesInvitationReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stages invitation read params
func (o *StagesInvitationReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stages invitation read params
func (o *StagesInvitationReadParams) WithContext(ctx context.Context) *StagesInvitationReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stages invitation read params
func (o *StagesInvitationReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stages invitation read params
func (o *StagesInvitationReadParams) WithHTTPClient(client *http.Client) *StagesInvitationReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stages invitation read params
func (o *StagesInvitationReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStageUUID adds the stageUUID to the stages invitation read params
func (o *StagesInvitationReadParams) WithStageUUID(stageUUID strfmt.UUID) *StagesInvitationReadParams {
	o.SetStageUUID(stageUUID)
	return o
}

// SetStageUUID adds the stageUuid to the stages invitation read params
func (o *StagesInvitationReadParams) SetStageUUID(stageUUID strfmt.UUID) {
	o.StageUUID = stageUUID
}

// WriteToRequest writes these params to a swagger request
func (o *StagesInvitationReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stage_uuid
	if err := r.SetPathParam("stage_uuid", o.StageUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
