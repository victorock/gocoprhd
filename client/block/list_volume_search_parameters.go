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

	/*Name
	  Name to match

	*/
	Name *string
	/*Project
	  Project URN to match

	*/
	Project *string
	/*Tag
	  Tag to match

	*/
	Tag *string
	/*Wwn
	  wwn to match

	*/
	Wwn *string

	timeout time.Duration
}

// WithName adds the name to the list volume search params
func (o *ListVolumeSearchParams) WithName(Name *string) *ListVolumeSearchParams {
	o.Name = Name
	return o
}

// WithProject adds the project to the list volume search params
func (o *ListVolumeSearchParams) WithProject(Project *string) *ListVolumeSearchParams {
	o.Project = Project
	return o
}

// WithTag adds the tag to the list volume search params
func (o *ListVolumeSearchParams) WithTag(Tag *string) *ListVolumeSearchParams {
	o.Tag = Tag
	return o
}

// WithWwn adds the wwn to the list volume search params
func (o *ListVolumeSearchParams) WithWwn(Wwn *string) *ListVolumeSearchParams {
	o.Wwn = Wwn
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListVolumeSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Project != nil {

		// query param project
		var qrProject string
		if o.Project != nil {
			qrProject = *o.Project
		}
		qProject := qrProject
		if qProject != "" {
			if err := r.SetQueryParam("project", qProject); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if o.Wwn != nil {

		// query param wwn
		var qrWwn string
		if o.Wwn != nil {
			qrWwn = *o.Wwn
		}
		qWwn := qrWwn
		if qWwn != "" {
			if err := r.SetQueryParam("wwn", qWwn); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
