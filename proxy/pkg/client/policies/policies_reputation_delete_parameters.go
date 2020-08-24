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

// NewPoliciesReputationDeleteParams creates a new PoliciesReputationDeleteParams object
// with the default values initialized.
func NewPoliciesReputationDeleteParams() *PoliciesReputationDeleteParams {
	var ()
	return &PoliciesReputationDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPoliciesReputationDeleteParamsWithTimeout creates a new PoliciesReputationDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPoliciesReputationDeleteParamsWithTimeout(timeout time.Duration) *PoliciesReputationDeleteParams {
	var ()
	return &PoliciesReputationDeleteParams{

		timeout: timeout,
	}
}

// NewPoliciesReputationDeleteParamsWithContext creates a new PoliciesReputationDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPoliciesReputationDeleteParamsWithContext(ctx context.Context) *PoliciesReputationDeleteParams {
	var ()
	return &PoliciesReputationDeleteParams{

		Context: ctx,
	}
}

// NewPoliciesReputationDeleteParamsWithHTTPClient creates a new PoliciesReputationDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPoliciesReputationDeleteParamsWithHTTPClient(client *http.Client) *PoliciesReputationDeleteParams {
	var ()
	return &PoliciesReputationDeleteParams{
		HTTPClient: client,
	}
}

/*PoliciesReputationDeleteParams contains all the parameters to send to the API endpoint
for the policies reputation delete operation typically these are written to a http.Request
*/
type PoliciesReputationDeleteParams struct {

	/*PolicyUUID
	  A UUID string identifying this Reputation Policy.

	*/
	PolicyUUID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the policies reputation delete params
func (o *PoliciesReputationDeleteParams) WithTimeout(timeout time.Duration) *PoliciesReputationDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the policies reputation delete params
func (o *PoliciesReputationDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the policies reputation delete params
func (o *PoliciesReputationDeleteParams) WithContext(ctx context.Context) *PoliciesReputationDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the policies reputation delete params
func (o *PoliciesReputationDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the policies reputation delete params
func (o *PoliciesReputationDeleteParams) WithHTTPClient(client *http.Client) *PoliciesReputationDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the policies reputation delete params
func (o *PoliciesReputationDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicyUUID adds the policyUUID to the policies reputation delete params
func (o *PoliciesReputationDeleteParams) WithPolicyUUID(policyUUID strfmt.UUID) *PoliciesReputationDeleteParams {
	o.SetPolicyUUID(policyUUID)
	return o
}

// SetPolicyUUID adds the policyUuid to the policies reputation delete params
func (o *PoliciesReputationDeleteParams) SetPolicyUUID(policyUUID strfmt.UUID) {
	o.PolicyUUID = policyUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PoliciesReputationDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
