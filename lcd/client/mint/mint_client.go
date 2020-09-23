// Code generated by go-swagger; DO NOT EDIT.

package mint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new mint API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mint API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetMintingAnnualProvisions(params *GetMintingAnnualProvisionsParams) (*GetMintingAnnualProvisionsOK, error)

	GetMintingInflation(params *GetMintingInflationParams) (*GetMintingInflationOK, error)

	GetMintingParameters(params *GetMintingParametersParams) (*GetMintingParametersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetMintingAnnualProvisions currents minting annual provisions value
*/
func (a *Client) GetMintingAnnualProvisions(params *GetMintingAnnualProvisionsParams) (*GetMintingAnnualProvisionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMintingAnnualProvisionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMintingAnnualProvisions",
		Method:             "GET",
		PathPattern:        "/minting/annual-provisions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMintingAnnualProvisionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMintingAnnualProvisionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMintingAnnualProvisions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMintingInflation currents minting inflation value
*/
func (a *Client) GetMintingInflation(params *GetMintingInflationParams) (*GetMintingInflationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMintingInflationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMintingInflation",
		Method:             "GET",
		PathPattern:        "/minting/inflation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMintingInflationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMintingInflationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMintingInflation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMintingParameters mintings module parameters
*/
func (a *Client) GetMintingParameters(params *GetMintingParametersParams) (*GetMintingParametersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMintingParametersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMintingParameters",
		Method:             "GET",
		PathPattern:        "/minting/parameters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMintingParametersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMintingParametersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMintingParameters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
