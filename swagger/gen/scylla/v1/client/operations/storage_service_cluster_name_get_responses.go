// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylla-manager/swagger/gen/scylla/v1/models"
)

// StorageServiceClusterNameGetReader is a Reader for the StorageServiceClusterNameGet structure.
type StorageServiceClusterNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceClusterNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceClusterNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceClusterNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceClusterNameGetOK creates a StorageServiceClusterNameGetOK with default headers values
func NewStorageServiceClusterNameGetOK() *StorageServiceClusterNameGetOK {
	return &StorageServiceClusterNameGetOK{}
}

/*StorageServiceClusterNameGetOK handles this case with default header values.

StorageServiceClusterNameGetOK storage service cluster name get o k
*/
type StorageServiceClusterNameGetOK struct {
	Payload string
}

func (o *StorageServiceClusterNameGetOK) GetPayload() string {
	return o.Payload
}

func (o *StorageServiceClusterNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceClusterNameGetDefault creates a StorageServiceClusterNameGetDefault with default headers values
func NewStorageServiceClusterNameGetDefault(code int) *StorageServiceClusterNameGetDefault {
	return &StorageServiceClusterNameGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceClusterNameGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceClusterNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service cluster name get default response
func (o *StorageServiceClusterNameGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceClusterNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceClusterNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceClusterNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
