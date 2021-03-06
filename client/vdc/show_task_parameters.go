package vdc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowTaskParams creates a new ShowTaskParams object
// with the default values initialized.
func NewShowTaskParams() *ShowTaskParams {
	var ()
	return &ShowTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowTaskParamsWithTimeout creates a new ShowTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowTaskParamsWithTimeout(timeout time.Duration) *ShowTaskParams {
	var ()
	return &ShowTaskParams{

		timeout: timeout,
	}
}

/*ShowTaskParams contains all the parameters to send to the API endpoint
for the show task operation typically these are written to a http.Request
*/
type ShowTaskParams struct {

	/*ID
	  The URN of a ViPR/CoprHD Task

	*/
	ID string

	timeout time.Duration
}

// WithID adds the id to the show task params
func (o *ShowTaskParams) WithID(ID string) *ShowTaskParams {
	o.ID = ID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ShowTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
