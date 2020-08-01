// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsReader is a Reader for the GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrants structure.
type GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK creates a GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK with default headers values
func NewGetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK() *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK {
	return &GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK{}
}

/*GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK handles this case with default header values.

grant informations (one of GenericGrantInfo and SendGrantInfo)
*/
type GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK struct {
	Payload []*GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0
}

func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK) Error() string {
	return fmt.Sprintf("[GET /msgauth/granters/{granterAddress}/grantees/{granteeAddress}/grants][%d] getMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK  %+v", 200, o.Payload)
}

func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK) GetPayload() []*GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0 {
	return o.Payload
}

func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0 get msgauth granters granter address grantees grantee address grants o k body items0
swagger:model GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0
*/
type GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0 struct {

	// generic grant info
	GenericGrantInfo *models.GenericGrantInfo `json:"GenericGrantInfo,omitempty"`

	// send grant info
	SendGrantInfo *models.SendGrantInfo `json:"SendGrantInfo,omitempty"`
}

// Validate validates this get msgauth granters granter address grantees grantee address grants o k body items0
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGenericGrantInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSendGrantInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0) validateGenericGrantInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.GenericGrantInfo) { // not required
		return nil
	}

	if o.GenericGrantInfo != nil {
		if err := o.GenericGrantInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GenericGrantInfo")
			}
			return err
		}
	}

	return nil
}

func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0) validateSendGrantInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.SendGrantInfo) { // not required
		return nil
	}

	if o.SendGrantInfo != nil {
		if err := o.SendGrantInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SendGrantInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetMsgauthGrantersGranterAddressGranteesGranteeAddressGrantsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}