// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla/models"
)

// CommitlogSegmentsActiveGetReader is a Reader for the CommitlogSegmentsActiveGet structure.
type CommitlogSegmentsActiveGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitlogSegmentsActiveGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitlogSegmentsActiveGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCommitlogSegmentsActiveGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCommitlogSegmentsActiveGetOK creates a CommitlogSegmentsActiveGetOK with default headers values
func NewCommitlogSegmentsActiveGetOK() *CommitlogSegmentsActiveGetOK {
	return &CommitlogSegmentsActiveGetOK{}
}

/*CommitlogSegmentsActiveGetOK handles this case with default header values.

CommitlogSegmentsActiveGetOK commitlog segments active get o k
*/
type CommitlogSegmentsActiveGetOK struct {
	Payload []string
}

func (o *CommitlogSegmentsActiveGetOK) GetPayload() []string {
	return o.Payload
}

func (o *CommitlogSegmentsActiveGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCommitlogSegmentsActiveGetDefault creates a CommitlogSegmentsActiveGetDefault with default headers values
func NewCommitlogSegmentsActiveGetDefault(code int) *CommitlogSegmentsActiveGetDefault {
	return &CommitlogSegmentsActiveGetDefault{
		_statusCode: code,
	}
}

/*CommitlogSegmentsActiveGetDefault handles this case with default header values.

internal server error
*/
type CommitlogSegmentsActiveGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the commitlog segments active get default response
func (o *CommitlogSegmentsActiveGetDefault) Code() int {
	return o._statusCode
}

func (o *CommitlogSegmentsActiveGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CommitlogSegmentsActiveGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CommitlogSegmentsActiveGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
