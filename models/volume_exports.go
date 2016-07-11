package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*VolumeExports volume exports

swagger:model VolumeExports
*/
type VolumeExports struct {

	/* itl
	 */
	Itl []*VolumeExportsItlItems0 `json:"itl,omitempty"`
}

// Validate validates this volume exports
func (m *VolumeExports) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItl(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeExports) validateItl(formats strfmt.Registry) error {

	if swag.IsZero(m.Itl) { // not required
		return nil
	}

	for i := 0; i < len(m.Itl); i++ {

		if swag.IsZero(m.Itl[i]) { // not required
			continue
		}

		if m.Itl[i] != nil {

			if err := m.Itl[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

/*VolumeExportsItlItems0 volume exports itl items0

swagger:model VolumeExportsItlItems0
*/
type VolumeExportsItlItems0 struct {

	/* device
	 */
	Device *VolumeExportsItlItems0Device `json:"device,omitempty"`

	/* export
	 */
	Export *VolumeExportsItlItems0Export `json:"export,omitempty"`

	/* hlu
	 */
	Hlu int64 `json:"hlu,omitempty"`

	/* initiator
	 */
	Initiator *VolumeExportsItlItems0Initiator `json:"initiator,omitempty"`

	/* san zone name
	 */
	SanZoneName string `json:"san_zone_name,omitempty"`

	/* target
	 */
	Target *VolumeExportsItlItems0Target `json:"target,omitempty"`
}

// Validate validates this volume exports itl items0
func (m *VolumeExportsItlItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExport(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInitiator(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeExportsItlItems0) validateDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {

		if err := m.Device.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VolumeExportsItlItems0) validateExport(formats strfmt.Registry) error {

	if swag.IsZero(m.Export) { // not required
		return nil
	}

	if m.Export != nil {

		if err := m.Export.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VolumeExportsItlItems0) validateInitiator(formats strfmt.Registry) error {

	if swag.IsZero(m.Initiator) { // not required
		return nil
	}

	if m.Initiator != nil {

		if err := m.Initiator.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *VolumeExportsItlItems0) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {

		if err := m.Target.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

/*VolumeExportsItlItems0Device volume exports itl items0 device

swagger:model VolumeExportsItlItems0Device
*/
type VolumeExportsItlItems0Device struct {

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* link
	 */
	Link *Link `json:"link,omitempty"`

	/* wwn
	 */
	Wwn string `json:"wwn,omitempty"`
}

// Validate validates this volume exports itl items0 device
func (m *VolumeExportsItlItems0Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeExportsItlItems0Device) validateLink(formats strfmt.Registry) error {

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

/*VolumeExportsItlItems0Export volume exports itl items0 export

swagger:model VolumeExportsItlItems0Export
*/
type VolumeExportsItlItems0Export struct {

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* link
	 */
	Link *Link `json:"link,omitempty"`

	/* name
	 */
	Name string `json:"name,omitempty"`
}

// Validate validates this volume exports itl items0 export
func (m *VolumeExportsItlItems0Export) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeExportsItlItems0Export) validateLink(formats strfmt.Registry) error {

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

/*VolumeExportsItlItems0Initiator volume exports itl items0 initiator

swagger:model VolumeExportsItlItems0Initiator
*/
type VolumeExportsItlItems0Initiator struct {

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* link
	 */
	Link *Link `json:"link,omitempty"`

	/* port
	 */
	Port string `json:"port,omitempty"`
}

// Validate validates this volume exports itl items0 initiator
func (m *VolumeExportsItlItems0Initiator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeExportsItlItems0Initiator) validateLink(formats strfmt.Registry) error {

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

/*VolumeExportsItlItems0Target volume exports itl items0 target

swagger:model VolumeExportsItlItems0Target
*/
type VolumeExportsItlItems0Target struct {

	/* id
	 */
	ID string `json:"id,omitempty"`

	/* link
	 */
	Link *Link `json:"link,omitempty"`

	/* port
	 */
	Port string `json:"port,omitempty"`
}

// Validate validates this volume exports itl items0 target
func (m *VolumeExportsItlItems0Target) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeExportsItlItems0Target) validateLink(formats strfmt.Registry) error {

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