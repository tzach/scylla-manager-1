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

// DeleteClusterClusterIDTaskTaskTypeTaskIDReader is a Reader for the DeleteClusterClusterIDTaskTaskTypeTaskID structure.
type DeleteClusterClusterIDTaskTaskTypeTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClusterClusterIDTaskTaskTypeTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteClusterClusterIDTaskTaskTypeTaskIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteClusterClusterIDTaskTaskTypeTaskIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDOK creates a DeleteClusterClusterIDTaskTaskTypeTaskIDOK with default headers values
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDOK() *DeleteClusterClusterIDTaskTaskTypeTaskIDOK {
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDOK{}
}

/*DeleteClusterClusterIDTaskTaskTypeTaskIDOK handles this case with default header values.

Task deleted
*/
type DeleteClusterClusterIDTaskTaskTypeTaskIDOK struct {
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] deleteClusterClusterIdTaskTaskTypeTaskIdOK ", 200)
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest creates a DeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest with default headers values
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest() *DeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest {
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest{}
}

/*DeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] deleteClusterClusterIdTaskTaskTypeTaskIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDNotFound creates a DeleteClusterClusterIDTaskTaskTypeTaskIDNotFound with default headers values
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDNotFound() *DeleteClusterClusterIDTaskTaskTypeTaskIDNotFound {
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDNotFound{}
}

/*DeleteClusterClusterIDTaskTaskTypeTaskIDNotFound handles this case with default header values.

Not found
*/
type DeleteClusterClusterIDTaskTaskTypeTaskIDNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] deleteClusterClusterIdTaskTaskTypeTaskIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError creates a DeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError with default headers values
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError() *DeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError {
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError{}
}

/*DeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError handles this case with default header values.

Server error
*/
type DeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] deleteClusterClusterIdTaskTaskTypeTaskIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterClusterIDTaskTaskTypeTaskIDDefault creates a DeleteClusterClusterIDTaskTaskTypeTaskIDDefault with default headers values
func NewDeleteClusterClusterIDTaskTaskTypeTaskIDDefault(code int) *DeleteClusterClusterIDTaskTaskTypeTaskIDDefault {
	return &DeleteClusterClusterIDTaskTaskTypeTaskIDDefault{
		_statusCode: code,
	}
}

/*DeleteClusterClusterIDTaskTaskTypeTaskIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteClusterClusterIDTaskTaskTypeTaskIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete cluster cluster ID task task type task ID default response
func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}/task/{task_type}/{task_id}][%d] DeleteClusterClusterIDTaskTaskTypeTaskID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterClusterIDTaskTaskTypeTaskIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
