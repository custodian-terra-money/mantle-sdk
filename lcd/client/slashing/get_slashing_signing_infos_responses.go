// Code generated by go-swagger; DO NOT EDIT.

package slashing

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

// GetSlashingSigningInfosReader is a Reader for the GetSlashingSigningInfos structure.
type GetSlashingSigningInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSlashingSigningInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSlashingSigningInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSlashingSigningInfosBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSlashingSigningInfosInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSlashingSigningInfosOK creates a GetSlashingSigningInfosOK with default headers values
func NewGetSlashingSigningInfosOK() *GetSlashingSigningInfosOK {
	return &GetSlashingSigningInfosOK{}
}

/*GetSlashingSigningInfosOK handles this case with default header values.

OK
*/
type GetSlashingSigningInfosOK struct {
	Payload *GetSlashingSigningInfosOKBody
}

func (o *GetSlashingSigningInfosOK) Error() string {
	return fmt.Sprintf("[GET /slashing/signing_infos][%d] getSlashingSigningInfosOK  %+v", 200, o.Payload)
}

func (o *GetSlashingSigningInfosOK) GetPayload() *GetSlashingSigningInfosOKBody {
	return o.Payload
}

func (o *GetSlashingSigningInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSlashingSigningInfosOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSlashingSigningInfosBadRequest creates a GetSlashingSigningInfosBadRequest with default headers values
func NewGetSlashingSigningInfosBadRequest() *GetSlashingSigningInfosBadRequest {
	return &GetSlashingSigningInfosBadRequest{}
}

/*GetSlashingSigningInfosBadRequest handles this case with default header values.

Invalid validator public key for one of the validators
*/
type GetSlashingSigningInfosBadRequest struct {
}

func (o *GetSlashingSigningInfosBadRequest) Error() string {
	return fmt.Sprintf("[GET /slashing/signing_infos][%d] getSlashingSigningInfosBadRequest ", 400)
}

func (o *GetSlashingSigningInfosBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSlashingSigningInfosInternalServerError creates a GetSlashingSigningInfosInternalServerError with default headers values
func NewGetSlashingSigningInfosInternalServerError() *GetSlashingSigningInfosInternalServerError {
	return &GetSlashingSigningInfosInternalServerError{}
}

/*GetSlashingSigningInfosInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetSlashingSigningInfosInternalServerError struct {
}

func (o *GetSlashingSigningInfosInternalServerError) Error() string {
	return fmt.Sprintf("[GET /slashing/signing_infos][%d] getSlashingSigningInfosInternalServerError ", 500)
}

func (o *GetSlashingSigningInfosInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetSlashingSigningInfosOKBody get slashing signing infos o k body
swagger:model GetSlashingSigningInfosOKBody
*/
type GetSlashingSigningInfosOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result []*models.SigningInfo `json:"result"`
}

// Validate validates this get slashing signing infos o k body
func (o *GetSlashingSigningInfosOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSlashingSigningInfosOKBody) validateResult(formats strfmt.Registry) error {

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
					return ve.ValidateName("getSlashingSigningInfosOK" + "." + "result" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSlashingSigningInfosOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSlashingSigningInfosOKBody) UnmarshalBinary(b []byte) error {
	var res GetSlashingSigningInfosOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
