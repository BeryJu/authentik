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

// NewStagesUserLogoutDeleteParams creates a new StagesUserLogoutDeleteParams object
// with the default values initialized.
func NewStagesUserLogoutDeleteParams() *StagesUserLogoutDeleteParams {
	var ()
	return &StagesUserLogoutDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStagesUserLogoutDeleteParamsWithTimeout creates a new StagesUserLogoutDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStagesUserLogoutDeleteParamsWithTimeout(timeout time.Duration) *StagesUserLogoutDeleteParams {
	var ()
	return &StagesUserLogoutDeleteParams{

		timeout: timeout,
	}
}

// NewStagesUserLogoutDeleteParamsWithContext creates a new StagesUserLogoutDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewStagesUserLogoutDeleteParamsWithContext(ctx context.Context) *StagesUserLogoutDeleteParams {
	var ()
	return &StagesUserLogoutDeleteParams{

		Context: ctx,
	}
}

// NewStagesUserLogoutDeleteParamsWithHTTPClient creates a new StagesUserLogoutDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStagesUserLogoutDeleteParamsWithHTTPClient(client *http.Client) *StagesUserLogoutDeleteParams {
	var ()
	return &StagesUserLogoutDeleteParams{
		HTTPClient: client,
	}
}

/*StagesUserLogoutDeleteParams contains all the parameters to send to the API endpoint
for the stages user logout delete operation typically these are written to a http.Request
*/
type StagesUserLogoutDeleteParams struct {

	/*StageUUID
	  A UUID string identifying this User Logout Stage.

	*/
	StageUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stages user logout delete params
func (o *StagesUserLogoutDeleteParams) WithTimeout(timeout time.Duration) *StagesUserLogoutDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stages user logout delete params
func (o *StagesUserLogoutDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stages user logout delete params
func (o *StagesUserLogoutDeleteParams) WithContext(ctx context.Context) *StagesUserLogoutDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stages user logout delete params
func (o *StagesUserLogoutDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stages user logout delete params
func (o *StagesUserLogoutDeleteParams) WithHTTPClient(client *http.Client) *StagesUserLogoutDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stages user logout delete params
func (o *StagesUserLogoutDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStageUUID adds the stageUUID to the stages user logout delete params
func (o *StagesUserLogoutDeleteParams) WithStageUUID(stageUUID strfmt.UUID) *StagesUserLogoutDeleteParams {
	o.SetStageUUID(stageUUID)
	return o
}

// SetStageUUID adds the stageUuid to the stages user logout delete params
func (o *StagesUserLogoutDeleteParams) SetStageUUID(stageUUID strfmt.UUID) {
	o.StageUUID = stageUUID
}

// WriteToRequest writes these params to a swagger request
func (o *StagesUserLogoutDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
