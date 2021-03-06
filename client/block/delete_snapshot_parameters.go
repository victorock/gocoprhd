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

// NewDeleteSnapshotParams creates a new DeleteSnapshotParams object
// with the default values initialized.
func NewDeleteSnapshotParams() *DeleteSnapshotParams {
	var ()
	return &DeleteSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnapshotParamsWithTimeout creates a new DeleteSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSnapshotParamsWithTimeout(timeout time.Duration) *DeleteSnapshotParams {
	var ()
	return &DeleteSnapshotParams{

		timeout: timeout,
	}
}

/*DeleteSnapshotParams contains all the parameters to send to the API endpoint
for the delete snapshot operation typically these are written to a http.Request
*/
type DeleteSnapshotParams struct {

	/*ID
	  The URN of a ViPR/CoprHD Snapshot

	*/
	ID string

	timeout time.Duration
}

// WithID adds the id to the delete snapshot params
func (o *DeleteSnapshotParams) WithID(ID string) *DeleteSnapshotParams {
	o.ID = ID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
