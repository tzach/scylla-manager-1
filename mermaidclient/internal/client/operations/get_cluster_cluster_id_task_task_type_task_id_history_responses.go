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

// GetClusterClusterIDTaskTaskTypeTaskIDHistoryReader is a Reader for the GetClusterClusterIDTaskTaskTypeTaskIDHistory structure.
type GetClusterClusterIDTaskTaskTypeTaskIDHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryOK creates a GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK with default headers values
func NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryOK() *GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK {
	return &GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK{}
}

/*GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK handles this case with default header values.

List of task runs
*/
type GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK struct {
	Payload []*models.TaskRun
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/{task_type}/{task_id}/history][%d] getClusterClusterIdTaskTaskTypeTaskIdHistoryOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK) GetPayload() []*models.TaskRun {
	return o.Payload
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest creates a GetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest with default headers values
func NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest() *GetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest {
	return &GetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest{}
}

/*GetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest handles this case with default header values.

Bad Request
*/
type GetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/{task_type}/{task_id}/history][%d] getClusterClusterIdTaskTaskTypeTaskIdHistoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound creates a GetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound with default headers values
func NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound() *GetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound {
	return &GetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound{}
}

/*GetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound handles this case with default header values.

Not found
*/
type GetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/{task_type}/{task_id}/history][%d] getClusterClusterIdTaskTaskTypeTaskIdHistoryNotFound  %+v", 404, o.Payload)
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError creates a GetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError with default headers values
func NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError() *GetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError {
	return &GetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError{}
}

/*GetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError handles this case with default header values.

Server error
*/
type GetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/{task_type}/{task_id}/history][%d] getClusterClusterIdTaskTaskTypeTaskIdHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault creates a GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault with default headers values
func NewGetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault(code int) *GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault {
	return &GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault{
		_statusCode: code,
	}
}

/*GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault handles this case with default header values.

Unexpected error
*/
type GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster cluster ID task task type task ID history default response
func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/task/{task_type}/{task_id}/history][%d] GetClusterClusterIDTaskTaskTypeTaskIDHistory default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDTaskTaskTypeTaskIDHistoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
