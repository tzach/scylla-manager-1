// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/agent/models"
)

// OperationsDeletefileReader is a Reader for the OperationsDeletefile structure.
type OperationsDeletefileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperationsDeletefileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperationsDeletefileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOperationsDeletefileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOperationsDeletefileOK creates a OperationsDeletefileOK with default headers values
func NewOperationsDeletefileOK() *OperationsDeletefileOK {
	return &OperationsDeletefileOK{}
}

/*OperationsDeletefileOK handles this case with default header values.

Job ID
*/
type OperationsDeletefileOK struct {
	Payload *models.Jobid
	JobID   int64
}

func (o *OperationsDeletefileOK) GetPayload() *models.Jobid {
	return o.Payload
}

func (o *OperationsDeletefileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Jobid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

// NewOperationsDeletefileDefault creates a OperationsDeletefileDefault with default headers values
func NewOperationsDeletefileDefault(code int) *OperationsDeletefileDefault {
	return &OperationsDeletefileDefault{
		_statusCode: code,
	}
}

/*OperationsDeletefileDefault handles this case with default header values.

Server error
*/
type OperationsDeletefileDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
	JobID   int64
}

// Code gets the status code for the operations deletefile default response
func (o *OperationsDeletefileDefault) Code() int {
	return o._statusCode
}

func (o *OperationsDeletefileDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *OperationsDeletefileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	if jobIDHeader := response.GetHeader("x-rclone-jobid"); jobIDHeader != "" {
		jobID, err := strconv.ParseInt(jobIDHeader, 10, 64)
		if err != nil {
			return err
		}

		o.JobID = jobID
	}
	return nil
}

func (o *OperationsDeletefileDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
