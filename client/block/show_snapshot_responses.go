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

// ShowSnapshotReader is a Reader for the ShowSnapshot structure.
type ShowSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ShowSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShowSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewShowSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewShowSnapshotOK creates a ShowSnapshotOK with default headers values
func NewShowSnapshotOK() *ShowSnapshotOK {
	return &ShowSnapshotOK{}
}

/*ShowSnapshotOK handles this case with default header values.

Snapshots Details
*/
type ShowSnapshotOK struct {
	Payload *models.Snapshot
}

func (o *ShowSnapshotOK) Error() string {
	return fmt.Sprintf("[GET /block/snapshots/{id}.json][%d] showSnapshotOK  %+v", 200, o.Payload)
}

func (o *ShowSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowSnapshotDefault creates a ShowSnapshotDefault with default headers values
func NewShowSnapshotDefault(code int) *ShowSnapshotDefault {
	return &ShowSnapshotDefault{
		_statusCode: code,
	}
}

/*ShowSnapshotDefault handles this case with default header values.

Error Message
*/
type ShowSnapshotDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the show snapshot default response
func (o *ShowSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *ShowSnapshotDefault) Error() string {
	return fmt.Sprintf("[GET /block/snapshots/{id}.json][%d] ShowSnapshot default  %+v", o._statusCode, o.Payload)
}

func (o *ShowSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}