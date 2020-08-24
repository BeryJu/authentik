// Code generated by go-swagger; DO NOT EDIT.

package stages

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

// StagesOtpTimeListReader is a Reader for the StagesOtpTimeList structure.
type StagesOtpTimeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StagesOtpTimeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStagesOtpTimeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStagesOtpTimeListOK creates a StagesOtpTimeListOK with default headers values
func NewStagesOtpTimeListOK() *StagesOtpTimeListOK {
	return &StagesOtpTimeListOK{}
}

/*StagesOtpTimeListOK handles this case with default header values.

StagesOtpTimeListOK stages otp time list o k
*/
type StagesOtpTimeListOK struct {
	Payload *StagesOtpTimeListOKBody
}

func (o *StagesOtpTimeListOK) Error() string {
	return fmt.Sprintf("[GET /stages/otp_time/][%d] stagesOtpTimeListOK  %+v", 200, o.Payload)
}

func (o *StagesOtpTimeListOK) GetPayload() *StagesOtpTimeListOKBody {
	return o.Payload
}

func (o *StagesOtpTimeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StagesOtpTimeListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StagesOtpTimeListOKBody stages otp time list o k body
swagger:model StagesOtpTimeListOKBody
*/
type StagesOtpTimeListOKBody struct {

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
	Results []*models.OTPTimeStage `json:"results"`
}

// Validate validates this stages otp time list o k body
func (o *StagesOtpTimeListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *StagesOtpTimeListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("stagesOtpTimeListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *StagesOtpTimeListOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("stagesOtpTimeListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *StagesOtpTimeListOKBody) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("stagesOtpTimeListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *StagesOtpTimeListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("stagesOtpTimeListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stagesOtpTimeListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StagesOtpTimeListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StagesOtpTimeListOKBody) UnmarshalBinary(b []byte) error {
	var res StagesOtpTimeListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
