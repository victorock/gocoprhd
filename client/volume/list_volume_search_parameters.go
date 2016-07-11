package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListVolumeSearchParams creates a new ListVolumeSearchParams object
// with the default values initialized.
func NewListVolumeSearchParams() *ListVolumeSearchParams {
	var ()
	return &ListVolumeSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListVolumeSearchParamsWithTimeout creates a new ListVolumeSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListVolumeSearchParamsWithTimeout(timeout time.Duration) *ListVolumeSearchParams {
	var ()
	return &ListVolumeSearchParams{

		timeout: timeout,
	}
}

/*ListVolumeSearchParams contains all the parameters to send to the API endpoint
for the list volume search operation typically these are written to a http.Request
*/
type ListVolumeSearchParams struct {

	/*Item
	  Item to search

	*/
	Item string
	/*Name
	  String to match

	*/
	Name string

	timeout time.Duration
}

// WithItem adds the item to the list volume search params
func (o *ListVolumeSearchParams) WithItem(Item string) *ListVolumeSearchParams {
	o.Item = Item
	return o
}

// WithName adds the name to the list volume search params
func (o *ListVolumeSearchParams) WithName(Name string) *ListVolumeSearchParams {
	o.Name = Name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListVolumeSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param item
	if err := r.SetPathParam("item", o.Item); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
