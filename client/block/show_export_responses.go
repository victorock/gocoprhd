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

// ShowExportReader is a Reader for the ShowExport structure.
type ShowExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ShowExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShowExportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewShowExportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewShowExportOK creates a ShowExportOK with default headers values
func NewShowExportOK() *ShowExportOK {
	return &ShowExportOK{}
}

/*ShowExportOK handles this case with default header values.

Export Object
*/
type ShowExportOK struct {
	Payload *models.Export
}

func (o *ShowExportOK) Error() string {
	return fmt.Sprintf("[GET /block/exports/{id}.json][%d] showExportOK  %+v", 200, o.Payload)
}

func (o *ShowExportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Export)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowExportDefault creates a ShowExportDefault with default headers values
func NewShowExportDefault(code int) *ShowExportDefault {
	return &ShowExportDefault{
		_statusCode: code,
	}
}

/*ShowExportDefault handles this case with default header values.

Error Message
*/
type ShowExportDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the show export default response
func (o *ShowExportDefault) Code() int {
	return o._statusCode
}

func (o *ShowExportDefault) Error() string {
	return fmt.Sprintf("[GET /block/exports/{id}.json][%d] ShowExport default  %+v", o._statusCode, o.Payload)
}

func (o *ShowExportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}