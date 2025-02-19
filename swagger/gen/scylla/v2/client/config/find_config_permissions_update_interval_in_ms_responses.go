// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v2/models"
)

// FindConfigPermissionsUpdateIntervalInMsReader is a Reader for the FindConfigPermissionsUpdateIntervalInMs structure.
type FindConfigPermissionsUpdateIntervalInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigPermissionsUpdateIntervalInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigPermissionsUpdateIntervalInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigPermissionsUpdateIntervalInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigPermissionsUpdateIntervalInMsOK creates a FindConfigPermissionsUpdateIntervalInMsOK with default headers values
func NewFindConfigPermissionsUpdateIntervalInMsOK() *FindConfigPermissionsUpdateIntervalInMsOK {
	return &FindConfigPermissionsUpdateIntervalInMsOK{}
}

/*FindConfigPermissionsUpdateIntervalInMsOK handles this case with default header values.

Config value
*/
type FindConfigPermissionsUpdateIntervalInMsOK struct {
	Payload int64
}

func (o *FindConfigPermissionsUpdateIntervalInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigPermissionsUpdateIntervalInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigPermissionsUpdateIntervalInMsDefault creates a FindConfigPermissionsUpdateIntervalInMsDefault with default headers values
func NewFindConfigPermissionsUpdateIntervalInMsDefault(code int) *FindConfigPermissionsUpdateIntervalInMsDefault {
	return &FindConfigPermissionsUpdateIntervalInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigPermissionsUpdateIntervalInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigPermissionsUpdateIntervalInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config permissions update interval in ms default response
func (o *FindConfigPermissionsUpdateIntervalInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigPermissionsUpdateIntervalInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigPermissionsUpdateIntervalInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigPermissionsUpdateIntervalInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
