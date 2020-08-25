// Code generated by go-swagger; DO NOT EDIT.

package outposts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/BeryJu/passbook/proxy/pkg/models"
)

// OutpostsOutpostsCreateReader is a Reader for the OutpostsOutpostsCreate structure.
type OutpostsOutpostsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OutpostsOutpostsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOutpostsOutpostsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOutpostsOutpostsCreateCreated creates a OutpostsOutpostsCreateCreated with default headers values
func NewOutpostsOutpostsCreateCreated() *OutpostsOutpostsCreateCreated {
	return &OutpostsOutpostsCreateCreated{}
}

/*OutpostsOutpostsCreateCreated handles this case with default header values.

OutpostsOutpostsCreateCreated outposts outposts create created
*/
type OutpostsOutpostsCreateCreated struct {
	Payload *models.Outpost
}

func (o *OutpostsOutpostsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /outposts/outposts/][%d] outpostsOutpostsCreateCreated  %+v", 201, o.Payload)
}

func (o *OutpostsOutpostsCreateCreated) GetPayload() *models.Outpost {
	return o.Payload
}

func (o *OutpostsOutpostsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Outpost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
