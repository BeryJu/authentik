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

// NewStagesUserWriteDeleteParams creates a new StagesUserWriteDeleteParams object
// with the default values initialized.
func NewStagesUserWriteDeleteParams() *StagesUserWriteDeleteParams {
	var ()
	return &StagesUserWriteDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStagesUserWriteDeleteParamsWithTimeout creates a new StagesUserWriteDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStagesUserWriteDeleteParamsWithTimeout(timeout time.Duration) *StagesUserWriteDeleteParams {
	var ()
	return &StagesUserWriteDeleteParams{

		timeout: timeout,
	}
}

// NewStagesUserWriteDeleteParamsWithContext creates a new StagesUserWriteDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewStagesUserWriteDeleteParamsWithContext(ctx context.Context) *StagesUserWriteDeleteParams {
	var ()
	return &StagesUserWriteDeleteParams{

		Context: ctx,
	}
}

// NewStagesUserWriteDeleteParamsWithHTTPClient creates a new StagesUserWriteDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStagesUserWriteDeleteParamsWithHTTPClient(client *http.Client) *StagesUserWriteDeleteParams {
	var ()
	return &StagesUserWriteDeleteParams{
		HTTPClient: client,
	}
}

/*StagesUserWriteDeleteParams contains all the parameters to send to the API endpoint
for the stages user write delete operation typically these are written to a http.Request
*/
type StagesUserWriteDeleteParams struct {

	/*StageUUID
	  A UUID string identifying this User Write Stage.

	*/
	StageUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stages user write delete params
func (o *StagesUserWriteDeleteParams) WithTimeout(timeout time.Duration) *StagesUserWriteDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stages user write delete params
func (o *StagesUserWriteDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stages user write delete params
func (o *StagesUserWriteDeleteParams) WithContext(ctx context.Context) *StagesUserWriteDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stages user write delete params
func (o *StagesUserWriteDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stages user write delete params
func (o *StagesUserWriteDeleteParams) WithHTTPClient(client *http.Client) *StagesUserWriteDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stages user write delete params
func (o *StagesUserWriteDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStageUUID adds the stageUUID to the stages user write delete params
func (o *StagesUserWriteDeleteParams) WithStageUUID(stageUUID strfmt.UUID) *StagesUserWriteDeleteParams {
	o.SetStageUUID(stageUUID)
	return o
}

// SetStageUUID adds the stageUuid to the stages user write delete params
func (o *StagesUserWriteDeleteParams) SetStageUUID(stageUUID strfmt.UUID) {
	o.StageUUID = stageUUID
}

// WriteToRequest writes these params to a swagger request
func (o *StagesUserWriteDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
