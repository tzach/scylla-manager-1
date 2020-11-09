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

// FindConfigConsistentRangemovementReader is a Reader for the FindConfigConsistentRangemovement structure.
type FindConfigConsistentRangemovementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigConsistentRangemovementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigConsistentRangemovementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigConsistentRangemovementDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigConsistentRangemovementOK creates a FindConfigConsistentRangemovementOK with default headers values
func NewFindConfigConsistentRangemovementOK() *FindConfigConsistentRangemovementOK {
	return &FindConfigConsistentRangemovementOK{}
}

/*FindConfigConsistentRangemovementOK handles this case with default header values.

Config value
*/
type FindConfigConsistentRangemovementOK struct {
	Payload bool
}

func (o *FindConfigConsistentRangemovementOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigConsistentRangemovementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigConsistentRangemovementDefault creates a FindConfigConsistentRangemovementDefault with default headers values
func NewFindConfigConsistentRangemovementDefault(code int) *FindConfigConsistentRangemovementDefault {
	return &FindConfigConsistentRangemovementDefault{
		_statusCode: code,
	}
}

/*FindConfigConsistentRangemovementDefault handles this case with default header values.

unexpected error
*/
type FindConfigConsistentRangemovementDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config consistent rangemovement default response
func (o *FindConfigConsistentRangemovementDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigConsistentRangemovementDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigConsistentRangemovementDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigConsistentRangemovementDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
