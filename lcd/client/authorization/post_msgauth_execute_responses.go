// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/terra-money/mantle-sdk/lcd/models"
)

// PostMsgauthExecuteReader is a Reader for the PostMsgauthExecute structure.
type PostMsgauthExecuteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMsgauthExecuteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostMsgauthExecuteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostMsgauthExecuteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostMsgauthExecuteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostMsgauthExecuteOK creates a PostMsgauthExecuteOK with default headers values
func NewPostMsgauthExecuteOK() *PostMsgauthExecuteOK {
	return &PostMsgauthExecuteOK{}
}

/*PostMsgauthExecuteOK handles this case with default header values.

OK
*/
type PostMsgauthExecuteOK struct {
	Payload *models.StdTx
}

func (o *PostMsgauthExecuteOK) Error() string {
	return fmt.Sprintf("[POST /msgauth/execute][%d] postMsgauthExecuteOK  %+v", 200, o.Payload)
}

func (o *PostMsgauthExecuteOK) GetPayload() *models.StdTx {
	return o.Payload
}

func (o *PostMsgauthExecuteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StdTx)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMsgauthExecuteBadRequest creates a PostMsgauthExecuteBadRequest with default headers values
func NewPostMsgauthExecuteBadRequest() *PostMsgauthExecuteBadRequest {
	return &PostMsgauthExecuteBadRequest{}
}

/*PostMsgauthExecuteBadRequest handles this case with default header values.

Bad request
*/
type PostMsgauthExecuteBadRequest struct {
}

func (o *PostMsgauthExecuteBadRequest) Error() string {
	return fmt.Sprintf("[POST /msgauth/execute][%d] postMsgauthExecuteBadRequest ", 400)
}

func (o *PostMsgauthExecuteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMsgauthExecuteInternalServerError creates a PostMsgauthExecuteInternalServerError with default headers values
func NewPostMsgauthExecuteInternalServerError() *PostMsgauthExecuteInternalServerError {
	return &PostMsgauthExecuteInternalServerError{}
}

/*PostMsgauthExecuteInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostMsgauthExecuteInternalServerError struct {
}

func (o *PostMsgauthExecuteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /msgauth/execute][%d] postMsgauthExecuteInternalServerError ", 500)
}

func (o *PostMsgauthExecuteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
