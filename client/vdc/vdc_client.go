package vdc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vdc API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vdc API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ListTasks lists tasks

Returns a list of tasks for the specified tenant

*/
func (a *Client) ListTasks(params *ListTasksParams, authInfo runtime.ClientAuthInfoWriter) (*ListTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListTasks",
		Method:             "GET",
		PathPattern:        "/vdc/tasks.json",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTasksOK), nil
}

/*
ShowTask shows task

Show Task details.

*/
func (a *Client) ShowTask(params *ShowTaskParams, authInfo runtime.ClientAuthInfoWriter) (*ShowTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ShowTask",
		Method:             "GET",
		PathPattern:        "/vdc/tasks/{id}.json",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ShowTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ShowTaskOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
