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

// CacheServiceRowCacheSavePeriodGetReader is a Reader for the CacheServiceRowCacheSavePeriodGet structure.
type CacheServiceRowCacheSavePeriodGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceRowCacheSavePeriodGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceRowCacheSavePeriodGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceRowCacheSavePeriodGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceRowCacheSavePeriodGetOK creates a CacheServiceRowCacheSavePeriodGetOK with default headers values
func NewCacheServiceRowCacheSavePeriodGetOK() *CacheServiceRowCacheSavePeriodGetOK {
	return &CacheServiceRowCacheSavePeriodGetOK{}
}

/*CacheServiceRowCacheSavePeriodGetOK handles this case with default header values.

CacheServiceRowCacheSavePeriodGetOK cache service row cache save period get o k
*/
type CacheServiceRowCacheSavePeriodGetOK struct {
	Payload int32
}

func (o *CacheServiceRowCacheSavePeriodGetOK) GetPayload() int32 {
	return o.Payload
}

func (o *CacheServiceRowCacheSavePeriodGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceRowCacheSavePeriodGetDefault creates a CacheServiceRowCacheSavePeriodGetDefault with default headers values
func NewCacheServiceRowCacheSavePeriodGetDefault(code int) *CacheServiceRowCacheSavePeriodGetDefault {
	return &CacheServiceRowCacheSavePeriodGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceRowCacheSavePeriodGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceRowCacheSavePeriodGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service row cache save period get default response
func (o *CacheServiceRowCacheSavePeriodGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceRowCacheSavePeriodGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceRowCacheSavePeriodGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceRowCacheSavePeriodGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
