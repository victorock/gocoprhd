package put

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/victorock/gocoprhd/models"
)

// UpdateExportReader is a Reader for the UpdateExport structure.
type UpdateExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateExportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewUpdateExportOK creates a UpdateExportOK with default headers values
func NewUpdateExportOK() *UpdateExportOK {
	return &UpdateExportOK{}
}

/*UpdateExportOK handles this case with default header values.

Task Object
*/
type UpdateExportOK struct {
	Payload *models.Task
}

func (o *UpdateExportOK) Error() string {
	return fmt.Sprintf("[PUT /block/exports/{id}.json][%d] updateExportOK  %+v", 200, o.Payload)
}

func (o *UpdateExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateExportDefault creates a UpdateExportDefault with default headers values
func NewUpdateExportDefault(code int) *UpdateExportDefault {
	return &UpdateExportDefault{
		_statusCode: code,
	}
}

/*UpdateExportDefault handles this case with default header values.

Error Message
*/
type UpdateExportDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update export default response
func (o *UpdateExportDefault) Code() int {
	return o._statusCode
}

func (o *UpdateExportDefault) Error() string {
	return fmt.Sprintf("[PUT /block/exports/{id}.json][%d] UpdateExport default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateExportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
