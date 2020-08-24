// Code generated by go-swagger; DO NOT EDIT.

package core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/BeryJu/passbook/proxy/pkg/models"
)

// CoreUsersUpdateReader is a Reader for the CoreUsersUpdate structure.
type CoreUsersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoreUsersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCoreUsersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCoreUsersUpdateOK creates a CoreUsersUpdateOK with default headers values
func NewCoreUsersUpdateOK() *CoreUsersUpdateOK {
	return &CoreUsersUpdateOK{}
}

/*CoreUsersUpdateOK handles this case with default header values.

CoreUsersUpdateOK core users update o k
*/
type CoreUsersUpdateOK struct {
	Payload *models.User
}

func (o *CoreUsersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /core/users/{id}/][%d] coreUsersUpdateOK  %+v", 200, o.Payload)
}

func (o *CoreUsersUpdateOK) GetPayload() *models.User {
	return o.Payload
}

func (o *CoreUsersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
