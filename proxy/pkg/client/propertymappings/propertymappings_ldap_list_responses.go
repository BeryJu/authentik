// Code generated by go-swagger; DO NOT EDIT.

package propertymappings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/BeryJu/passbook/proxy/pkg/models"
)

// PropertymappingsLdapListReader is a Reader for the PropertymappingsLdapList structure.
type PropertymappingsLdapListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PropertymappingsLdapListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPropertymappingsLdapListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPropertymappingsLdapListOK creates a PropertymappingsLdapListOK with default headers values
func NewPropertymappingsLdapListOK() *PropertymappingsLdapListOK {
	return &PropertymappingsLdapListOK{}
}

/*PropertymappingsLdapListOK handles this case with default header values.

PropertymappingsLdapListOK propertymappings ldap list o k
*/
type PropertymappingsLdapListOK struct {
	Payload *PropertymappingsLdapListOKBody
}

func (o *PropertymappingsLdapListOK) Error() string {
	return fmt.Sprintf("[GET /propertymappings/ldap/][%d] propertymappingsLdapListOK  %+v", 200, o.Payload)
}

func (o *PropertymappingsLdapListOK) GetPayload() *PropertymappingsLdapListOKBody {
	return o.Payload
}

func (o *PropertymappingsLdapListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PropertymappingsLdapListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PropertymappingsLdapListOKBody propertymappings ldap list o k body
swagger:model PropertymappingsLdapListOKBody
*/
type PropertymappingsLdapListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.LDAPPropertyMapping `json:"results"`
}

// Validate validates this propertymappings ldap list o k body
func (o *PropertymappingsLdapListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PropertymappingsLdapListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("propertymappingsLdapListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *PropertymappingsLdapListOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("propertymappingsLdapListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PropertymappingsLdapListOKBody) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("propertymappingsLdapListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PropertymappingsLdapListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("propertymappingsLdapListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("propertymappingsLdapListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PropertymappingsLdapListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PropertymappingsLdapListOKBody) UnmarshalBinary(b []byte) error {
	var res PropertymappingsLdapListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
