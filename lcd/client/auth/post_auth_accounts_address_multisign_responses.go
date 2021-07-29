// Code generated by go-swagger; DO NOT EDIT.

package auth

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

// PostAuthAccountsAddressMultisignReader is a Reader for the PostAuthAccountsAddressMultisign structure.
type PostAuthAccountsAddressMultisignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthAccountsAddressMultisignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuthAccountsAddressMultisignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostAuthAccountsAddressMultisignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAuthAccountsAddressMultisignOK creates a PostAuthAccountsAddressMultisignOK with default headers values
func NewPostAuthAccountsAddressMultisignOK() *PostAuthAccountsAddressMultisignOK {
	return &PostAuthAccountsAddressMultisignOK{}
}

/*PostAuthAccountsAddressMultisignOK handles this case with default header values.

Generated multisigned tx or multisig
*/
type PostAuthAccountsAddressMultisignOK struct {
	Payload *PostAuthAccountsAddressMultisignOKBody
}

func (o *PostAuthAccountsAddressMultisignOK) Error() string {
	return fmt.Sprintf("[POST /auth/accounts/{address}/multisign][%d] postAuthAccountsAddressMultisignOK  %+v", 200, o.Payload)
}

func (o *PostAuthAccountsAddressMultisignOK) GetPayload() *PostAuthAccountsAddressMultisignOKBody {
	return o.Payload
}

func (o *PostAuthAccountsAddressMultisignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostAuthAccountsAddressMultisignOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthAccountsAddressMultisignInternalServerError creates a PostAuthAccountsAddressMultisignInternalServerError with default headers values
func NewPostAuthAccountsAddressMultisignInternalServerError() *PostAuthAccountsAddressMultisignInternalServerError {
	return &PostAuthAccountsAddressMultisignInternalServerError{}
}

/*PostAuthAccountsAddressMultisignInternalServerError handles this case with default header values.

Server internal error
*/
type PostAuthAccountsAddressMultisignInternalServerError struct {
}

func (o *PostAuthAccountsAddressMultisignInternalServerError) Error() string {
	return fmt.Sprintf("[POST /auth/accounts/{address}/multisign][%d] postAuthAccountsAddressMultisignInternalServerError ", 500)
}

func (o *PostAuthAccountsAddressMultisignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostAuthAccountsAddressMultisignOKBody post auth accounts address multisign o k body
swagger:model PostAuthAccountsAddressMultisignOKBody
*/
type PostAuthAccountsAddressMultisignOKBody struct {

	// multi sig
	MultiSig *models.StdSignature `json:"MultiSig,omitempty"`

	// multi signed tx
	MultiSignedTx *models.StdTx `json:"MultiSignedTx,omitempty"`
}

// Validate validates this post auth accounts address multisign o k body
func (o *PostAuthAccountsAddressMultisignOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMultiSig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMultiSignedTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostAuthAccountsAddressMultisignOKBody) validateMultiSig(formats strfmt.Registry) error {

	if swag.IsZero(o.MultiSig) { // not required
		return nil
	}

	if o.MultiSig != nil {
		if err := o.MultiSig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postAuthAccountsAddressMultisignOK" + "." + "MultiSig")
			}
			return err
		}
	}

	return nil
}

func (o *PostAuthAccountsAddressMultisignOKBody) validateMultiSignedTx(formats strfmt.Registry) error {

	if swag.IsZero(o.MultiSignedTx) { // not required
		return nil
	}

	if o.MultiSignedTx != nil {
		if err := o.MultiSignedTx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postAuthAccountsAddressMultisignOK" + "." + "MultiSignedTx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostAuthAccountsAddressMultisignOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAuthAccountsAddressMultisignOKBody) UnmarshalBinary(b []byte) error {
	var res PostAuthAccountsAddressMultisignOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
