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

// StorageProxyMetricsRangeUnavailablesRatesGetReader is a Reader for the StorageProxyMetricsRangeUnavailablesRatesGet structure.
type StorageProxyMetricsRangeUnavailablesRatesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageProxyMetricsRangeUnavailablesRatesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageProxyMetricsRangeUnavailablesRatesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageProxyMetricsRangeUnavailablesRatesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageProxyMetricsRangeUnavailablesRatesGetOK creates a StorageProxyMetricsRangeUnavailablesRatesGetOK with default headers values
func NewStorageProxyMetricsRangeUnavailablesRatesGetOK() *StorageProxyMetricsRangeUnavailablesRatesGetOK {
	return &StorageProxyMetricsRangeUnavailablesRatesGetOK{}
}

/*StorageProxyMetricsRangeUnavailablesRatesGetOK handles this case with default header values.

StorageProxyMetricsRangeUnavailablesRatesGetOK storage proxy metrics range unavailables rates get o k
*/
type StorageProxyMetricsRangeUnavailablesRatesGetOK struct {
	Payload *models.RateMovingAverage
}

func (o *StorageProxyMetricsRangeUnavailablesRatesGetOK) GetPayload() *models.RateMovingAverage {
	return o.Payload
}

func (o *StorageProxyMetricsRangeUnavailablesRatesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateMovingAverage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageProxyMetricsRangeUnavailablesRatesGetDefault creates a StorageProxyMetricsRangeUnavailablesRatesGetDefault with default headers values
func NewStorageProxyMetricsRangeUnavailablesRatesGetDefault(code int) *StorageProxyMetricsRangeUnavailablesRatesGetDefault {
	return &StorageProxyMetricsRangeUnavailablesRatesGetDefault{
		_statusCode: code,
	}
}

/*StorageProxyMetricsRangeUnavailablesRatesGetDefault handles this case with default header values.

internal server error
*/
type StorageProxyMetricsRangeUnavailablesRatesGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the storage proxy metrics range unavailables rates get default response
func (o *StorageProxyMetricsRangeUnavailablesRatesGetDefault) Code() int {
	return o._statusCode
}

func (o *StorageProxyMetricsRangeUnavailablesRatesGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *StorageProxyMetricsRangeUnavailablesRatesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *StorageProxyMetricsRangeUnavailablesRatesGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
