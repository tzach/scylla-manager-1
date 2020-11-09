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

// StorageServiceIsStartingGetReader is a Reader for the StorageServiceIsStartingGet structure.
type StorageServiceIsStartingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageServiceIsStartingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageServiceIsStartingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageServiceIsStartingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageServiceIsStartingGetOK creates a StorageServiceIsStartingGetOK with default headers values
func NewStorageServiceIsStartingGetOK() *StorageServiceIsStartingGetOK {
	return &StorageServiceIsStartingGetOK{}
}

/*StorageServiceIsStartingGetOK handles this case with default header values.

StorageServiceIsStartingGetOK storage service is starting get o k
*/
type StorageServiceIsStartingGetOK struct {
	Payload bool
}

func (o *StorageServiceIsStartingGetOK) GetPayload() bool {
	return o.Payload
}

func (o *StorageServiceIsStartingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageServiceIsStartingGetDefault creates a StorageServiceIsStartingGetDefault with default headers values
func NewStorageServiceIsStartingGetDefault(code int) *StorageServiceIsStartingGetDefault {
	return &StorageServiceIsStartingGetDefault{
		_statusCode: code,
	}
}

/*StorageServiceIsStartingGetDefault handles this case with default header values.

internal server error
*/
type StorageServiceIsStartingGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage service is starting get default response
func (o *StorageServiceIsStartingGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageServiceIsStartingGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageServiceIsStartingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageServiceIsStartingGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
