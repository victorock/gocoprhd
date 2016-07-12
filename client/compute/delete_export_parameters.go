package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteExportParams creates a new DeleteExportParams object
// with the default values initialized.
func NewDeleteExportParams() *DeleteExportParams {
	var ()
	return &DeleteExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExportParamsWithTimeout creates a new DeleteExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteExportParamsWithTimeout(timeout time.Duration) *DeleteExportParams {
	var ()
	return &DeleteExportParams{

		timeout: timeout,
	}
}

/*DeleteExportParams contains all the parameters to send to the API endpoint
for the delete export operation typically these are written to a http.Request
*/
type DeleteExportParams struct {

	/*ID
	  The URN of a ViPR/CoprHD Export

	*/
	ID string

	timeout time.Duration
}

// WithID adds the id to the delete export params
func (o *DeleteExportParams) WithID(ID string) *DeleteExportParams {
	o.ID = ID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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