// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

	"github.com/terra-project/mantle/lcd/models"
)

// GetDistributionValidatorsValidatorAddrOutstandingRewardsReader is a Reader for the GetDistributionValidatorsValidatorAddrOutstandingRewards structure.
type GetDistributionValidatorsValidatorAddrOutstandingRewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistributionValidatorsValidatorAddrOutstandingRewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDistributionValidatorsValidatorAddrOutstandingRewardsOK creates a GetDistributionValidatorsValidatorAddrOutstandingRewardsOK with default headers values
func NewGetDistributionValidatorsValidatorAddrOutstandingRewardsOK() *GetDistributionValidatorsValidatorAddrOutstandingRewardsOK {
	return &GetDistributionValidatorsValidatorAddrOutstandingRewardsOK{}
}

/*GetDistributionValidatorsValidatorAddrOutstandingRewardsOK handles this case with default header values.

OK
*/
type GetDistributionValidatorsValidatorAddrOutstandingRewardsOK struct {
	Payload *GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody
}

func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsOK) Error() string {
	return fmt.Sprintf("[GET /distribution/validators/{validatorAddr}/outstanding_rewards][%d] getDistributionValidatorsValidatorAddrOutstandingRewardsOK  %+v", 200, o.Payload)
}

func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsOK) GetPayload() *GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody {
	return o.Payload
}

func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError creates a GetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError with default headers values
func NewGetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError() *GetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError {
	return &GetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError{}
}

/*GetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError struct {
}

func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /distribution/validators/{validatorAddr}/outstanding_rewards][%d] getDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError ", 500)
}

func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody get distribution validators validator addr outstanding rewards o k body
swagger:model GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody
*/
type GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result []*models.Coin `json:"result"`
}

// Validate validates this get distribution validators validator addr outstanding rewards o k body
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	for i := 0; i < len(o.Result); i++ {
		if swag.IsZero(o.Result[i]) { // not required
			continue
		}

		if o.Result[i] != nil {
			if err := o.Result[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionValidatorsValidatorAddrOutstandingRewardsOK" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody) UnmarshalBinary(b []byte) error {
	var res GetDistributionValidatorsValidatorAddrOutstandingRewardsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
