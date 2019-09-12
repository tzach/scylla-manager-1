// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/mermaid/mermaidclient/internal/models"
)

// PutClusterClusterIDTaskTaskTypeTaskIDStopReader is a Reader for the PutClusterClusterIDTaskTaskTypeTaskIDStop structure.
type PutClusterClusterIDTaskTaskTypeTaskIDStopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutClusterClusterIDTaskTaskTypeTaskIDStopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutClusterClusterIDTaskTaskTypeTaskIDStopNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutClusterClusterIDTaskTaskTypeTaskIDStopDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStopOK creates a PutClusterClusterIDTaskTaskTypeTaskIDStopOK with default headers values
func NewPutClusterClusterIDTaskTaskTypeTaskIDStopOK() *PutClusterClusterIDTaskTaskTypeTaskIDStopOK {
	return &PutClusterClusterIDTaskTaskTypeTaskIDStopOK{}
}

/*PutClusterClusterIDTaskTaskTypeTaskIDStopOK handles this case with default header values.

Task stopped
*/
type PutClusterClusterIDTaskTaskTypeTaskIDStopOK struct {
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopOK) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/task/{task_type}/{task_id}/stop][%d] putClusterClusterIdTaskTaskTypeTaskIdStopOK ", 200)
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest creates a PutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest with default headers values
func NewPutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest() *PutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest {
	return &PutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest{}
}

/*PutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest handles this case with default header values.

Bad Request
*/
type PutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/task/{task_type}/{task_id}/stop][%d] putClusterClusterIdTaskTaskTypeTaskIdStopBadRequest  %+v", 400, o.Payload)
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStopNotFound creates a PutClusterClusterIDTaskTaskTypeTaskIDStopNotFound with default headers values
func NewPutClusterClusterIDTaskTaskTypeTaskIDStopNotFound() *PutClusterClusterIDTaskTaskTypeTaskIDStopNotFound {
	return &PutClusterClusterIDTaskTaskTypeTaskIDStopNotFound{}
}

/*PutClusterClusterIDTaskTaskTypeTaskIDStopNotFound handles this case with default header values.

Not found
*/
type PutClusterClusterIDTaskTaskTypeTaskIDStopNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopNotFound) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/task/{task_type}/{task_id}/stop][%d] putClusterClusterIdTaskTaskTypeTaskIdStopNotFound  %+v", 404, o.Payload)
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError creates a PutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError with default headers values
func NewPutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError() *PutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError {
	return &PutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError{}
}

/*PutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError handles this case with default header values.

Server error
*/
type PutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/task/{task_type}/{task_id}/stop][%d] putClusterClusterIdTaskTaskTypeTaskIdStopInternalServerError  %+v", 500, o.Payload)
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutClusterClusterIDTaskTaskTypeTaskIDStopDefault creates a PutClusterClusterIDTaskTaskTypeTaskIDStopDefault with default headers values
func NewPutClusterClusterIDTaskTaskTypeTaskIDStopDefault(code int) *PutClusterClusterIDTaskTaskTypeTaskIDStopDefault {
	return &PutClusterClusterIDTaskTaskTypeTaskIDStopDefault{
		_statusCode: code,
	}
}

/*PutClusterClusterIDTaskTaskTypeTaskIDStopDefault handles this case with default header values.

Unexpected error
*/
type PutClusterClusterIDTaskTaskTypeTaskIDStopDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the put cluster cluster ID task task type task ID stop default response
func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopDefault) Code() int {
	return o._statusCode
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopDefault) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/task/{task_type}/{task_id}/stop][%d] PutClusterClusterIDTaskTaskTypeTaskIDStop default  %+v", o._statusCode, o.Payload)
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PutClusterClusterIDTaskTaskTypeTaskIDStopDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
