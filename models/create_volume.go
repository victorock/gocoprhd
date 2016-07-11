package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*CreateVolume create volume

swagger:model CreateVolume
*/
type CreateVolume struct {

	/* consistency group
	 */
	ConsistencyGroup string `json:"consistency_group,omitempty"`

	/* count
	 */
	Count int32 `json:"count,omitempty"`

	/* name
	 */
	Name string `json:"name,omitempty"`

	/* project
	 */
	Project string `json:"project,omitempty"`

	/* size
	 */
	Size int32 `json:"size,omitempty"`

	/* varray
	 */
	Varray string `json:"varray,omitempty"`

	/* vpool
	 */
	Vpool string `json:"vpool,omitempty"`
}

// Validate validates this create volume
func (m *CreateVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}