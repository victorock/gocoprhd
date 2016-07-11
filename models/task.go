package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Task task

swagger:model Task
*/
type Task struct {

	/* description
	 */
	Description string `json:"description,omitempty"`

	/* link
	 */
	Link *Link `json:"link,omitempty"`

	/* message
	 */
	Message string `json:"message,omitempty"`

	/* op id
	 */
	OpID string `json:"op_id,omitempty"`

	/* resource
	 */
	Resource *Resource `json:"resource,omitempty"`

	/* start time
	 */
	StartTime string `json:"start_time,omitempty"`

	/* state
	 */
	State string `json:"state,omitempty"`
}

// Validate validates this task
func (m *Task) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Task) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {

		if err := m.Link.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Task) validateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {

		if err := m.Resource.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
