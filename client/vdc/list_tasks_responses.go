package vdc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/victorock/gocoprhd/models"
)

// ListTasksReader is a Reader for the ListTasks structure.
type ListTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ListTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewListTasksOK creates a ListTasksOK with default headers values
func NewListTasksOK() *ListTasksOK {
	return &ListTasksOK{}
}

/*ListTasksOK handles this case with default header values.

Task List
*/
type ListTasksOK struct {
	Payload *models.Tasks
}

func (o *ListTasksOK) Error() string {
	return fmt.Sprintf("[GET /vdc/tasks.json][%d] listTasksOK  %+v", 200, o.Payload)
}

func (o *ListTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tasks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksDefault creates a ListTasksDefault with default headers values
func NewListTasksDefault(code int) *ListTasksDefault {
	return &ListTasksDefault{
		_statusCode: code,
	}
}

/*ListTasksDefault handles this case with default header values.

Error Message
*/
type ListTasksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list tasks default response
func (o *ListTasksDefault) Code() int {
	return o._statusCode
}

func (o *ListTasksDefault) Error() string {
	return fmt.Sprintf("[GET /vdc/tasks.json][%d] ListTasks default  %+v", o._statusCode, o.Payload)
}

func (o *ListTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}