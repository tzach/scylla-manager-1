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

// SystemLoggerGetReader is a Reader for the SystemLoggerGet structure.
type SystemLoggerGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemLoggerGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemLoggerGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSystemLoggerGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSystemLoggerGetOK creates a SystemLoggerGetOK with default headers values
func NewSystemLoggerGetOK() *SystemLoggerGetOK {
	return &SystemLoggerGetOK{}
}

/*SystemLoggerGetOK handles this case with default header values.

SystemLoggerGetOK system logger get o k
*/
type SystemLoggerGetOK struct {
	Payload []string
}

func (o *SystemLoggerGetOK) GetPayload() []string {
	return o.Payload
}

func (o *SystemLoggerGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemLoggerGetDefault creates a SystemLoggerGetDefault with default headers values
func NewSystemLoggerGetDefault(code int) *SystemLoggerGetDefault {
	return &SystemLoggerGetDefault{
		_statusCode: code,
	}
}

/*SystemLoggerGetDefault handles this case with default header values.

internal server error
*/
type SystemLoggerGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the system logger get default response
func (o *SystemLoggerGetDefault) Code() int {
	return o._statusCode
}

func (o *SystemLoggerGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *SystemLoggerGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *SystemLoggerGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
