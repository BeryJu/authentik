// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewPoliciesHaveibeenpwnedUpdateParams creates a new PoliciesHaveibeenpwnedUpdateParams object
// with the default values initialized.
func NewPoliciesHaveibeenpwnedUpdateParams() *PoliciesHaveibeenpwnedUpdateParams {
	var ()
	return &PoliciesHaveibeenpwnedUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPoliciesHaveibeenpwnedUpdateParamsWithTimeout creates a new PoliciesHaveibeenpwnedUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPoliciesHaveibeenpwnedUpdateParamsWithTimeout(timeout time.Duration) *PoliciesHaveibeenpwnedUpdateParams {
	var ()
	return &PoliciesHaveibeenpwnedUpdateParams{

		timeout: timeout,
	}
}

// NewPoliciesHaveibeenpwnedUpdateParamsWithContext creates a new PoliciesHaveibeenpwnedUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPoliciesHaveibeenpwnedUpdateParamsWithContext(ctx context.Context) *PoliciesHaveibeenpwnedUpdateParams {
	var ()
	return &PoliciesHaveibeenpwnedUpdateParams{

		Context: ctx,
	}
}

// NewPoliciesHaveibeenpwnedUpdateParamsWithHTTPClient creates a new PoliciesHaveibeenpwnedUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPoliciesHaveibeenpwnedUpdateParamsWithHTTPClient(client *http.Client) *PoliciesHaveibeenpwnedUpdateParams {
	var ()
	return &PoliciesHaveibeenpwnedUpdateParams{
		HTTPClient: client,
	}
}

/*PoliciesHaveibeenpwnedUpdateParams contains all the parameters to send to the API endpoint
for the policies haveibeenpwned update operation typically these are written to a http.Request
*/
type PoliciesHaveibeenpwnedUpdateParams struct {

	/*Data*/
	Data *models.HaveIBeenPwendPolicy
	/*PolicyUUID
	  A UUID string identifying this Have I Been Pwned Policy.

	*/
	PolicyUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) WithTimeout(timeout time.Duration) *PoliciesHaveibeenpwnedUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) WithContext(ctx context.Context) *PoliciesHaveibeenpwnedUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) WithHTTPClient(client *http.Client) *PoliciesHaveibeenpwnedUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) WithData(data *models.HaveIBeenPwendPolicy) *PoliciesHaveibeenpwnedUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) SetData(data *models.HaveIBeenPwendPolicy) {
	o.Data = data
}

// WithPolicyUUID adds the policyUUID to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) WithPolicyUUID(policyUUID strfmt.UUID) *PoliciesHaveibeenpwnedUpdateParams {
	o.SetPolicyUUID(policyUUID)
	return o
}

// SetPolicyUUID adds the policyUuid to the policies haveibeenpwned update params
func (o *PoliciesHaveibeenpwnedUpdateParams) SetPolicyUUID(policyUUID strfmt.UUID) {
	o.PolicyUUID = policyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PoliciesHaveibeenpwnedUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param policy_uuid
	if err := r.SetPathParam("policy_uuid", o.PolicyUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
