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

// PoliciesPasswordCreateReader is a Reader for the PoliciesPasswordCreate structure.
type PoliciesPasswordCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoliciesPasswordCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPoliciesPasswordCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPoliciesPasswordCreateCreated creates a PoliciesPasswordCreateCreated with default headers values
func NewPoliciesPasswordCreateCreated() *PoliciesPasswordCreateCreated {
	return &PoliciesPasswordCreateCreated{}
}

/*PoliciesPasswordCreateCreated handles this case with default header values.

PoliciesPasswordCreateCreated policies password create created
*/
type PoliciesPasswordCreateCreated struct {
	Payload *models.PasswordPolicy
}

func (o *PoliciesPasswordCreateCreated) Error() string {
	return fmt.Sprintf("[POST /policies/password/][%d] policiesPasswordCreateCreated  %+v", 201, o.Payload)
}

func (o *PoliciesPasswordCreateCreated) GetPayload() *models.PasswordPolicy {
	return o.Payload
}

func (o *PoliciesPasswordCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PasswordPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
