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

// NewPoliciesHaveibeenpwnedDeleteParams creates a new PoliciesHaveibeenpwnedDeleteParams object
// with the default values initialized.
func NewPoliciesHaveibeenpwnedDeleteParams() *PoliciesHaveibeenpwnedDeleteParams {
	var ()
	return &PoliciesHaveibeenpwnedDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPoliciesHaveibeenpwnedDeleteParamsWithTimeout creates a new PoliciesHaveibeenpwnedDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPoliciesHaveibeenpwnedDeleteParamsWithTimeout(timeout time.Duration) *PoliciesHaveibeenpwnedDeleteParams {
	var ()
	return &PoliciesHaveibeenpwnedDeleteParams{

		timeout: timeout,
	}
}

// NewPoliciesHaveibeenpwnedDeleteParamsWithContext creates a new PoliciesHaveibeenpwnedDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPoliciesHaveibeenpwnedDeleteParamsWithContext(ctx context.Context) *PoliciesHaveibeenpwnedDeleteParams {
	var ()
	return &PoliciesHaveibeenpwnedDeleteParams{

		Context: ctx,
	}
}

// NewPoliciesHaveibeenpwnedDeleteParamsWithHTTPClient creates a new PoliciesHaveibeenpwnedDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPoliciesHaveibeenpwnedDeleteParamsWithHTTPClient(client *http.Client) *PoliciesHaveibeenpwnedDeleteParams {
	var ()
	return &PoliciesHaveibeenpwnedDeleteParams{
		HTTPClient: client,
	}
}

/*PoliciesHaveibeenpwnedDeleteParams contains all the parameters to send to the API endpoint
for the policies haveibeenpwned delete operation typically these are written to a http.Request
*/
type PoliciesHaveibeenpwnedDeleteParams struct {

	/*PolicyUUID
	  A UUID string identifying this Have I Been Pwned Policy.

	*/
	PolicyUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the policies haveibeenpwned delete params
func (o *PoliciesHaveibeenpwnedDeleteParams) WithTimeout(timeout time.Duration) *PoliciesHaveibeenpwnedDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the policies haveibeenpwned delete params
func (o *PoliciesHaveibeenpwnedDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the policies haveibeenpwned delete params
func (o *PoliciesHaveibeenpwnedDeleteParams) WithContext(ctx context.Context) *PoliciesHaveibeenpwnedDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the policies haveibeenpwned delete params
func (o *PoliciesHaveibeenpwnedDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the policies haveibeenpwned delete params
func (o *PoliciesHaveibeenpwnedDeleteParams) WithHTTPClient(client *http.Client) *PoliciesHaveibeenpwnedDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the policies haveibeenpwned delete params
func (o *PoliciesHaveibeenpwnedDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicyUUID adds the policyUUID to the policies haveibeenpwned delete params
func (o *PoliciesHaveibeenpwnedDeleteParams) WithPolicyUUID(policyUUID strfmt.UUID) *PoliciesHaveibeenpwnedDeleteParams {
	o.SetPolicyUUID(policyUUID)
	return o
}

// SetPolicyUUID adds the policyUuid to the policies haveibeenpwned delete params
func (o *PoliciesHaveibeenpwnedDeleteParams) SetPolicyUUID(policyUUID strfmt.UUID) {
	o.PolicyUUID = policyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PoliciesHaveibeenpwnedDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
