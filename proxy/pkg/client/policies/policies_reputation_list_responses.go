// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// PoliciesReputationListReader is a Reader for the PoliciesReputationList structure.
type PoliciesReputationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PoliciesReputationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPoliciesReputationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPoliciesReputationListOK creates a PoliciesReputationListOK with default headers values
func NewPoliciesReputationListOK() *PoliciesReputationListOK {
	return &PoliciesReputationListOK{}
}

/*PoliciesReputationListOK handles this case with default header values.

PoliciesReputationListOK policies reputation list o k
*/
type PoliciesReputationListOK struct {
	Payload *PoliciesReputationListOKBody
}

func (o *PoliciesReputationListOK) Error() string {
	return fmt.Sprintf("[GET /policies/reputation/][%d] policiesReputationListOK  %+v", 200, o.Payload)
}

func (o *PoliciesReputationListOK) GetPayload() *PoliciesReputationListOKBody {
	return o.Payload
}

func (o *PoliciesReputationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PoliciesReputationListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PoliciesReputationListOKBody policies reputation list o k body
swagger:model PoliciesReputationListOKBody
*/
type PoliciesReputationListOKBody struct {

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
	Results []*models.ReputationPolicy `json:"results"`
}

// Validate validates this policies reputation list o k body
func (o *PoliciesReputationListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PoliciesReputationListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("policiesReputationListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *PoliciesReputationListOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("policiesReputationListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PoliciesReputationListOKBody) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("policiesReputationListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *PoliciesReputationListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("policiesReputationListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policiesReputationListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PoliciesReputationListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PoliciesReputationListOKBody) UnmarshalBinary(b []byte) error {
	var res PoliciesReputationListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
