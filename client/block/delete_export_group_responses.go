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

// DeleteExportGroupReader is a Reader for the DeleteExportGroup structure.
type DeleteExportGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteExportGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteExportGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteExportGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteExportGroupOK creates a DeleteExportGroupOK with default headers values
func NewDeleteExportGroupOK() *DeleteExportGroupOK {
	return &DeleteExportGroupOK{}
}

/*DeleteExportGroupOK handles this case with default header values.

Task Object
*/
type DeleteExportGroupOK struct {
	Payload *models.Task
}

func (o *DeleteExportGroupOK) Error() string {
	return fmt.Sprintf("[POST /block/exports/{id}/deactivate.json][%d] deleteExportGroupOK  %+v", 200, o.Payload)
}

func (o *DeleteExportGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExportGroupDefault creates a DeleteExportGroupDefault with default headers values
func NewDeleteExportGroupDefault(code int) *DeleteExportGroupDefault {
	return &DeleteExportGroupDefault{
		_statusCode: code,
	}
}

/*DeleteExportGroupDefault handles this case with default header values.

Error Message
*/
type DeleteExportGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete export group default response
func (o *DeleteExportGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteExportGroupDefault) Error() string {
	return fmt.Sprintf("[POST /block/exports/{id}/deactivate.json][%d] DeleteExportGroup default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteExportGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
