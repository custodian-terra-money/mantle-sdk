// Code generated by go-swagger; DO NOT EDIT.

package oracle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-money/mantle-sdk/lcd/models"
)

// PostOracleVotersValidatorAggregateVoteReader is a Reader for the PostOracleVotersValidatorAggregateVote structure.
type PostOracleVotersValidatorAggregateVoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOracleVotersValidatorAggregateVoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOracleVotersValidatorAggregateVoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOracleVotersValidatorAggregateVoteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOracleVotersValidatorAggregateVoteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOracleVotersValidatorAggregateVoteOK creates a PostOracleVotersValidatorAggregateVoteOK with default headers values
func NewPostOracleVotersValidatorAggregateVoteOK() *PostOracleVotersValidatorAggregateVoteOK {
	return &PostOracleVotersValidatorAggregateVoteOK{}
}

/*PostOracleVotersValidatorAggregateVoteOK handles this case with default header values.

OK
*/
type PostOracleVotersValidatorAggregateVoteOK struct {
	Payload *models.StdTx
}

func (o *PostOracleVotersValidatorAggregateVoteOK) Error() string {
	return fmt.Sprintf("[POST /oracle/voters/{validator}/aggregate_vote][%d] postOracleVotersValidatorAggregateVoteOK  %+v", 200, o.Payload)
}

func (o *PostOracleVotersValidatorAggregateVoteOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *PostOracleVotersValidatorAggregateVoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOracleVotersValidatorAggregateVoteBadRequest creates a PostOracleVotersValidatorAggregateVoteBadRequest with default headers values
func NewPostOracleVotersValidatorAggregateVoteBadRequest() *PostOracleVotersValidatorAggregateVoteBadRequest {
	return &PostOracleVotersValidatorAggregateVoteBadRequest{}
}

/*PostOracleVotersValidatorAggregateVoteBadRequest handles this case with default header values.

Bad request
*/
type PostOracleVotersValidatorAggregateVoteBadRequest struct {
}

func (o *PostOracleVotersValidatorAggregateVoteBadRequest) Error() string {
	return fmt.Sprintf("[POST /oracle/voters/{validator}/aggregate_vote][%d] postOracleVotersValidatorAggregateVoteBadRequest ", 400)
}

func (o *PostOracleVotersValidatorAggregateVoteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostOracleVotersValidatorAggregateVoteInternalServerError creates a PostOracleVotersValidatorAggregateVoteInternalServerError with default headers values
func NewPostOracleVotersValidatorAggregateVoteInternalServerError() *PostOracleVotersValidatorAggregateVoteInternalServerError {
	return &PostOracleVotersValidatorAggregateVoteInternalServerError{}
}

/*PostOracleVotersValidatorAggregateVoteInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostOracleVotersValidatorAggregateVoteInternalServerError struct {
}

func (o *PostOracleVotersValidatorAggregateVoteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oracle/voters/{validator}/aggregate_vote][%d] postOracleVotersValidatorAggregateVoteInternalServerError ", 500)
}

func (o *PostOracleVotersValidatorAggregateVoteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
