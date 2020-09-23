// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetOracleDenomsDenomExchangeRateReader is a Reader for the GetOracleDenomsDenomExchangeRate structure.
type GetOracleDenomsDenomExchangeRateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOracleDenomsDenomExchangeRateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOracleDenomsDenomExchangeRateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOracleDenomsDenomExchangeRateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOracleDenomsDenomExchangeRateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOracleDenomsDenomExchangeRateOK creates a GetOracleDenomsDenomExchangeRateOK with default headers values
func NewGetOracleDenomsDenomExchangeRateOK() *GetOracleDenomsDenomExchangeRateOK {
	return &GetOracleDenomsDenomExchangeRateOK{}
}

/*GetOracleDenomsDenomExchangeRateOK handles this case with default header values.

current exchange rate of denom i.e. "1000.0"
*/
type GetOracleDenomsDenomExchangeRateOK struct {
	Payload *GetOracleDenomsDenomExchangeRateOKBody
}

func (o *GetOracleDenomsDenomExchangeRateOK) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/{denom}/exchange_rate][%d] getOracleDenomsDenomExchangeRateOK  %+v", 200, o.Payload)
}

func (o *GetOracleDenomsDenomExchangeRateOK) GetPayload() *GetOracleDenomsDenomExchangeRateOKBody {
	return o.Payload
}

func (o *GetOracleDenomsDenomExchangeRateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOracleDenomsDenomExchangeRateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOracleDenomsDenomExchangeRateBadRequest creates a GetOracleDenomsDenomExchangeRateBadRequest with default headers values
func NewGetOracleDenomsDenomExchangeRateBadRequest() *GetOracleDenomsDenomExchangeRateBadRequest {
	return &GetOracleDenomsDenomExchangeRateBadRequest{}
}

/*GetOracleDenomsDenomExchangeRateBadRequest handles this case with default header values.

Bad Request
*/
type GetOracleDenomsDenomExchangeRateBadRequest struct {
}

func (o *GetOracleDenomsDenomExchangeRateBadRequest) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/{denom}/exchange_rate][%d] getOracleDenomsDenomExchangeRateBadRequest ", 400)
}

func (o *GetOracleDenomsDenomExchangeRateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOracleDenomsDenomExchangeRateInternalServerError creates a GetOracleDenomsDenomExchangeRateInternalServerError with default headers values
func NewGetOracleDenomsDenomExchangeRateInternalServerError() *GetOracleDenomsDenomExchangeRateInternalServerError {
	return &GetOracleDenomsDenomExchangeRateInternalServerError{}
}

/*GetOracleDenomsDenomExchangeRateInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetOracleDenomsDenomExchangeRateInternalServerError struct {
}

func (o *GetOracleDenomsDenomExchangeRateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /oracle/denoms/{denom}/exchange_rate][%d] getOracleDenomsDenomExchangeRateInternalServerError ", 500)
}

func (o *GetOracleDenomsDenomExchangeRateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetOracleDenomsDenomExchangeRateOKBody get oracle denoms denom exchange rate o k body
swagger:model GetOracleDenomsDenomExchangeRateOKBody
*/
type GetOracleDenomsDenomExchangeRateOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result string `json:"result,omitempty"`
}

// Validate validates this get oracle denoms denom exchange rate o k body
func (o *GetOracleDenomsDenomExchangeRateOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOracleDenomsDenomExchangeRateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOracleDenomsDenomExchangeRateOKBody) UnmarshalBinary(b []byte) error {
	var res GetOracleDenomsDenomExchangeRateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
