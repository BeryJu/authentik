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

// NewStagesAllReadParams creates a new StagesAllReadParams object
// with the default values initialized.
func NewStagesAllReadParams() *StagesAllReadParams {
	var ()
	return &StagesAllReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStagesAllReadParamsWithTimeout creates a new StagesAllReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStagesAllReadParamsWithTimeout(timeout time.Duration) *StagesAllReadParams {
	var ()
	return &StagesAllReadParams{

		timeout: timeout,
	}
}

// NewStagesAllReadParamsWithContext creates a new StagesAllReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewStagesAllReadParamsWithContext(ctx context.Context) *StagesAllReadParams {
	var ()
	return &StagesAllReadParams{

		Context: ctx,
	}
}

// NewStagesAllReadParamsWithHTTPClient creates a new StagesAllReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStagesAllReadParamsWithHTTPClient(client *http.Client) *StagesAllReadParams {
	var ()
	return &StagesAllReadParams{
		HTTPClient: client,
	}
}

/*StagesAllReadParams contains all the parameters to send to the API endpoint
for the stages all read operation typically these are written to a http.Request
*/
type StagesAllReadParams struct {

	/*StageUUID
	  A UUID string identifying this stage.

	*/
	StageUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stages all read params
func (o *StagesAllReadParams) WithTimeout(timeout time.Duration) *StagesAllReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stages all read params
func (o *StagesAllReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stages all read params
func (o *StagesAllReadParams) WithContext(ctx context.Context) *StagesAllReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stages all read params
func (o *StagesAllReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stages all read params
func (o *StagesAllReadParams) WithHTTPClient(client *http.Client) *StagesAllReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stages all read params
func (o *StagesAllReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStageUUID adds the stageUUID to the stages all read params
func (o *StagesAllReadParams) WithStageUUID(stageUUID strfmt.UUID) *StagesAllReadParams {
	o.SetStageUUID(stageUUID)
	return o
}

// SetStageUUID adds the stageUuid to the stages all read params
func (o *StagesAllReadParams) SetStageUUID(stageUUID strfmt.UUID) {
	o.StageUUID = stageUUID
}

// WriteToRequest writes these params to a swagger request
func (o *StagesAllReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
