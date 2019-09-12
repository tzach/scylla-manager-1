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

// DeleteClusterClusterIDReader is a Reader for the DeleteClusterClusterID structure.
type DeleteClusterClusterIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterClusterIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClusterClusterIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteClusterClusterIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteClusterClusterIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteClusterClusterIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteClusterClusterIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterClusterIDOK creates a DeleteClusterClusterIDOK with default headers values
func NewDeleteClusterClusterIDOK() *DeleteClusterClusterIDOK {
	return &DeleteClusterClusterIDOK{}
}

/*DeleteClusterClusterIDOK handles this case with default header values.

Cluster deleted
*/
type DeleteClusterClusterIDOK struct {
}

func (o *DeleteClusterClusterIDOK) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}][%d] deleteClusterClusterIdOK ", 200)
}

func (o *DeleteClusterClusterIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClusterClusterIDBadRequest creates a DeleteClusterClusterIDBadRequest with default headers values
func NewDeleteClusterClusterIDBadRequest() *DeleteClusterClusterIDBadRequest {
	return &DeleteClusterClusterIDBadRequest{}
}

/*DeleteClusterClusterIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteClusterClusterIDBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DeleteClusterClusterIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}][%d] deleteClusterClusterIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteClusterClusterIDBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterClusterIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterClusterIDNotFound creates a DeleteClusterClusterIDNotFound with default headers values
func NewDeleteClusterClusterIDNotFound() *DeleteClusterClusterIDNotFound {
	return &DeleteClusterClusterIDNotFound{}
}

/*DeleteClusterClusterIDNotFound handles this case with default header values.

Not found
*/
type DeleteClusterClusterIDNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DeleteClusterClusterIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}][%d] deleteClusterClusterIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteClusterClusterIDNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterClusterIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterClusterIDInternalServerError creates a DeleteClusterClusterIDInternalServerError with default headers values
func NewDeleteClusterClusterIDInternalServerError() *DeleteClusterClusterIDInternalServerError {
	return &DeleteClusterClusterIDInternalServerError{}
}

/*DeleteClusterClusterIDInternalServerError handles this case with default header values.

Server error
*/
type DeleteClusterClusterIDInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *DeleteClusterClusterIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}][%d] deleteClusterClusterIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteClusterClusterIDInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterClusterIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterClusterIDDefault creates a DeleteClusterClusterIDDefault with default headers values
func NewDeleteClusterClusterIDDefault(code int) *DeleteClusterClusterIDDefault {
	return &DeleteClusterClusterIDDefault{
		_statusCode: code,
	}
}

/*DeleteClusterClusterIDDefault handles this case with default header values.

Unexpected error
*/
type DeleteClusterClusterIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete cluster cluster ID default response
func (o *DeleteClusterClusterIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClusterClusterIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /cluster/{cluster_id}][%d] DeleteClusterClusterID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterClusterIDDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DeleteClusterClusterIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
