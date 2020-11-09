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

// ColumnFamilyNameGetReader is a Reader for the ColumnFamilyNameGet structure.
type ColumnFamilyNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyNameGetOK creates a ColumnFamilyNameGetOK with default headers values
func NewColumnFamilyNameGetOK() *ColumnFamilyNameGetOK {
	return &ColumnFamilyNameGetOK{}
}

/*ColumnFamilyNameGetOK handles this case with default header values.

ColumnFamilyNameGetOK column family name get o k
*/
type ColumnFamilyNameGetOK struct {
	Payload []string
}

func (o *ColumnFamilyNameGetOK) GetPayload() []string {
	return o.Payload
}

func (o *ColumnFamilyNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColumnFamilyNameGetDefault creates a ColumnFamilyNameGetDefault with default headers values
func NewColumnFamilyNameGetDefault(code int) *ColumnFamilyNameGetDefault {
	return &ColumnFamilyNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family name get default response
func (o *ColumnFamilyNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
