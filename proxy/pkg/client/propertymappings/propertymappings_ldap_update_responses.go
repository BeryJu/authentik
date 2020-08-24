// Code generated by go-swagger; DO NOT EDIT.

package propertymappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/BeryJu/passbook/proxy/pkg/models"
)

// PropertymappingsLdapUpdateReader is a Reader for the PropertymappingsLdapUpdate structure.
type PropertymappingsLdapUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PropertymappingsLdapUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPropertymappingsLdapUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPropertymappingsLdapUpdateOK creates a PropertymappingsLdapUpdateOK with default headers values
func NewPropertymappingsLdapUpdateOK() *PropertymappingsLdapUpdateOK {
	return &PropertymappingsLdapUpdateOK{}
}

/*PropertymappingsLdapUpdateOK handles this case with default header values.

PropertymappingsLdapUpdateOK propertymappings ldap update o k
*/
type PropertymappingsLdapUpdateOK struct {
	Payload *models.LDAPPropertyMapping
}

func (o *PropertymappingsLdapUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /propertymappings/ldap/{pm_uuid}/][%d] propertymappingsLdapUpdateOK  %+v", 200, o.Payload)
}

func (o *PropertymappingsLdapUpdateOK) GetPayload() *models.LDAPPropertyMapping {
	return o.Payload
}

func (o *PropertymappingsLdapUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LDAPPropertyMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
