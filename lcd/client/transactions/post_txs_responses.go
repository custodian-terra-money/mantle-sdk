// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// PostTxsReader is a Reader for the PostTxs structure.
type PostTxsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTxsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTxsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostTxsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTxsOK creates a PostTxsOK with default headers values
func NewPostTxsOK() *PostTxsOK {
	return &PostTxsOK{}
}

/*PostTxsOK handles this case with default header values.

Tx broadcasting result
*/
type PostTxsOK struct {
	Payload *models.BroadcastTxCommitResult
}

func (o *PostTxsOK) Error() string {
	return fmt.Sprintf("[POST /txs][%d] postTxsOK  %+v", 200, o.Payload)
}

func (o *PostTxsOK) GetPayload() *models.BroadcastTxCommitResult {
	return o.Payload
}

func (o *PostTxsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BroadcastTxCommitResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTxsInternalServerError creates a PostTxsInternalServerError with default headers values
func NewPostTxsInternalServerError() *PostTxsInternalServerError {
	return &PostTxsInternalServerError{}
}

/*PostTxsInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostTxsInternalServerError struct {
}

func (o *PostTxsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /txs][%d] postTxsInternalServerError ", 500)
}

func (o *PostTxsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostTxsBody post txs body
swagger:model PostTxsBody
*/
type PostTxsBody struct {

	// mode
	Mode string `json:"mode,omitempty"`

	// tx
	Tx *models.StdTx `json:"tx,omitempty"`
}

// Validate validates this post txs body
func (o *PostTxsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsBody) validateTx(formats strfmt.Registry) error {

	if swag.IsZero(o.Tx) { // not required
		return nil
	}

	if o.Tx != nil {
		if err := o.Tx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("txBroadcast" + "." + "tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsBody) UnmarshalBinary(b []byte) error {
	var res PostTxsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
