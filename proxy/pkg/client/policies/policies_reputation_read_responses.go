// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/BeryJu/passbook/proxy/pkg/models"
)

// PoliciesReputationReadReader is a Reader for the PoliciesReputationRead structure.
type PoliciesReputationReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoliciesReputationReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPoliciesReputationReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPoliciesReputationReadOK creates a PoliciesReputationReadOK with default headers values
func NewPoliciesReputationReadOK() *PoliciesReputationReadOK {
	return &PoliciesReputationReadOK{}
}

/*PoliciesReputationReadOK handles this case with default header values.

PoliciesReputationReadOK policies reputation read o k
*/
type PoliciesReputationReadOK struct {
	Payload *models.ReputationPolicy
}

func (o *PoliciesReputationReadOK) Error() string {
	return fmt.Sprintf("[GET /policies/reputation/{policy_uuid}/][%d] policiesReputationReadOK  %+v", 200, o.Payload)
}

func (o *PoliciesReputationReadOK) GetPayload() *models.ReputationPolicy {
	return o.Payload
}

func (o *PoliciesReputationReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReputationPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
