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

// NewListVolumesParams creates a new ListVolumesParams object
// with the default values initialized.
func NewListVolumesParams() *ListVolumesParams {

	return &ListVolumesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListVolumesParamsWithTimeout creates a new ListVolumesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListVolumesParamsWithTimeout(timeout time.Duration) *ListVolumesParams {

	return &ListVolumesParams{

		timeout: timeout,
	}
}

/*ListVolumesParams contains all the parameters to send to the API endpoint
for the list volumes operation typically these are written to a http.Request
*/
type ListVolumesParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *ListVolumesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}