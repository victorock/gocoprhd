package block

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/victorock/gocoprhd/models"
)

// DeleteSnapshotReader is a Reader for the DeleteSnapshot structure.
type DeleteSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteSnapshotAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteSnapshotAccepted creates a DeleteSnapshotAccepted with default headers values
func NewDeleteSnapshotAccepted() *DeleteSnapshotAccepted {
	return &DeleteSnapshotAccepted{}
}

/*DeleteSnapshotAccepted handles this case with default header values.

Task Object
*/
type DeleteSnapshotAccepted struct {
	Payload *models.Tasks
}

func (o *DeleteSnapshotAccepted) Error() string {
	return fmt.Sprintf("[POST /block/snapshots/{id}/deactivate.json][%d] deleteSnapshotAccepted  %+v", 202, o.Payload)
}

func (o *DeleteSnapshotAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tasks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSnapshotDefault creates a DeleteSnapshotDefault with default headers values
func NewDeleteSnapshotDefault(code int) *DeleteSnapshotDefault {
	return &DeleteSnapshotDefault{
		_statusCode: code,
	}
}

/*DeleteSnapshotDefault handles this case with default header values.

Error Message
*/
type DeleteSnapshotDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete snapshot default response
func (o *DeleteSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSnapshotDefault) Error() string {
	return fmt.Sprintf("[POST /block/snapshots/{id}/deactivate.json][%d] DeleteSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
