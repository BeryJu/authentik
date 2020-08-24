// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PoliciesPasswordExpiryDeleteReader is a Reader for the PoliciesPasswordExpiryDelete structure.
type PoliciesPasswordExpiryDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoliciesPasswordExpiryDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPoliciesPasswordExpiryDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPoliciesPasswordExpiryDeleteNoContent creates a PoliciesPasswordExpiryDeleteNoContent with default headers values
func NewPoliciesPasswordExpiryDeleteNoContent() *PoliciesPasswordExpiryDeleteNoContent {
	return &PoliciesPasswordExpiryDeleteNoContent{}
}

/*PoliciesPasswordExpiryDeleteNoContent handles this case with default header values.

PoliciesPasswordExpiryDeleteNoContent policies password expiry delete no content
*/
type PoliciesPasswordExpiryDeleteNoContent struct {
}

func (o *PoliciesPasswordExpiryDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /policies/password_expiry/{policy_uuid}/][%d] policiesPasswordExpiryDeleteNoContent ", 204)
}

func (o *PoliciesPasswordExpiryDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
