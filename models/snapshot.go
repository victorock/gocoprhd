package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Snapshot snapshot

swagger:model Snapshot
*/
type Snapshot struct {
	Volume

	/* is sync active
	 */
	IsSyncActive bool `json:"is_sync_active,omitempty"`

	/* replica state
	 */
	ReplicaState string `json:"replica_state,omitempty"`

	/* snapset label
	 */
	SnapsetLabel string `json:"snapset_label,omitempty"`

	/* source native id
	 */
	SourceNativeID string `json:"source_native_id,omitempty"`
}

// Validate validates this snapshot
func (m *Snapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.Volume.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
