package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Initiator initiator

swagger:model Initiator
*/
type Initiator struct {
	Resource

	/* creation time
	 */
	CreationTime int64 `json:"creation_time,omitempty"`

	/* host
	 */
	Host *Resource `json:"host,omitempty"`

	/* hostname
	 */
	Hostname string `json:"hostname,omitempty"`

	/* inactive
	 */
	Inactive bool `json:"inactive,omitempty"`

	/* initiator node
	 */
	InitiatorNode string `json:"initiator_node,omitempty"`

	/* initiator port
	 */
	InitiatorPort string `json:"initiator_port,omitempty"`

	/* protocol
	 */
	Protocol string `json:"protocol,omitempty"`

	/* tags
	 */
	Tags []string `json:"tags,omitempty"`
}

// Validate validates this initiator
func (m *Initiator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Initiator) validateHost(formats strfmt.Registry) error {

	if m.Host != nil {

		if err := m.Host.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var initiatorTypeProtocolPropEnum []interface{}

// property enum
func (m *Initiator) validateProtocolEnum(path, location string, value string) error {
	if initiatorTypeProtocolPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["FC","iSCSI","ScaleIO"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			initiatorTypeProtocolPropEnum = append(initiatorTypeProtocolPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, initiatorTypeProtocolPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Initiator) validateProtocol(formats strfmt.Registry) error {

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *Initiator) validateTags(formats strfmt.Registry) error {

	return nil
}
