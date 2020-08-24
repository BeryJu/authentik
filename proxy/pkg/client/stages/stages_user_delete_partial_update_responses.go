// Code generated by go-swagger; DO NOT EDIT.

package stages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/BeryJu/passbook/proxy/pkg/models"
)

// StagesUserDeletePartialUpdateReader is a Reader for the StagesUserDeletePartialUpdate structure.
type StagesUserDeletePartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StagesUserDeletePartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStagesUserDeletePartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStagesUserDeletePartialUpdateOK creates a StagesUserDeletePartialUpdateOK with default headers values
func NewStagesUserDeletePartialUpdateOK() *StagesUserDeletePartialUpdateOK {
	return &StagesUserDeletePartialUpdateOK{}
}

/*StagesUserDeletePartialUpdateOK handles this case with default header values.

StagesUserDeletePartialUpdateOK stages user delete partial update o k
*/
type StagesUserDeletePartialUpdateOK struct {
	Payload *models.UserDeleteStage
}

func (o *StagesUserDeletePartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /stages/user_delete/{stage_uuid}/][%d] stagesUserDeletePartialUpdateOK  %+v", 200, o.Payload)
}

func (o *StagesUserDeletePartialUpdateOK) GetPayload() *models.UserDeleteStage {
	return o.Payload
}

func (o *StagesUserDeletePartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserDeleteStage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
