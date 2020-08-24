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

// StagesUserWritePartialUpdateReader is a Reader for the StagesUserWritePartialUpdate structure.
type StagesUserWritePartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StagesUserWritePartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStagesUserWritePartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStagesUserWritePartialUpdateOK creates a StagesUserWritePartialUpdateOK with default headers values
func NewStagesUserWritePartialUpdateOK() *StagesUserWritePartialUpdateOK {
	return &StagesUserWritePartialUpdateOK{}
}

/*StagesUserWritePartialUpdateOK handles this case with default header values.

StagesUserWritePartialUpdateOK stages user write partial update o k
*/
type StagesUserWritePartialUpdateOK struct {
	Payload *models.UserWriteStage
}

func (o *StagesUserWritePartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /stages/user_write/{stage_uuid}/][%d] stagesUserWritePartialUpdateOK  %+v", 200, o.Payload)
}

func (o *StagesUserWritePartialUpdateOK) GetPayload() *models.UserWriteStage {
	return o.Payload
}

func (o *StagesUserWritePartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserWriteStage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
