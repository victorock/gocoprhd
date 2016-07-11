package snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSnapshotsParams creates a new ListSnapshotsParams object
// with the default values initialized.
func NewListSnapshotsParams() *ListSnapshotsParams {

	return &ListSnapshotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSnapshotsParamsWithTimeout creates a new ListSnapshotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSnapshotsParamsWithTimeout(timeout time.Duration) *ListSnapshotsParams {

	return &ListSnapshotsParams{

		timeout: timeout,
	}
}

/*ListSnapshotsParams contains all the parameters to send to the API endpoint
for the list snapshots operation typically these are written to a http.Request
*/
type ListSnapshotsParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *ListSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
