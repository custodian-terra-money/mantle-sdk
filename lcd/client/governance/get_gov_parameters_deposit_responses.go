// Code generated by go-swagger; DO NOT EDIT.

package governance

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

	"github.com/terra-money/mantle-sdk/lcd/models"
)

// GetGovParametersDepositReader is a Reader for the GetGovParametersDeposit structure.
type GetGovParametersDepositReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGovParametersDepositReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGovParametersDepositOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGovParametersDepositBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGovParametersDepositNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGovParametersDepositInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGovParametersDepositOK creates a GetGovParametersDepositOK with default headers values
func NewGetGovParametersDepositOK() *GetGovParametersDepositOK {
	return &GetGovParametersDepositOK{}
}

/*GetGovParametersDepositOK handles this case with default header values.

OK
*/
type GetGovParametersDepositOK struct {
	Payload *GetGovParametersDepositOKBody
}

func (o *GetGovParametersDepositOK) Error() string {
	return fmt.Sprintf("[GET /gov/parameters/deposit][%d] getGovParametersDepositOK  %+v", 200, o.Payload)
}

func (o *GetGovParametersDepositOK) GetPayload() *GetGovParametersDepositOKBody {
	return o.Payload
}

func (o *GetGovParametersDepositOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetGovParametersDepositOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGovParametersDepositBadRequest creates a GetGovParametersDepositBadRequest with default headers values
func NewGetGovParametersDepositBadRequest() *GetGovParametersDepositBadRequest {
	return &GetGovParametersDepositBadRequest{}
}

/*GetGovParametersDepositBadRequest handles this case with default header values.

<other_path> is not a valid query request path
*/
type GetGovParametersDepositBadRequest struct {
}

func (o *GetGovParametersDepositBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/parameters/deposit][%d] getGovParametersDepositBadRequest ", 400)
}

func (o *GetGovParametersDepositBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovParametersDepositNotFound creates a GetGovParametersDepositNotFound with default headers values
func NewGetGovParametersDepositNotFound() *GetGovParametersDepositNotFound {
	return &GetGovParametersDepositNotFound{}
}

/*GetGovParametersDepositNotFound handles this case with default header values.

Found no deposit parameters
*/
type GetGovParametersDepositNotFound struct {
}

func (o *GetGovParametersDepositNotFound) Error() string {
	return fmt.Sprintf("[GET /gov/parameters/deposit][%d] getGovParametersDepositNotFound ", 404)
}

func (o *GetGovParametersDepositNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovParametersDepositInternalServerError creates a GetGovParametersDepositInternalServerError with default headers values
func NewGetGovParametersDepositInternalServerError() *GetGovParametersDepositInternalServerError {
	return &GetGovParametersDepositInternalServerError{}
}

/*GetGovParametersDepositInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetGovParametersDepositInternalServerError struct {
}

func (o *GetGovParametersDepositInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/parameters/deposit][%d] getGovParametersDepositInternalServerError ", 500)
}

func (o *GetGovParametersDepositInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetGovParametersDepositOKBody get gov parameters deposit o k body
swagger:model GetGovParametersDepositOKBody
*/
type GetGovParametersDepositOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *GetGovParametersDepositOKBodyResult `json:"result,omitempty"`
}

// Validate validates this get gov parameters deposit o k body
func (o *GetGovParametersDepositOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGovParametersDepositOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getGovParametersDepositOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGovParametersDepositOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovParametersDepositOKBody) UnmarshalBinary(b []byte) error {
	var res GetGovParametersDepositOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetGovParametersDepositOKBodyResult get gov parameters deposit o k body result
swagger:model GetGovParametersDepositOKBodyResult
*/
type GetGovParametersDepositOKBodyResult struct {

	// max deposit period
	MaxDepositPeriod string `json:"max_deposit_period,omitempty"`

	// min deposit
	MinDeposit []*models.Coin `json:"min_deposit"`
}

// Validate validates this get gov parameters deposit o k body result
func (o *GetGovParametersDepositOKBodyResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMinDeposit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGovParametersDepositOKBodyResult) validateMinDeposit(formats strfmt.Registry) error {

	if swag.IsZero(o.MinDeposit) { // not required
		return nil
	}

	for i := 0; i < len(o.MinDeposit); i++ {
		if swag.IsZero(o.MinDeposit[i]) { // not required
			continue
		}

		if o.MinDeposit[i] != nil {
			if err := o.MinDeposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getGovParametersDepositOK" + "." + "result" + "." + "min_deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGovParametersDepositOKBodyResult) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovParametersDepositOKBodyResult) UnmarshalBinary(b []byte) error {
	var res GetGovParametersDepositOKBodyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
