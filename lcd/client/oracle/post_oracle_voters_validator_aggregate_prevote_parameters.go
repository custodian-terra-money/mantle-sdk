// Code generated by go-swagger; DO NOT EDIT.

package oracle

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

// NewPostOracleVotersValidatorAggregatePrevoteParams creates a new PostOracleVotersValidatorAggregatePrevoteParams object
// with the default values initialized.
func NewPostOracleVotersValidatorAggregatePrevoteParams() *PostOracleVotersValidatorAggregatePrevoteParams {
	var ()
	return &PostOracleVotersValidatorAggregatePrevoteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOracleVotersValidatorAggregatePrevoteParamsWithTimeout creates a new PostOracleVotersValidatorAggregatePrevoteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOracleVotersValidatorAggregatePrevoteParamsWithTimeout(timeout time.Duration) *PostOracleVotersValidatorAggregatePrevoteParams {
	var ()
	return &PostOracleVotersValidatorAggregatePrevoteParams{

		timeout: timeout,
	}
}

// NewPostOracleVotersValidatorAggregatePrevoteParamsWithContext creates a new PostOracleVotersValidatorAggregatePrevoteParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOracleVotersValidatorAggregatePrevoteParamsWithContext(ctx context.Context) *PostOracleVotersValidatorAggregatePrevoteParams {
	var ()
	return &PostOracleVotersValidatorAggregatePrevoteParams{

		Context: ctx,
	}
}

// NewPostOracleVotersValidatorAggregatePrevoteParamsWithHTTPClient creates a new PostOracleVotersValidatorAggregatePrevoteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOracleVotersValidatorAggregatePrevoteParamsWithHTTPClient(client *http.Client) *PostOracleVotersValidatorAggregatePrevoteParams {
	var ()
	return &PostOracleVotersValidatorAggregatePrevoteParams{
		HTTPClient: client,
	}
}

/*PostOracleVotersValidatorAggregatePrevoteParams contains all the parameters to send to the API endpoint
for the post oracle voters validator aggregate prevote operation typically these are written to a http.Request
*/
type PostOracleVotersValidatorAggregatePrevoteParams struct {

	/*AggregatePrevoteRequestBody*/
	AggregatePrevoteRequestBody *models.AggregatePrevoteReq
	/*Validator
	  oracle operator

	*/
	Validator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) WithTimeout(timeout time.Duration) *PostOracleVotersValidatorAggregatePrevoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) WithContext(ctx context.Context) *PostOracleVotersValidatorAggregatePrevoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) WithHTTPClient(client *http.Client) *PostOracleVotersValidatorAggregatePrevoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAggregatePrevoteRequestBody adds the aggregatePrevoteRequestBody to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) WithAggregatePrevoteRequestBody(aggregatePrevoteRequestBody *models.AggregatePrevoteReq) *PostOracleVotersValidatorAggregatePrevoteParams {
	o.SetAggregatePrevoteRequestBody(aggregatePrevoteRequestBody)
	return o
}

// SetAggregatePrevoteRequestBody adds the aggregatePrevoteRequestBody to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) SetAggregatePrevoteRequestBody(aggregatePrevoteRequestBody *models.AggregatePrevoteReq) {
	o.AggregatePrevoteRequestBody = aggregatePrevoteRequestBody
}

// WithValidator adds the validator to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) WithValidator(validator string) *PostOracleVotersValidatorAggregatePrevoteParams {
	o.SetValidator(validator)
	return o
}

// SetValidator adds the validator to the post oracle voters validator aggregate prevote params
func (o *PostOracleVotersValidatorAggregatePrevoteParams) SetValidator(validator string) {
	o.Validator = validator
}

// WriteToRequest writes these params to a swagger request
func (o *PostOracleVotersValidatorAggregatePrevoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AggregatePrevoteRequestBody != nil {
		if err := r.SetBodyParam(o.AggregatePrevoteRequestBody); err != nil {
			return err
		}
	}

	// path param validator
	if err := r.SetPathParam("validator", o.Validator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
