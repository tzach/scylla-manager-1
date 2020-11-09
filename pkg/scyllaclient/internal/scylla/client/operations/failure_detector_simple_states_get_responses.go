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

// FailureDetectorSimpleStatesGetReader is a Reader for the FailureDetectorSimpleStatesGet structure.
type FailureDetectorSimpleStatesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FailureDetectorSimpleStatesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFailureDetectorSimpleStatesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFailureDetectorSimpleStatesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFailureDetectorSimpleStatesGetOK creates a FailureDetectorSimpleStatesGetOK with default headers values
func NewFailureDetectorSimpleStatesGetOK() *FailureDetectorSimpleStatesGetOK {
	return &FailureDetectorSimpleStatesGetOK{}
}

/*FailureDetectorSimpleStatesGetOK handles this case with default header values.

FailureDetectorSimpleStatesGetOK failure detector simple states get o k
*/
type FailureDetectorSimpleStatesGetOK struct {
	Payload []*models.Mapper
}

func (o *FailureDetectorSimpleStatesGetOK) GetPayload() []*models.Mapper {
	return o.Payload
}

func (o *FailureDetectorSimpleStatesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFailureDetectorSimpleStatesGetDefault creates a FailureDetectorSimpleStatesGetDefault with default headers values
func NewFailureDetectorSimpleStatesGetDefault(code int) *FailureDetectorSimpleStatesGetDefault {
	return &FailureDetectorSimpleStatesGetDefault{
		_statusCode: code,
	}
}

/*FailureDetectorSimpleStatesGetDefault handles this case with default header values.

internal server error
*/
type FailureDetectorSimpleStatesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the failure detector simple states get default response
func (o *FailureDetectorSimpleStatesGetDefault) Code() int {
	return o._statusCode
}

func (o *FailureDetectorSimpleStatesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FailureDetectorSimpleStatesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FailureDetectorSimpleStatesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
