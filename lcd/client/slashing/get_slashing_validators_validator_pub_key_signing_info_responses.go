// Code generated by go-swagger; DO NOT EDIT.

package slashing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/terra-project/mantle/lcd/models"
)

// GetSlashingValidatorsValidatorPubKeySigningInfoReader is a Reader for the GetSlashingValidatorsValidatorPubKeySigningInfo structure.
type GetSlashingValidatorsValidatorPubKeySigningInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSlashingValidatorsValidatorPubKeySigningInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetSlashingValidatorsValidatorPubKeySigningInfoNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSlashingValidatorsValidatorPubKeySigningInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSlashingValidatorsValidatorPubKeySigningInfoOK creates a GetSlashingValidatorsValidatorPubKeySigningInfoOK with default headers values
func NewGetSlashingValidatorsValidatorPubKeySigningInfoOK() *GetSlashingValidatorsValidatorPubKeySigningInfoOK {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoOK{}
}

/*GetSlashingValidatorsValidatorPubKeySigningInfoOK handles this case with default header values.

OK
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoOK struct {
	Payload *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOK) Error() string {
	return fmt.Sprintf("[GET /slashing/validators/{validatorPubKey}/signing_info][%d] getSlashingValidatorsValidatorPubKeySigningInfoOK  %+v", 200, o.Payload)
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOK) GetPayload() *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody {
	return o.Payload
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSlashingValidatorsValidatorPubKeySigningInfoOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSlashingValidatorsValidatorPubKeySigningInfoNoContent creates a GetSlashingValidatorsValidatorPubKeySigningInfoNoContent with default headers values
func NewGetSlashingValidatorsValidatorPubKeySigningInfoNoContent() *GetSlashingValidatorsValidatorPubKeySigningInfoNoContent {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoNoContent{}
}

/*GetSlashingValidatorsValidatorPubKeySigningInfoNoContent handles this case with default header values.

No sign info of this validator
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoNoContent struct {
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoNoContent) Error() string {
	return fmt.Sprintf("[GET /slashing/validators/{validatorPubKey}/signing_info][%d] getSlashingValidatorsValidatorPubKeySigningInfoNoContent ", 204)
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSlashingValidatorsValidatorPubKeySigningInfoBadRequest creates a GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest with default headers values
func NewGetSlashingValidatorsValidatorPubKeySigningInfoBadRequest() *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest{}
}

/*GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest handles this case with default header values.

Invalid validator public key
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest struct {
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest) Error() string {
	return fmt.Sprintf("[GET /slashing/validators/{validatorPubKey}/signing_info][%d] getSlashingValidatorsValidatorPubKeySigningInfoBadRequest ", 400)
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError creates a GetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError with default headers values
func NewGetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError() *GetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError {
	return &GetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError{}
}

/*GetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError struct {
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /slashing/validators/{validatorPubKey}/signing_info][%d] getSlashingValidatorsValidatorPubKeySigningInfoInternalServerError ", 500)
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetSlashingValidatorsValidatorPubKeySigningInfoOKBody get slashing validators validator pub key signing info o k body
swagger:model GetSlashingValidatorsValidatorPubKeySigningInfoOKBody
*/
type GetSlashingValidatorsValidatorPubKeySigningInfoOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *models.SigningInfo `json:"result,omitempty"`
}

// Validate validates this get slashing validators validator pub key signing info o k body
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSlashingValidatorsValidatorPubKeySigningInfoOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSlashingValidatorsValidatorPubKeySigningInfoOKBody) UnmarshalBinary(b []byte) error {
	var res GetSlashingValidatorsValidatorPubKeySigningInfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
