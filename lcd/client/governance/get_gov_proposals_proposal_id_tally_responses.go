// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// GetGovProposalsProposalIDTallyReader is a Reader for the GetGovProposalsProposalIDTally structure.
type GetGovProposalsProposalIDTallyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGovProposalsProposalIDTallyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGovProposalsProposalIDTallyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGovProposalsProposalIDTallyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGovProposalsProposalIDTallyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGovProposalsProposalIDTallyOK creates a GetGovProposalsProposalIDTallyOK with default headers values
func NewGetGovProposalsProposalIDTallyOK() *GetGovProposalsProposalIDTallyOK {
	return &GetGovProposalsProposalIDTallyOK{}
}

/*GetGovProposalsProposalIDTallyOK handles this case with default header values.

OK
*/
type GetGovProposalsProposalIDTallyOK struct {
	Payload *GetGovProposalsProposalIDTallyOKBody
}

func (o *GetGovProposalsProposalIDTallyOK) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/tally][%d] getGovProposalsProposalIdTallyOK  %+v", 200, o.Payload)
}

func (o *GetGovProposalsProposalIDTallyOK) GetPayload() *GetGovProposalsProposalIDTallyOKBody {
	return o.Payload
}

func (o *GetGovProposalsProposalIDTallyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetGovProposalsProposalIDTallyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGovProposalsProposalIDTallyBadRequest creates a GetGovProposalsProposalIDTallyBadRequest with default headers values
func NewGetGovProposalsProposalIDTallyBadRequest() *GetGovProposalsProposalIDTallyBadRequest {
	return &GetGovProposalsProposalIDTallyBadRequest{}
}

/*GetGovProposalsProposalIDTallyBadRequest handles this case with default header values.

Invalid proposal id
*/
type GetGovProposalsProposalIDTallyBadRequest struct {
}

func (o *GetGovProposalsProposalIDTallyBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/tally][%d] getGovProposalsProposalIdTallyBadRequest ", 400)
}

func (o *GetGovProposalsProposalIDTallyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDTallyInternalServerError creates a GetGovProposalsProposalIDTallyInternalServerError with default headers values
func NewGetGovProposalsProposalIDTallyInternalServerError() *GetGovProposalsProposalIDTallyInternalServerError {
	return &GetGovProposalsProposalIDTallyInternalServerError{}
}

/*GetGovProposalsProposalIDTallyInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetGovProposalsProposalIDTallyInternalServerError struct {
}

func (o *GetGovProposalsProposalIDTallyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/tally][%d] getGovProposalsProposalIdTallyInternalServerError ", 500)
}

func (o *GetGovProposalsProposalIDTallyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetGovProposalsProposalIDTallyOKBody get gov proposals proposal ID tally o k body
swagger:model GetGovProposalsProposalIDTallyOKBody
*/
type GetGovProposalsProposalIDTallyOKBody struct {

	// height
	Height string `json:"height,omitempty"`

	// result
	Result *models.TallyResult `json:"result,omitempty"`
}

// Validate validates this get gov proposals proposal ID tally o k body
func (o *GetGovProposalsProposalIDTallyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGovProposalsProposalIDTallyOKBody) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(o.Result) { // not required
		return nil
	}

	if o.Result != nil {
		if err := o.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getGovProposalsProposalIdTallyOK" + "." + "result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGovProposalsProposalIDTallyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovProposalsProposalIDTallyOKBody) UnmarshalBinary(b []byte) error {
	var res GetGovProposalsProposalIDTallyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
