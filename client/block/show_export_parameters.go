package block

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewShowExportParams creates a new ShowExportParams object
// with the default values initialized.
func NewShowExportParams() *ShowExportParams {
	var ()
	return &ShowExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowExportParamsWithTimeout creates a new ShowExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowExportParamsWithTimeout(timeout time.Duration) *ShowExportParams {
	var ()
	return &ShowExportParams{

		timeout: timeout,
	}
}

/*ShowExportParams contains all the parameters to send to the API endpoint
for the show export operation typically these are written to a http.Request
*/
type ShowExportParams struct {

	/*ID
	  The URN of a ViPR/CoprHD ExportGroup

	*/
	ID string

	timeout time.Duration
}

// WithID adds the id to the show export params
func (o *ShowExportParams) WithID(ID string) *ShowExportParams {
	o.ID = ID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ShowExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
