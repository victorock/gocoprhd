package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*SearchResources search resources

swagger:model SearchResources
*/
type SearchResources struct {

	/* resource
	 */
	Resource []*SearchResourcesResourceItems0 `json:"resource,omitempty"`
}

// Validate validates this search resources
func (m *SearchResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchResources) validateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	for i := 0; i < len(m.Resource); i++ {

		if swag.IsZero(m.Resource[i]) { // not required
			continue
		}

		if m.Resource[i] != nil {

			if err := m.Resource[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

/*SearchResourcesResourceItems0 search resources resource items0

swagger:model SearchResourcesResourceItems0
*/
type SearchResourcesResourceItems0 struct {
	Resource

	/* match
	 */
	Match string `json:"match,omitempty"`
}

// Validate validates this search resources resource items0
func (m *SearchResourcesResourceItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
