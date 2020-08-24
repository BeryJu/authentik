// Code generated by go-swagger; DO NOT EDIT.

package sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/BeryJu/passbook/proxy/pkg/models"
)

// SourcesSamlUpdateReader is a Reader for the SourcesSamlUpdate structure.
type SourcesSamlUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SourcesSamlUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSourcesSamlUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSourcesSamlUpdateOK creates a SourcesSamlUpdateOK with default headers values
func NewSourcesSamlUpdateOK() *SourcesSamlUpdateOK {
	return &SourcesSamlUpdateOK{}
}

/*SourcesSamlUpdateOK handles this case with default header values.

SourcesSamlUpdateOK sources saml update o k
*/
type SourcesSamlUpdateOK struct {
	Payload *models.SAMLSource
}

func (o *SourcesSamlUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /sources/saml/{pbm_uuid}/][%d] sourcesSamlUpdateOK  %+v", 200, o.Payload)
}

func (o *SourcesSamlUpdateOK) GetPayload() *models.SAMLSource {
	return o.Payload
}

func (o *SourcesSamlUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SAMLSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
