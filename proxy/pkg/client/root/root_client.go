// Code generated by go-swagger; DO NOT EDIT.

package root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new root API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for root API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	RootConfigList(params *RootConfigListParams, authInfo runtime.ClientAuthInfoWriter) (*RootConfigListOK, error)

	RootMessagesList(params *RootMessagesListParams, authInfo runtime.ClientAuthInfoWriter) (*RootMessagesListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  RootConfigList Retrieve passbook configurations
*/
func (a *Client) RootConfigList(params *RootConfigListParams, authInfo runtime.ClientAuthInfoWriter) (*RootConfigListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRootConfigListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "root_config_list",
		Method:             "GET",
		PathPattern:        "/root/config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RootConfigListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RootConfigListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for root_config_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RootMessagesList List current messages and pass into Serializer
*/
func (a *Client) RootMessagesList(params *RootMessagesListParams, authInfo runtime.ClientAuthInfoWriter) (*RootMessagesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRootMessagesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "root_messages_list",
		Method:             "GET",
		PathPattern:        "/root/messages/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RootMessagesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RootMessagesListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for root_messages_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
