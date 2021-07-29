// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// PostTxsEstimateFeeReader is a Reader for the PostTxsEstimateFee structure.
type PostTxsEstimateFeeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTxsEstimateFeeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTxsEstimateFeeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTxsEstimateFeeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTxsEstimateFeeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTxsEstimateFeeOK creates a PostTxsEstimateFeeOK with default headers values
func NewPostTxsEstimateFeeOK() *PostTxsEstimateFeeOK {
	return &PostTxsEstimateFeeOK{}
}

/*PostTxsEstimateFeeOK handles this case with default header values.

Estimated gas and fees object
*/
type PostTxsEstimateFeeOK struct {
	Payload *models.EstimateFeeResp
}

func (o *PostTxsEstimateFeeOK) Error() string {
	return fmt.Sprintf("[POST /txs/estimate_fee][%d] postTxsEstimateFeeOK  %+v", 200, o.Payload)
}

func (o *PostTxsEstimateFeeOK) GetPayload() *models.EstimateFeeResp {
	return o.Payload
}

func (o *PostTxsEstimateFeeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EstimateFeeResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTxsEstimateFeeBadRequest creates a PostTxsEstimateFeeBadRequest with default headers values
func NewPostTxsEstimateFeeBadRequest() *PostTxsEstimateFeeBadRequest {
	return &PostTxsEstimateFeeBadRequest{}
}

/*PostTxsEstimateFeeBadRequest handles this case with default header values.

The tx was malformed
*/
type PostTxsEstimateFeeBadRequest struct {
}

func (o *PostTxsEstimateFeeBadRequest) Error() string {
	return fmt.Sprintf("[POST /txs/estimate_fee][%d] postTxsEstimateFeeBadRequest ", 400)
}

func (o *PostTxsEstimateFeeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTxsEstimateFeeInternalServerError creates a PostTxsEstimateFeeInternalServerError with default headers values
func NewPostTxsEstimateFeeInternalServerError() *PostTxsEstimateFeeInternalServerError {
	return &PostTxsEstimateFeeInternalServerError{}
}

/*PostTxsEstimateFeeInternalServerError handles this case with default header values.

Server internal error
*/
type PostTxsEstimateFeeInternalServerError struct {
}

func (o *PostTxsEstimateFeeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /txs/estimate_fee][%d] postTxsEstimateFeeInternalServerError ", 500)
}

func (o *PostTxsEstimateFeeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostTxsEstimateFeeBody post txs estimate fee body
swagger:model PostTxsEstimateFeeBody
*/
type PostTxsEstimateFeeBody struct {

	// only used when tx's gas is "0", default "1"
	GasAdjustment string `json:"gas_adjustment,omitempty"`

	// used to compute gas fees = (gas * gas_prices), default 0
	GasPrices []*models.DecCoin `json:"gas_prices"`

	// tx
	Tx *models.StdTx `json:"tx,omitempty"`
}

// Validate validates this post txs estimate fee body
func (o *PostTxsEstimateFeeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGasPrices(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsEstimateFeeBody) validateGasPrices(formats strfmt.Registry) error {

	if swag.IsZero(o.GasPrices) { // not required
		return nil
	}

	for i := 0; i < len(o.GasPrices); i++ {
		if swag.IsZero(o.GasPrices[i]) { // not required
			continue
		}

		if o.GasPrices[i] != nil {
			if err := o.GasPrices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("estimate_req" + "." + "gas_prices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostTxsEstimateFeeBody) validateTx(formats strfmt.Registry) error {

	if swag.IsZero(o.Tx) { // not required
		return nil
	}

	if o.Tx != nil {
		if err := o.Tx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("estimate_req" + "." + "tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsEstimateFeeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsEstimateFeeBody) UnmarshalBinary(b []byte) error {
	var res PostTxsEstimateFeeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
