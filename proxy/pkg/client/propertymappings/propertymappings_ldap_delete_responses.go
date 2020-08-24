// Code generated by go-swagger; DO NOT EDIT.

package propertymappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PropertymappingsLdapDeleteReader is a Reader for the PropertymappingsLdapDelete structure.
type PropertymappingsLdapDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PropertymappingsLdapDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPropertymappingsLdapDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPropertymappingsLdapDeleteNoContent creates a PropertymappingsLdapDeleteNoContent with default headers values
func NewPropertymappingsLdapDeleteNoContent() *PropertymappingsLdapDeleteNoContent {
	return &PropertymappingsLdapDeleteNoContent{}
}

/*PropertymappingsLdapDeleteNoContent handles this case with default header values.

PropertymappingsLdapDeleteNoContent propertymappings ldap delete no content
*/
type PropertymappingsLdapDeleteNoContent struct {
}

func (o *PropertymappingsLdapDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /propertymappings/ldap/{pm_uuid}/][%d] propertymappingsLdapDeleteNoContent ", 204)
}

func (o *PropertymappingsLdapDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
