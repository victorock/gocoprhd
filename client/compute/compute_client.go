package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new compute API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for compute API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateExport creates export

Block export method is use to export one or more volumes to one or more
hosts. This is a required step for a host to be able to access a block
volume, although in some scenarios, additional configurations may be
required. There are three main types of export group to meet the common
use cases:

Create an initiator type export group so that a single host can see one
or more volumes. An example would be an export group for a host boot lun
or a private volume that is meant to be used by only one host. The
assumption is, in this case the user wants the boot or private volume
to be accessed via known initiators. For this type of export, the
request object is expected to have only initiators (i.e. no hosts or
clusters). Further, the initiators are expected to belong to the same
host. While an initiator type export group can belong to only one host,
this does not mean the host can only have the initiator type export
group. A hosts can be part of many export groups of any type. The export
group type {@link ExportGroupType#Initiator} should be specified in the
request for this type of export.

Create an export group so that one or more hosts, which are not part of
a cluster, can access one or more volumes. This is the use case of a
shared data lun. In this case, it is assumed that the user wants all the
hosts initiators that are connected to the storage array (up to the
maximum specified by the virtual pool) to be able to access the volume.
The export group type {@link ExportGroupType#Host} should be specified
in the request for this type of export.

Create an export group so that one or more clusters of hosts can access
one or more volumes. This is the same use case of shared data lun as the
{@link ExportGroupType#Host} use case with the exception that the user
is managing a cluster of hosts as opposed to individual hosts. In this
case, the same assumption about the initiators as in the previous case
is made. The export group type {@link ExportGroupType#Cluster} should be
specified in the request for this type of export.

Note that the above discussion only mentions volumes but mirrors and
snapshots can also be used in export groups.

Once a block export is created, following incremental changes can be
applied to it:
  - add volume or volume snapshot to the shared storage pool
  - remove volume or volume snapshot from the shared storage pool
  - add new server to the cluster by adding initiator from that server
    to the block export - remove visibility of shared storage to a
    server by removing initiators from the block export

Similar to block storage provisioning, block export is also created
within the scope of a varray. Hence, volumes and snapshots being added
to a block export must belong to the same varray. Fibre Channel and
iSCSI initiators must be part of SANs belonging to the same varray as
block export.

For Fibre Channel initiators, SAN zones will also be created when the
export group is created if the networks are discovered and:

at least one of the Network Systems can provision the Vsan or Fabric in
which the each endpoint exists, and the VirtualArray has
"auto_san_zoning" set to true.

The SAN zones each consists of an initiator (from the arguments) and a
storage port that is selected. The number of zones created will be
determined from the number of required initiator/storage-port
communication paths.

*/
func (a *Client) CreateExport(params *CreateExportParams, authInfo runtime.ClientAuthInfoWriter) (*CreateExportAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateExport",
		Method:             "POST",
		PathPattern:        "/block/exports.json",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateExportReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateExportAccepted), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
