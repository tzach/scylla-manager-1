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

// FindConfigCommitlogSyncPeriodInMsReader is a Reader for the FindConfigCommitlogSyncPeriodInMs structure.
type FindConfigCommitlogSyncPeriodInMsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCommitlogSyncPeriodInMsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCommitlogSyncPeriodInMsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCommitlogSyncPeriodInMsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCommitlogSyncPeriodInMsOK creates a FindConfigCommitlogSyncPeriodInMsOK with default headers values
func NewFindConfigCommitlogSyncPeriodInMsOK() *FindConfigCommitlogSyncPeriodInMsOK {
	return &FindConfigCommitlogSyncPeriodInMsOK{}
}

/*FindConfigCommitlogSyncPeriodInMsOK handles this case with default header values.

Config value
*/
type FindConfigCommitlogSyncPeriodInMsOK struct {
	Payload int64
}

func (o *FindConfigCommitlogSyncPeriodInMsOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigCommitlogSyncPeriodInMsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCommitlogSyncPeriodInMsDefault creates a FindConfigCommitlogSyncPeriodInMsDefault with default headers values
func NewFindConfigCommitlogSyncPeriodInMsDefault(code int) *FindConfigCommitlogSyncPeriodInMsDefault {
	return &FindConfigCommitlogSyncPeriodInMsDefault{
		_statusCode: code,
	}
}

/*FindConfigCommitlogSyncPeriodInMsDefault handles this case with default header values.

unexpected error
*/
type FindConfigCommitlogSyncPeriodInMsDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config commitlog sync period in ms default response
func (o *FindConfigCommitlogSyncPeriodInMsDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCommitlogSyncPeriodInMsDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCommitlogSyncPeriodInMsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCommitlogSyncPeriodInMsDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
