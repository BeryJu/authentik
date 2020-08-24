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

	"github.com/BeryJu/passbook/proxy/pkg/models"
)

// NewStagesUserLogoutPartialUpdateParams creates a new StagesUserLogoutPartialUpdateParams object
// with the default values initialized.
func NewStagesUserLogoutPartialUpdateParams() *StagesUserLogoutPartialUpdateParams {
	var ()
	return &StagesUserLogoutPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStagesUserLogoutPartialUpdateParamsWithTimeout creates a new StagesUserLogoutPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStagesUserLogoutPartialUpdateParamsWithTimeout(timeout time.Duration) *StagesUserLogoutPartialUpdateParams {
	var ()
	return &StagesUserLogoutPartialUpdateParams{

		timeout: timeout,
	}
}

// NewStagesUserLogoutPartialUpdateParamsWithContext creates a new StagesUserLogoutPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewStagesUserLogoutPartialUpdateParamsWithContext(ctx context.Context) *StagesUserLogoutPartialUpdateParams {
	var ()
	return &StagesUserLogoutPartialUpdateParams{

		Context: ctx,
	}
}

// NewStagesUserLogoutPartialUpdateParamsWithHTTPClient creates a new StagesUserLogoutPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStagesUserLogoutPartialUpdateParamsWithHTTPClient(client *http.Client) *StagesUserLogoutPartialUpdateParams {
	var ()
	return &StagesUserLogoutPartialUpdateParams{
		HTTPClient: client,
	}
}

/*StagesUserLogoutPartialUpdateParams contains all the parameters to send to the API endpoint
for the stages user logout partial update operation typically these are written to a http.Request
*/
type StagesUserLogoutPartialUpdateParams struct {

	/*Data*/
	Data *models.UserLogoutStage
	/*StageUUID
	  A UUID string identifying this User Logout Stage.

	*/
	StageUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) WithTimeout(timeout time.Duration) *StagesUserLogoutPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) WithContext(ctx context.Context) *StagesUserLogoutPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) WithHTTPClient(client *http.Client) *StagesUserLogoutPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) WithData(data *models.UserLogoutStage) *StagesUserLogoutPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) SetData(data *models.UserLogoutStage) {
	o.Data = data
}

// WithStageUUID adds the stageUUID to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) WithStageUUID(stageUUID strfmt.UUID) *StagesUserLogoutPartialUpdateParams {
	o.SetStageUUID(stageUUID)
	return o
}

// SetStageUUID adds the stageUuid to the stages user logout partial update params
func (o *StagesUserLogoutPartialUpdateParams) SetStageUUID(stageUUID strfmt.UUID) {
	o.StageUUID = stageUUID
}

// WriteToRequest writes these params to a swagger request
func (o *StagesUserLogoutPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param stage_uuid
	if err := r.SetPathParam("stage_uuid", o.StageUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
