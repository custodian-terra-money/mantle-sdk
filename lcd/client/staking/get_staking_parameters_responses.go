// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetStakingParametersReader is a Reader for the GetStakingParameters structure.
type GetStakingParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStakingParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStakingParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetStakingParametersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStakingParametersOK creates a GetStakingParametersOK with default headers values
func NewGetStakingParametersOK() *GetStakingParametersOK {
	return &GetStakingParametersOK{}
}

/*GetStakingParametersOK handles this case with default header values.

OK
*/
type GetStakingParametersOK struct {
	Payload *GetStakingParametersOKBody
}

func (o *GetStakingParametersOK) Error() string {
	return fmt.Sprintf("[GET /staking/parameters][%d] getStakingParametersOK  %+v", 200, o.Payload)
}

func (o *GetStakingParametersOK) GetPayload() *GetStakingParametersOKBody {
	return o.Payload
}

func (o *GetStakingParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStakingParametersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStakingParametersInternalServerError creates a GetStakingParametersInternalServerError with default headers values
func NewGetStakingParametersInternalServerError() *GetStakingParametersInternalServerError {
	return &GetStakingParametersInternalServerError{}
}

/*GetStakingParametersInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetStakingParametersInternalServerError struct {
}

func (o *GetStakingParametersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/parameters][%d] getStakingParametersInternalServerError ", 500)
}

func (o *GetStakingParametersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetStakingParametersOKBody get staking parameters o k body
swagger:model GetStakingParametersOKBody
*/
type GetStakingParametersOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *GetStakingParametersOKBodyResult `json:"result,omitempty"`
}

// Validate validates this get staking parameters o k body
func (o *GetStakingParametersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStakingParametersOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getStakingParametersOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingParametersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingParametersOKBody) UnmarshalBinary(b []byte) error {
	var res GetStakingParametersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetStakingParametersOKBodyResult get staking parameters o k body result
swagger:model GetStakingParametersOKBodyResult
*/
type GetStakingParametersOKBodyResult struct {

	// bond denom
	BondDenom string `json:"bond_denom,omitempty"`

	// max entries
	MaxEntries int64 `json:"max_entries,omitempty"`

	// max validators
	MaxValidators int64 `json:"max_validators,omitempty"`

	// unbonding time
	UnbondingTime string `json:"unbonding_time,omitempty"`
}

// Validate validates this get staking parameters o k body result
func (o *GetStakingParametersOKBodyResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingParametersOKBodyResult) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingParametersOKBodyResult) UnmarshalBinary(b []byte) error {
	var res GetStakingParametersOKBodyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
