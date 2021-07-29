// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// PostGovProposalsRewardWeightUpdateReader is a Reader for the PostGovProposalsRewardWeightUpdate structure.
type PostGovProposalsRewardWeightUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGovProposalsRewardWeightUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGovProposalsRewardWeightUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGovProposalsRewardWeightUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGovProposalsRewardWeightUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGovProposalsRewardWeightUpdateOK creates a PostGovProposalsRewardWeightUpdateOK with default headers values
func NewPostGovProposalsRewardWeightUpdateOK() *PostGovProposalsRewardWeightUpdateOK {
	return &PostGovProposalsRewardWeightUpdateOK{}
}

/*PostGovProposalsRewardWeightUpdateOK handles this case with default header values.

The transaction was successfully generated
*/
type PostGovProposalsRewardWeightUpdateOK struct {
	Payload *models.StdTx
}

func (o *PostGovProposalsRewardWeightUpdateOK) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/reward_weight_update][%d] postGovProposalsRewardWeightUpdateOK  %+v", 200, o.Payload)
}

func (o *PostGovProposalsRewardWeightUpdateOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *PostGovProposalsRewardWeightUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGovProposalsRewardWeightUpdateBadRequest creates a PostGovProposalsRewardWeightUpdateBadRequest with default headers values
func NewPostGovProposalsRewardWeightUpdateBadRequest() *PostGovProposalsRewardWeightUpdateBadRequest {
	return &PostGovProposalsRewardWeightUpdateBadRequest{}
}

/*PostGovProposalsRewardWeightUpdateBadRequest handles this case with default header values.

Invalid proposal body
*/
type PostGovProposalsRewardWeightUpdateBadRequest struct {
}

func (o *PostGovProposalsRewardWeightUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/reward_weight_update][%d] postGovProposalsRewardWeightUpdateBadRequest ", 400)
}

func (o *PostGovProposalsRewardWeightUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostGovProposalsRewardWeightUpdateInternalServerError creates a PostGovProposalsRewardWeightUpdateInternalServerError with default headers values
func NewPostGovProposalsRewardWeightUpdateInternalServerError() *PostGovProposalsRewardWeightUpdateInternalServerError {
	return &PostGovProposalsRewardWeightUpdateInternalServerError{}
}

/*PostGovProposalsRewardWeightUpdateInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostGovProposalsRewardWeightUpdateInternalServerError struct {
}

func (o *PostGovProposalsRewardWeightUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/reward_weight_update][%d] postGovProposalsRewardWeightUpdateInternalServerError ", 500)
}

func (o *PostGovProposalsRewardWeightUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostGovProposalsRewardWeightUpdateBody post gov proposals reward weight update body
swagger:model PostGovProposalsRewardWeightUpdateBody
*/
type PostGovProposalsRewardWeightUpdateBody struct {

	// base req
	BaseReq *models.BaseReq `json:"base_req,omitempty"`

	// deposit
	Deposit []*models.Coin `json:"deposit"`

	// description
	Description string `json:"description,omitempty"`

	// proposer
	Proposer models.Address `json:"proposer,omitempty"`

	// reward weight
	RewardWeight float32 `json:"reward_weight,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this post gov proposals reward weight update body
func (o *PostGovProposalsRewardWeightUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProposer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsRewardWeightUpdateBody) validateBaseReq(formats strfmt.Registry) error {

	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_proposal_body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostGovProposalsRewardWeightUpdateBody) validateDeposit(formats strfmt.Registry) error {

	if swag.IsZero(o.Deposit) { // not required
		return nil
	}

	for i := 0; i < len(o.Deposit); i++ {
		if swag.IsZero(o.Deposit[i]) { // not required
			continue
		}

		if o.Deposit[i] != nil {
			if err := o.Deposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostGovProposalsRewardWeightUpdateBody) validateProposer(formats strfmt.Registry) error {

	if swag.IsZero(o.Proposer) { // not required
		return nil
	}

	if err := o.Proposer.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("post_proposal_body" + "." + "proposer")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsRewardWeightUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsRewardWeightUpdateBody) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsRewardWeightUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
