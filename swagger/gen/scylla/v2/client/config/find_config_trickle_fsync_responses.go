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

// FindConfigTrickleFsyncReader is a Reader for the FindConfigTrickleFsync structure.
type FindConfigTrickleFsyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigTrickleFsyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigTrickleFsyncOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigTrickleFsyncDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigTrickleFsyncOK creates a FindConfigTrickleFsyncOK with default headers values
func NewFindConfigTrickleFsyncOK() *FindConfigTrickleFsyncOK {
	return &FindConfigTrickleFsyncOK{}
}

/*FindConfigTrickleFsyncOK handles this case with default header values.

Config value
*/
type FindConfigTrickleFsyncOK struct {
	Payload bool
}

func (o *FindConfigTrickleFsyncOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigTrickleFsyncOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigTrickleFsyncDefault creates a FindConfigTrickleFsyncDefault with default headers values
func NewFindConfigTrickleFsyncDefault(code int) *FindConfigTrickleFsyncDefault {
	return &FindConfigTrickleFsyncDefault{
		_statusCode: code,
	}
}

/*FindConfigTrickleFsyncDefault handles this case with default header values.

unexpected error
*/
type FindConfigTrickleFsyncDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config trickle fsync default response
func (o *FindConfigTrickleFsyncDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigTrickleFsyncDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigTrickleFsyncDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigTrickleFsyncDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
