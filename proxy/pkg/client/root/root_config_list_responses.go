// Code generated by go-swagger; DO NOT EDIT.

package root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/BeryJu/passbook/proxy/pkg/models"
)

// RootConfigListReader is a Reader for the RootConfigList structure.
type RootConfigListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RootConfigListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRootConfigListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRootConfigListOK creates a RootConfigListOK with default headers values
func NewRootConfigListOK() *RootConfigListOK {
	return &RootConfigListOK{}
}

/*RootConfigListOK handles this case with default header values.

RootConfigListOK root config list o k
*/
type RootConfigListOK struct {
	Payload []*models.Config
}

func (o *RootConfigListOK) Error() string {
	return fmt.Sprintf("[GET /root/config/][%d] rootConfigListOK  %+v", 200, o.Payload)
}

func (o *RootConfigListOK) GetPayload() []*models.Config {
	return o.Payload
}

func (o *RootConfigListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
