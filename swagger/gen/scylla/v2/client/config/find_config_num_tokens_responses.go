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

// FindConfigNumTokensReader is a Reader for the FindConfigNumTokens structure.
type FindConfigNumTokensReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigNumTokensReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigNumTokensOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigNumTokensDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigNumTokensOK creates a FindConfigNumTokensOK with default headers values
func NewFindConfigNumTokensOK() *FindConfigNumTokensOK {
	return &FindConfigNumTokensOK{}
}

/*FindConfigNumTokensOK handles this case with default header values.

Config value
*/
type FindConfigNumTokensOK struct {
	Payload int64
}

func (o *FindConfigNumTokensOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigNumTokensOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigNumTokensDefault creates a FindConfigNumTokensDefault with default headers values
func NewFindConfigNumTokensDefault(code int) *FindConfigNumTokensDefault {
	return &FindConfigNumTokensDefault{
		_statusCode: code,
	}
}

/*FindConfigNumTokensDefault handles this case with default header values.

unexpected error
*/
type FindConfigNumTokensDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config num tokens default response
func (o *FindConfigNumTokensDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigNumTokensDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigNumTokensDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigNumTokensDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
