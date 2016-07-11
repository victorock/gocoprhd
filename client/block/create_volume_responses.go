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

// CreateVolumeReader is a Reader for the CreateVolume structure.
type CreateVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateVolumeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewCreateVolumeOK creates a CreateVolumeOK with default headers values
func NewCreateVolumeOK() *CreateVolumeOK {
	return &CreateVolumeOK{}
}

/*CreateVolumeOK handles this case with default header values.

Task Object
*/
type CreateVolumeOK struct {
	Payload *models.Task
}

func (o *CreateVolumeOK) Error() string {
	return fmt.Sprintf("[POST /block/volumes.json][%d] createVolumeOK  %+v", 200, o.Payload)
}

func (o *CreateVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVolumeDefault creates a CreateVolumeDefault with default headers values
func NewCreateVolumeDefault(code int) *CreateVolumeDefault {
	return &CreateVolumeDefault{
		_statusCode: code,
	}
}

/*CreateVolumeDefault handles this case with default header values.

Error Message
*/
type CreateVolumeDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create volume default response
func (o *CreateVolumeDefault) Code() int {
	return o._statusCode
}

func (o *CreateVolumeDefault) Error() string {
	return fmt.Sprintf("[POST /block/volumes.json][%d] CreateVolume default  %+v", o._statusCode, o.Payload)
}

func (o *CreateVolumeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
