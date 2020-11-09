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

// CacheServiceMetricsKeyHitsGetReader is a Reader for the CacheServiceMetricsKeyHitsGet structure.
type CacheServiceMetricsKeyHitsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CacheServiceMetricsKeyHitsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCacheServiceMetricsKeyHitsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCacheServiceMetricsKeyHitsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCacheServiceMetricsKeyHitsGetOK creates a CacheServiceMetricsKeyHitsGetOK with default headers values
func NewCacheServiceMetricsKeyHitsGetOK() *CacheServiceMetricsKeyHitsGetOK {
	return &CacheServiceMetricsKeyHitsGetOK{}
}

/*CacheServiceMetricsKeyHitsGetOK handles this case with default header values.

CacheServiceMetricsKeyHitsGetOK cache service metrics key hits get o k
*/
type CacheServiceMetricsKeyHitsGetOK struct {
	Payload interface{}
}

func (o *CacheServiceMetricsKeyHitsGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CacheServiceMetricsKeyHitsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCacheServiceMetricsKeyHitsGetDefault creates a CacheServiceMetricsKeyHitsGetDefault with default headers values
func NewCacheServiceMetricsKeyHitsGetDefault(code int) *CacheServiceMetricsKeyHitsGetDefault {
	return &CacheServiceMetricsKeyHitsGetDefault{
		_statusCode: code,
	}
}

/*CacheServiceMetricsKeyHitsGetDefault handles this case with default header values.

internal server error
*/
type CacheServiceMetricsKeyHitsGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the cache service metrics key hits get default response
func (o *CacheServiceMetricsKeyHitsGetDefault) Code() int {
	return o._statusCode
}

func (o *CacheServiceMetricsKeyHitsGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CacheServiceMetricsKeyHitsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CacheServiceMetricsKeyHitsGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
