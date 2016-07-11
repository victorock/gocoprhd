package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*CreateExport create export

swagger:model CreateExport
*/
type CreateExport struct {

	/* clusters
	 */
	Clusters []string `json:"clusters,omitempty"`

	/* hosts
	 */
	Hosts []string `json:"hosts,omitempty"`

	/* initiators
	 */
	Initiators []string `json:"initiators,omitempty"`

	/* name
	 */
	Name string `json:"name,omitempty"`

	/* project
	 */
	Project string `json:"project,omitempty"`

	/* type
	 */
	Type string `json:"type,omitempty"`

	/* varray
	 */
	Varray string `json:"varray,omitempty"`

	/* volumes
	 */
	Volumes []*CreateExportVolumesItems0 `json:"volumes,omitempty"`
}

// Validate validates this create export
func (m *CreateExport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInitiators(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateExport) validateClusters(formats strfmt.Registry) error {

	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	return nil
}

func (m *CreateExport) validateHosts(formats strfmt.Registry) error {

	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	return nil
}

func (m *CreateExport) validateInitiators(formats strfmt.Registry) error {

	if swag.IsZero(m.Initiators) { // not required
		return nil
	}

	return nil
}

func (m *CreateExport) validateVolumes(formats strfmt.Registry) error {

	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {

		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {

			if err := m.Volumes[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

/*CreateExportVolumesItems0 create export volumes items0

swagger:model CreateExportVolumesItems0
*/
type CreateExportVolumesItems0 struct {

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* lun
	 */
	Lun string `json:"lun,omitempty"`
}

// Validate validates this create export volumes items0
func (m *CreateExportVolumesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}