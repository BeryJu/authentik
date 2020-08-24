// Code generated by go-swagger; DO NOT EDIT.

package stages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StagesUserLoginDeleteReader is a Reader for the StagesUserLoginDelete structure.
type StagesUserLoginDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StagesUserLoginDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewStagesUserLoginDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStagesUserLoginDeleteNoContent creates a StagesUserLoginDeleteNoContent with default headers values
func NewStagesUserLoginDeleteNoContent() *StagesUserLoginDeleteNoContent {
	return &StagesUserLoginDeleteNoContent{}
}

/*StagesUserLoginDeleteNoContent handles this case with default header values.

StagesUserLoginDeleteNoContent stages user login delete no content
*/
type StagesUserLoginDeleteNoContent struct {
}

func (o *StagesUserLoginDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /stages/user_login/{stage_uuid}/][%d] stagesUserLoginDeleteNoContent ", 204)
}

func (o *StagesUserLoginDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
