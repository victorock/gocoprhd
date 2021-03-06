package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*UpdateExport update export

swagger:model UpdateExport
*/
type UpdateExport struct {

	/* cluster changes
	 */
	ClusterChanges *UpdateExportClusterChanges `json:"cluster_changes,omitempty"`

	/* host changes
	 */
	HostChanges *UpdateExportHostChanges `json:"host_changes,omitempty"`

	/* initiator changes
	 */
	InitiatorChanges *UpdateExportInitiatorChanges `json:"initiator_changes,omitempty"`

	/* volume changes
	 */
	VolumeChanges string `json:"volume_changes,omitempty"`
}

// Validate validates this update export
func (m *UpdateExport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterChanges(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHostChanges(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInitiatorChanges(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateExport) validateClusterChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterChanges) { // not required
		return nil
	}

	if m.ClusterChanges != nil {

		if err := m.ClusterChanges.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *UpdateExport) validateHostChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.HostChanges) { // not required
		return nil
	}

	if m.HostChanges != nil {

		if err := m.HostChanges.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *UpdateExport) validateInitiatorChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.InitiatorChanges) { // not required
		return nil
	}

	if m.InitiatorChanges != nil {

		if err := m.InitiatorChanges.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

/*UpdateExportClusterChanges update export cluster changes

swagger:model UpdateExportClusterChanges
*/
type UpdateExportClusterChanges struct {

	/* add
	 */
	Add []string `json:"add,omitempty"`

	/* remove
	 */
	Remove []string `json:"remove,omitempty"`
}

// Validate validates this update export cluster changes
func (m *UpdateExportClusterChanges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRemove(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateExportClusterChanges) validateAdd(formats strfmt.Registry) error {

	if swag.IsZero(m.Add) { // not required
		return nil
	}

	return nil
}

func (m *UpdateExportClusterChanges) validateRemove(formats strfmt.Registry) error {

	if swag.IsZero(m.Remove) { // not required
		return nil
	}

	return nil
}

/*UpdateExportHostChanges update export host changes

swagger:model UpdateExportHostChanges
*/
type UpdateExportHostChanges struct {

	/* add
	 */
	Add []string `json:"add,omitempty"`

	/* remove
	 */
	Remove []string `json:"remove,omitempty"`
}

// Validate validates this update export host changes
func (m *UpdateExportHostChanges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRemove(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateExportHostChanges) validateAdd(formats strfmt.Registry) error {

	if swag.IsZero(m.Add) { // not required
		return nil
	}

	return nil
}

func (m *UpdateExportHostChanges) validateRemove(formats strfmt.Registry) error {

	if swag.IsZero(m.Remove) { // not required
		return nil
	}

	return nil
}

/*UpdateExportInitiatorChanges update export initiator changes

swagger:model UpdateExportInitiatorChanges
*/
type UpdateExportInitiatorChanges struct {

	/* add
	 */
	Add []string `json:"add,omitempty"`

	/* remove
	 */
	Remove []string `json:"remove,omitempty"`
}

// Validate validates this update export initiator changes
func (m *UpdateExportInitiatorChanges) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRemove(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateExportInitiatorChanges) validateAdd(formats strfmt.Registry) error {

	if swag.IsZero(m.Add) { // not required
		return nil
	}

	return nil
}

func (m *UpdateExportInitiatorChanges) validateRemove(formats strfmt.Registry) error {

	if swag.IsZero(m.Remove) { // not required
		return nil
	}

	return nil
}
