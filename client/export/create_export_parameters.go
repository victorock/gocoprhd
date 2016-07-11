package export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/victorock/gocoprhd/models"
)

// NewCreateExportParams creates a new CreateExportParams object
// with the default values initialized.
func NewCreateExportParams() *CreateExportParams {
	var ()
	return &CreateExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateExportParamsWithTimeout creates a new CreateExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateExportParamsWithTimeout(timeout time.Duration) *CreateExportParams {
	var ()
	return &CreateExportParams{

		timeout: timeout,
	}
}

/*CreateExportParams contains all the parameters to send to the API endpoint
for the create export operation typically these are written to a http.Request
*/
type CreateExportParams struct {

	/*Body*/
	Body *models.CreateExport

	timeout time.Duration
}

// WithBody adds the body to the create export params
func (o *CreateExportParams) WithBody(Body *models.CreateExport) *CreateExportParams {
	o.Body = Body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.CreateExport)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}