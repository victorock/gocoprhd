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

// ListSnapshotsReader is a Reader for the ListSnapshots structure.
type ListSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ListSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListSnapshotsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewListSnapshotsOK creates a ListSnapshotsOK with default headers values
func NewListSnapshotsOK() *ListSnapshotsOK {
	return &ListSnapshotsOK{}
}

/*ListSnapshotsOK handles this case with default header values.

Snapshots List
*/
type ListSnapshotsOK struct {
	Payload *models.Snapshots
}

func (o *ListSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /block/snapshots/bulk.json][%d] listSnapshotsOK  %+v", 200, o.Payload)
}

func (o *ListSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snapshots)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSnapshotsDefault creates a ListSnapshotsDefault with default headers values
func NewListSnapshotsDefault(code int) *ListSnapshotsDefault {
	return &ListSnapshotsDefault{
		_statusCode: code,
	}
}

/*ListSnapshotsDefault handles this case with default header values.

Error Message
*/
type ListSnapshotsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list snapshots default response
func (o *ListSnapshotsDefault) Code() int {
	return o._statusCode
}

func (o *ListSnapshotsDefault) Error() string {
	return fmt.Sprintf("[GET /block/snapshots/bulk.json][%d] ListSnapshots default  %+v", o._statusCode, o.Payload)
}

func (o *ListSnapshotsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
