// Code generated by go-swagger; DO NOT EDIT.

package providers

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

// ProvidersSamlListReader is a Reader for the ProvidersSamlList structure.
type ProvidersSamlListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProvidersSamlListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProvidersSamlListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProvidersSamlListOK creates a ProvidersSamlListOK with default headers values
func NewProvidersSamlListOK() *ProvidersSamlListOK {
	return &ProvidersSamlListOK{}
}

/*ProvidersSamlListOK handles this case with default header values.

ProvidersSamlListOK providers saml list o k
*/
type ProvidersSamlListOK struct {
	Payload *ProvidersSamlListOKBody
}

func (o *ProvidersSamlListOK) Error() string {
	return fmt.Sprintf("[GET /providers/saml/][%d] providersSamlListOK  %+v", 200, o.Payload)
}

func (o *ProvidersSamlListOK) GetPayload() *ProvidersSamlListOKBody {
	return o.Payload
}

func (o *ProvidersSamlListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ProvidersSamlListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProvidersSamlListOKBody providers saml list o k body
swagger:model ProvidersSamlListOKBody
*/
type ProvidersSamlListOKBody struct {

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
	Results []*models.SAMLProvider `json:"results"`
}

// Validate validates this providers saml list o k body
func (o *ProvidersSamlListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *ProvidersSamlListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("providersSamlListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *ProvidersSamlListOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("providersSamlListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ProvidersSamlListOKBody) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("providersSamlListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ProvidersSamlListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("providersSamlListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("providersSamlListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ProvidersSamlListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProvidersSamlListOKBody) UnmarshalBinary(b []byte) error {
	var res ProvidersSamlListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
