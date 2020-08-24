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

// StagesUserLogoutReadReader is a Reader for the StagesUserLogoutRead structure.
type StagesUserLogoutReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StagesUserLogoutReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStagesUserLogoutReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStagesUserLogoutReadOK creates a StagesUserLogoutReadOK with default headers values
func NewStagesUserLogoutReadOK() *StagesUserLogoutReadOK {
	return &StagesUserLogoutReadOK{}
}

/*StagesUserLogoutReadOK handles this case with default header values.

StagesUserLogoutReadOK stages user logout read o k
*/
type StagesUserLogoutReadOK struct {
	Payload *models.UserLogoutStage
}

func (o *StagesUserLogoutReadOK) Error() string {
	return fmt.Sprintf("[GET /stages/user_logout/{stage_uuid}/][%d] stagesUserLogoutReadOK  %+v", 200, o.Payload)
}

func (o *StagesUserLogoutReadOK) GetPayload() *models.UserLogoutStage {
	return o.Payload
}

func (o *StagesUserLogoutReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserLogoutStage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
