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

// FindConfigCounterCacheSavePeriodReader is a Reader for the FindConfigCounterCacheSavePeriod structure.
type FindConfigCounterCacheSavePeriodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCounterCacheSavePeriodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCounterCacheSavePeriodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCounterCacheSavePeriodDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCounterCacheSavePeriodOK creates a FindConfigCounterCacheSavePeriodOK with default headers values
func NewFindConfigCounterCacheSavePeriodOK() *FindConfigCounterCacheSavePeriodOK {
	return &FindConfigCounterCacheSavePeriodOK{}
}

/*FindConfigCounterCacheSavePeriodOK handles this case with default header values.

Config value
*/
type FindConfigCounterCacheSavePeriodOK struct {
	Payload int64
}

func (o *FindConfigCounterCacheSavePeriodOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigCounterCacheSavePeriodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCounterCacheSavePeriodDefault creates a FindConfigCounterCacheSavePeriodDefault with default headers values
func NewFindConfigCounterCacheSavePeriodDefault(code int) *FindConfigCounterCacheSavePeriodDefault {
	return &FindConfigCounterCacheSavePeriodDefault{
		_statusCode: code,
	}
}

/*FindConfigCounterCacheSavePeriodDefault handles this case with default header values.

unexpected error
*/
type FindConfigCounterCacheSavePeriodDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config counter cache save period default response
func (o *FindConfigCounterCacheSavePeriodDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCounterCacheSavePeriodDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCounterCacheSavePeriodDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCounterCacheSavePeriodDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
