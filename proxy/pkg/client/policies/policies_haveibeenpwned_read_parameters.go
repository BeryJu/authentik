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
)

// NewPoliciesHaveibeenpwnedReadParams creates a new PoliciesHaveibeenpwnedReadParams object
// with the default values initialized.
func NewPoliciesHaveibeenpwnedReadParams() *PoliciesHaveibeenpwnedReadParams {
	var ()
	return &PoliciesHaveibeenpwnedReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPoliciesHaveibeenpwnedReadParamsWithTimeout creates a new PoliciesHaveibeenpwnedReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPoliciesHaveibeenpwnedReadParamsWithTimeout(timeout time.Duration) *PoliciesHaveibeenpwnedReadParams {
	var ()
	return &PoliciesHaveibeenpwnedReadParams{

		timeout: timeout,
	}
}

// NewPoliciesHaveibeenpwnedReadParamsWithContext creates a new PoliciesHaveibeenpwnedReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewPoliciesHaveibeenpwnedReadParamsWithContext(ctx context.Context) *PoliciesHaveibeenpwnedReadParams {
	var ()
	return &PoliciesHaveibeenpwnedReadParams{

		Context: ctx,
	}
}

// NewPoliciesHaveibeenpwnedReadParamsWithHTTPClient creates a new PoliciesHaveibeenpwnedReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPoliciesHaveibeenpwnedReadParamsWithHTTPClient(client *http.Client) *PoliciesHaveibeenpwnedReadParams {
	var ()
	return &PoliciesHaveibeenpwnedReadParams{
		HTTPClient: client,
	}
}

/*PoliciesHaveibeenpwnedReadParams contains all the parameters to send to the API endpoint
for the policies haveibeenpwned read operation typically these are written to a http.Request
*/
type PoliciesHaveibeenpwnedReadParams struct {

	/*PolicyUUID
	  A UUID string identifying this Have I Been Pwned Policy.

	*/
	PolicyUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the policies haveibeenpwned read params
func (o *PoliciesHaveibeenpwnedReadParams) WithTimeout(timeout time.Duration) *PoliciesHaveibeenpwnedReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the policies haveibeenpwned read params
func (o *PoliciesHaveibeenpwnedReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the policies haveibeenpwned read params
func (o *PoliciesHaveibeenpwnedReadParams) WithContext(ctx context.Context) *PoliciesHaveibeenpwnedReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the policies haveibeenpwned read params
func (o *PoliciesHaveibeenpwnedReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the policies haveibeenpwned read params
func (o *PoliciesHaveibeenpwnedReadParams) WithHTTPClient(client *http.Client) *PoliciesHaveibeenpwnedReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the policies haveibeenpwned read params
func (o *PoliciesHaveibeenpwnedReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicyUUID adds the policyUUID to the policies haveibeenpwned read params
func (o *PoliciesHaveibeenpwnedReadParams) WithPolicyUUID(policyUUID strfmt.UUID) *PoliciesHaveibeenpwnedReadParams {
	o.SetPolicyUUID(policyUUID)
	return o
}

// SetPolicyUUID adds the policyUuid to the policies haveibeenpwned read params
func (o *PoliciesHaveibeenpwnedReadParams) SetPolicyUUID(policyUUID strfmt.UUID) {
	o.PolicyUUID = policyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PoliciesHaveibeenpwnedReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param policy_uuid
	if err := r.SetPathParam("policy_uuid", o.PolicyUUID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
