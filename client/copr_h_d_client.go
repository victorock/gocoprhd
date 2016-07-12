package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/victorock/gocoprhd/client/authentication"
	"github.com/victorock/gocoprhd/client/block"
	"github.com/victorock/gocoprhd/client/vdc"
)

// Default copr h d HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new copr h d HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CoprHD {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/", []string{"https", "http"})
	return New(transport, formats)
}

// New creates a new copr h d client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CoprHD {
	cli := new(CoprHD)
	cli.Transport = transport

	cli.Authentication = authentication.New(transport, formats)

	cli.Block = block.New(transport, formats)

	cli.Vdc = vdc.New(transport, formats)

	return cli
}

// CoprHD is a client for copr h d
type CoprHD struct {
	Authentication *authentication.Client

	Block *block.Client

	Vdc *vdc.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CoprHD) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Authentication.SetTransport(transport)

	c.Block.SetTransport(transport)

	c.Vdc.SetTransport(transport)

}
