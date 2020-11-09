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

// ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetReader is a Reader for the ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGet structure.
type ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK creates a ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK with default headers values
func NewColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK() *ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK {
	return &ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK{}
}

/*ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK handles this case with default header values.

ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK column family metrics write latency estimated recent histogram by name get o k
*/
type ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK struct {
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault creates a ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault with default headers values
func NewColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault(code int) *ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault {
	return &ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault{
		_statusCode: code,
	}
}

/*ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault handles this case with default header values.

internal server error
*/
type ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the column family metrics write latency estimated recent histogram by name get default response
func (o *ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault) Code() int {
	return o._statusCode
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *ColumnFamilyMetricsWriteLatencyEstimatedRecentHistogramByNameGetDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
