// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigSeedProviderReader is a Reader for the FindConfigSeedProvider structure.
type FindConfigSeedProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigSeedProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigSeedProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigSeedProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigSeedProviderOK creates a FindConfigSeedProviderOK with default headers values
func NewFindConfigSeedProviderOK() *FindConfigSeedProviderOK {
	return &FindConfigSeedProviderOK{}
}

/*FindConfigSeedProviderOK handles this case with default header values.

Config value
*/
type FindConfigSeedProviderOK struct {
	Payload []string
}

func (o *FindConfigSeedProviderOK) GetPayload() []string {
	return o.Payload
}

func (o *FindConfigSeedProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigSeedProviderDefault creates a FindConfigSeedProviderDefault with default headers values
func NewFindConfigSeedProviderDefault(code int) *FindConfigSeedProviderDefault {
	return &FindConfigSeedProviderDefault{
		_statusCode: code,
	}
}

/*FindConfigSeedProviderDefault handles this case with default header values.

unexpected error
*/
type FindConfigSeedProviderDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config seed provider default response
func (o *FindConfigSeedProviderDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigSeedProviderDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigSeedProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigSeedProviderDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
