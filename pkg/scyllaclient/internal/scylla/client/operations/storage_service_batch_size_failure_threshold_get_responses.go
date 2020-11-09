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

// StorageServiceBatchSizeFailureThresholdGetReader is a Reader for the StorageServiceBatchSizeFailureThresholdGet structure.
type StorageServiceBatchSizeFailureThresholdGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceBatchSizeFailureThresholdGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceBatchSizeFailureThresholdGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceBatchSizeFailureThresholdGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceBatchSizeFailureThresholdGetOK creates a StorageServiceBatchSizeFailureThresholdGetOK with default headers values
func NewStorageServiceBatchSizeFailureThresholdGetOK() *StorageServiceBatchSizeFailureThresholdGetOK {
	return &StorageServiceBatchSizeFailureThresholdGetOK{}
}

/*StorageServiceBatchSizeFailureThresholdGetOK handles this case with default header values.

StorageServiceBatchSizeFailureThresholdGetOK storage service batch size failure threshold get o k
*/
type StorageServiceBatchSizeFailureThresholdGetOK struct {
	Payload int32
}

func (o *StorageServiceBatchSizeFailureThresholdGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *StorageServiceBatchSizeFailureThresholdGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceBatchSizeFailureThresholdGetDefault creates a StorageServiceBatchSizeFailureThresholdGetDefault with default headers values
func NewStorageServiceBatchSizeFailureThresholdGetDefault(code int) *StorageServiceBatchSizeFailureThresholdGetDefault {
	return &StorageServiceBatchSizeFailureThresholdGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceBatchSizeFailureThresholdGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceBatchSizeFailureThresholdGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service batch size failure threshold get default response
func (o *StorageServiceBatchSizeFailureThresholdGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceBatchSizeFailureThresholdGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceBatchSizeFailureThresholdGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceBatchSizeFailureThresholdGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
