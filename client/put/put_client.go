package put

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new put API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for put API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
UpdateExport updates export group

Update an export group which includes:
    Add/Remove block objects (volumes, mirrors and snapshots)
    Add/remove clusters
    Add/remove hosts
    Add/remove initiators

Depending on the export group type (Initiator, Host or Cluster), the
request is restricted to enforce the same rules as
{@link #createExportGroup(ExportCreateParam)}:
    For initiator type groups, only initiators are accepted in the
    request. Further the initiators must be in the same host as the
    existing initiators.

    For host type groups, only hosts and initiators that belong to
    existing hosts will be accepted.

    For cluster type groups, only clusters, hosts and initiators will
    be accepted. Hosts and initiators must belong to existing clusters
    and hosts.

Note: The export group name, project and varray can not be modified.

*/
func (a *Client) UpdateExport(params *UpdateExportParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateExport",
		Method:             "PUT",
		PathPattern:        "/block/exports/{id}.json",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateExportReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateExportOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
