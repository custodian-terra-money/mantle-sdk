// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/terra-money/mantle-sdk/lcd/models"
)

// GetOracleVotersValidatorFeederReader is a Reader for the GetOracleVotersValidatorFeeder structure.
type GetOracleVotersValidatorFeederReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleVotersValidatorFeederReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOracleVotersValidatorFeederOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOracleVotersValidatorFeederBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOracleVotersValidatorFeederInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOracleVotersValidatorFeederOK creates a GetOracleVotersValidatorFeederOK with default headers values
func NewGetOracleVotersValidatorFeederOK() *GetOracleVotersValidatorFeederOK {
	return &GetOracleVotersValidatorFeederOK{}
}

/*GetOracleVotersValidatorFeederOK handles this case with default header values.

OK
*/
type GetOracleVotersValidatorFeederOK struct {
	Payload *GetOracleVotersValidatorFeederOKBody
}

func (o *GetOracleVotersValidatorFeederOK) Error() string {
	return fmt.Sprintf("[GET /oracle/voters/{validator}/feeder][%d] getOracleVotersValidatorFeederOK  %+v", 200, o.Payload)
}

func (o *GetOracleVotersValidatorFeederOK) GetPayload() *GetOracleVotersValidatorFeederOKBody {
	return o.Payload
}

func (o *GetOracleVotersValidatorFeederOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOracleVotersValidatorFeederOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleVotersValidatorFeederBadRequest creates a GetOracleVotersValidatorFeederBadRequest with default headers values
func NewGetOracleVotersValidatorFeederBadRequest() *GetOracleVotersValidatorFeederBadRequest {
	return &GetOracleVotersValidatorFeederBadRequest{}
}

/*GetOracleVotersValidatorFeederBadRequest handles this case with default header values.

Bad Request
*/
type GetOracleVotersValidatorFeederBadRequest struct {
}

func (o *GetOracleVotersValidatorFeederBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracle/voters/{validator}/feeder][%d] getOracleVotersValidatorFeederBadRequest ", 400)
}

func (o *GetOracleVotersValidatorFeederBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOracleVotersValidatorFeederInternalServerError creates a GetOracleVotersValidatorFeederInternalServerError with default headers values
func NewGetOracleVotersValidatorFeederInternalServerError() *GetOracleVotersValidatorFeederInternalServerError {
	return &GetOracleVotersValidatorFeederInternalServerError{}
}

/*GetOracleVotersValidatorFeederInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOracleVotersValidatorFeederInternalServerError struct {
}

func (o *GetOracleVotersValidatorFeederInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oracle/voters/{validator}/feeder][%d] getOracleVotersValidatorFeederInternalServerError ", 500)
}

func (o *GetOracleVotersValidatorFeederInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetOracleVotersValidatorFeederOKBody get oracle voters validator feeder o k body
swagger:model GetOracleVotersValidatorFeederOKBody
*/
type GetOracleVotersValidatorFeederOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result models.Address `json:"result,omitempty"`
}

// Validate validates this get oracle voters validator feeder o k body
func (o *GetOracleVotersValidatorFeederOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOracleVotersValidatorFeederOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if err := o.Result.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getOracleVotersValidatorFeederOK" + "." + "result")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOracleVotersValidatorFeederOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOracleVotersValidatorFeederOKBody) UnmarshalBinary(b []byte) error {
	var res GetOracleVotersValidatorFeederOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
