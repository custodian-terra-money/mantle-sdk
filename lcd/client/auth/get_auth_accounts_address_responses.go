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

// GetAuthAccountsAddressReader is a Reader for the GetAuthAccountsAddress structure.
type GetAuthAccountsAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthAccountsAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthAccountsAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetAuthAccountsAddressNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAuthAccountsAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthAccountsAddressOK creates a GetAuthAccountsAddressOK with default headers values
func NewGetAuthAccountsAddressOK() *GetAuthAccountsAddressOK {
	return &GetAuthAccountsAddressOK{}
}

/*GetAuthAccountsAddressOK handles this case with default header values.

Account information on the blockchain (one of Account and LazyGradedVestingAccount)
*/
type GetAuthAccountsAddressOK struct {
	Payload *GetAuthAccountsAddressOKBody
}

func (o *GetAuthAccountsAddressOK) Error() string {
	return fmt.Sprintf("[GET /auth/accounts/{address}][%d] getAuthAccountsAddressOK  %+v", 200, o.Payload)
}

func (o *GetAuthAccountsAddressOK) GetPayload() *GetAuthAccountsAddressOKBody {
	return o.Payload
}

func (o *GetAuthAccountsAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAuthAccountsAddressOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthAccountsAddressNoContent creates a GetAuthAccountsAddressNoContent with default headers values
func NewGetAuthAccountsAddressNoContent() *GetAuthAccountsAddressNoContent {
	return &GetAuthAccountsAddressNoContent{}
}

/*GetAuthAccountsAddressNoContent handles this case with default header values.

No content about this account address
*/
type GetAuthAccountsAddressNoContent struct {
}

func (o *GetAuthAccountsAddressNoContent) Error() string {
	return fmt.Sprintf("[GET /auth/accounts/{address}][%d] getAuthAccountsAddressNoContent ", 204)
}

func (o *GetAuthAccountsAddressNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAuthAccountsAddressInternalServerError creates a GetAuthAccountsAddressInternalServerError with default headers values
func NewGetAuthAccountsAddressInternalServerError() *GetAuthAccountsAddressInternalServerError {
	return &GetAuthAccountsAddressInternalServerError{}
}

/*GetAuthAccountsAddressInternalServerError handles this case with default header values.

Server internal error
*/
type GetAuthAccountsAddressInternalServerError struct {
}

func (o *GetAuthAccountsAddressInternalServerError) Error() string {
	return fmt.Sprintf("[GET /auth/accounts/{address}][%d] getAuthAccountsAddressInternalServerError ", 500)
}

func (o *GetAuthAccountsAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetAuthAccountsAddressOKBody get auth accounts address o k body
swagger:model GetAuthAccountsAddressOKBody
*/
type GetAuthAccountsAddressOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *GetAuthAccountsAddressOKBodyResult `json:"result,omitempty"`
}

// Validate validates this get auth accounts address o k body
func (o *GetAuthAccountsAddressOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAuthAccountsAddressOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAuthAccountsAddressOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAuthAccountsAddressOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAuthAccountsAddressOKBody) UnmarshalBinary(b []byte) error {
	var res GetAuthAccountsAddressOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAuthAccountsAddressOKBodyResult get auth accounts address o k body result
swagger:model GetAuthAccountsAddressOKBodyResult
*/
type GetAuthAccountsAddressOKBodyResult struct {

	// account
	Account *models.Account `json:"Account,omitempty"`

	// lazy graded vesting account
	LazyGradedVestingAccount *models.LazyGradedVestingAccount `json:"LazyGradedVestingAccount,omitempty"`
}

// Validate validates this get auth accounts address o k body result
func (o *GetAuthAccountsAddressOKBodyResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLazyGradedVestingAccount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAuthAccountsAddressOKBodyResult) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(o.Account) { // not required
		return nil
	}

	if o.Account != nil {
		if err := o.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAuthAccountsAddressOK" + "." + "result" + "." + "Account")
			}
			return err
		}
	}

	return nil
}

func (o *GetAuthAccountsAddressOKBodyResult) validateLazyGradedVestingAccount(formats strfmt.Registry) error {

	if swag.IsZero(o.LazyGradedVestingAccount) { // not required
		return nil
	}

	if o.LazyGradedVestingAccount != nil {
		if err := o.LazyGradedVestingAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAuthAccountsAddressOK" + "." + "result" + "." + "LazyGradedVestingAccount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAuthAccountsAddressOKBodyResult) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAuthAccountsAddressOKBodyResult) UnmarshalBinary(b []byte) error {
	var res GetAuthAccountsAddressOKBodyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
