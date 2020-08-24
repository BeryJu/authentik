// Code generated by go-swagger; DO NOT EDIT.

package sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SourcesSamlDeleteReader is a Reader for the SourcesSamlDelete structure.
type SourcesSamlDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SourcesSamlDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSourcesSamlDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSourcesSamlDeleteNoContent creates a SourcesSamlDeleteNoContent with default headers values
func NewSourcesSamlDeleteNoContent() *SourcesSamlDeleteNoContent {
	return &SourcesSamlDeleteNoContent{}
}

/*SourcesSamlDeleteNoContent handles this case with default header values.

SourcesSamlDeleteNoContent sources saml delete no content
*/
type SourcesSamlDeleteNoContent struct {
}

func (o *SourcesSamlDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /sources/saml/{pbm_uuid}/][%d] sourcesSamlDeleteNoContent ", 204)
}

func (o *SourcesSamlDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
