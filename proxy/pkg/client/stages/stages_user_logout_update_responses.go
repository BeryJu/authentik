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

// StagesUserLogoutUpdateReader is a Reader for the StagesUserLogoutUpdate structure.
type StagesUserLogoutUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StagesUserLogoutUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStagesUserLogoutUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStagesUserLogoutUpdateOK creates a StagesUserLogoutUpdateOK with default headers values
func NewStagesUserLogoutUpdateOK() *StagesUserLogoutUpdateOK {
	return &StagesUserLogoutUpdateOK{}
}

/*StagesUserLogoutUpdateOK handles this case with default header values.

StagesUserLogoutUpdateOK stages user logout update o k
*/
type StagesUserLogoutUpdateOK struct {
	Payload *models.UserLogoutStage
}

func (o *StagesUserLogoutUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /stages/user_logout/{stage_uuid}/][%d] stagesUserLogoutUpdateOK  %+v", 200, o.Payload)
}

func (o *StagesUserLogoutUpdateOK) GetPayload() *models.UserLogoutStage {
	return o.Payload
}

func (o *StagesUserLogoutUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserLogoutStage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
