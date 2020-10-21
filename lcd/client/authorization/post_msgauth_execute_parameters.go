// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/terra-project/mantle-sdk/lcd/models"
)

// NewPostMsgauthExecuteParams creates a new PostMsgauthExecuteParams object
// with the default values initialized.
func NewPostMsgauthExecuteParams() *PostMsgauthExecuteParams {
	var ()
	return &PostMsgauthExecuteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMsgauthExecuteParamsWithTimeout creates a new PostMsgauthExecuteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMsgauthExecuteParamsWithTimeout(timeout time.Duration) *PostMsgauthExecuteParams {
	var ()
	return &PostMsgauthExecuteParams{

		timeout: timeout,
	}
}

// NewPostMsgauthExecuteParamsWithContext creates a new PostMsgauthExecuteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMsgauthExecuteParamsWithContext(ctx context.Context) *PostMsgauthExecuteParams {
	var ()
	return &PostMsgauthExecuteParams{

		Context: ctx,
	}
}

// NewPostMsgauthExecuteParamsWithHTTPClient creates a new PostMsgauthExecuteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMsgauthExecuteParamsWithHTTPClient(client *http.Client) *PostMsgauthExecuteParams {
	var ()
	return &PostMsgauthExecuteParams{
		HTTPClient: client,
	}
}

/*PostMsgauthExecuteParams contains all the parameters to send to the API endpoint
for the post msgauth execute operation typically these are written to a http.Request
*/
type PostMsgauthExecuteParams struct {

	/*RevokeRequestBody*/
	RevokeRequestBody *models.ExecuteGrantReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post msgauth execute params
func (o *PostMsgauthExecuteParams) WithTimeout(timeout time.Duration) *PostMsgauthExecuteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post msgauth execute params
func (o *PostMsgauthExecuteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post msgauth execute params
func (o *PostMsgauthExecuteParams) WithContext(ctx context.Context) *PostMsgauthExecuteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post msgauth execute params
func (o *PostMsgauthExecuteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post msgauth execute params
func (o *PostMsgauthExecuteParams) WithHTTPClient(client *http.Client) *PostMsgauthExecuteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post msgauth execute params
func (o *PostMsgauthExecuteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRevokeRequestBody adds the revokeRequestBody to the post msgauth execute params
func (o *PostMsgauthExecuteParams) WithRevokeRequestBody(revokeRequestBody *models.ExecuteGrantReq) *PostMsgauthExecuteParams {
	o.SetRevokeRequestBody(revokeRequestBody)
	return o
}

// SetRevokeRequestBody adds the revokeRequestBody to the post msgauth execute params
func (o *PostMsgauthExecuteParams) SetRevokeRequestBody(revokeRequestBody *models.ExecuteGrantReq) {
	o.RevokeRequestBody = revokeRequestBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostMsgauthExecuteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RevokeRequestBody != nil {
		if err := r.SetBodyParam(o.RevokeRequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
