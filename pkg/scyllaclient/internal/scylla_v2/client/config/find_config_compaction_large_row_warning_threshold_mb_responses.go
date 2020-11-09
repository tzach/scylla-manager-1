// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigCompactionLargeRowWarningThresholdMbReader is a Reader for the FindConfigCompactionLargeRowWarningThresholdMb structure.
type FindConfigCompactionLargeRowWarningThresholdMbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigCompactionLargeRowWarningThresholdMbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigCompactionLargeRowWarningThresholdMbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigCompactionLargeRowWarningThresholdMbDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigCompactionLargeRowWarningThresholdMbOK creates a FindConfigCompactionLargeRowWarningThresholdMbOK with default headers values
func NewFindConfigCompactionLargeRowWarningThresholdMbOK() *FindConfigCompactionLargeRowWarningThresholdMbOK {
	return &FindConfigCompactionLargeRowWarningThresholdMbOK{}
}

/*FindConfigCompactionLargeRowWarningThresholdMbOK handles this case with default header values.

Config value
*/
type FindConfigCompactionLargeRowWarningThresholdMbOK struct {
	Payload int64
}

func (o *FindConfigCompactionLargeRowWarningThresholdMbOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigCompactionLargeRowWarningThresholdMbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigCompactionLargeRowWarningThresholdMbDefault creates a FindConfigCompactionLargeRowWarningThresholdMbDefault with default headers values
func NewFindConfigCompactionLargeRowWarningThresholdMbDefault(code int) *FindConfigCompactionLargeRowWarningThresholdMbDefault {
	return &FindConfigCompactionLargeRowWarningThresholdMbDefault{
		_statusCode: code,
	}
}

/*FindConfigCompactionLargeRowWarningThresholdMbDefault handles this case with default header values.

unexpected error
*/
type FindConfigCompactionLargeRowWarningThresholdMbDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config compaction large row warning threshold mb default response
func (o *FindConfigCompactionLargeRowWarningThresholdMbDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigCompactionLargeRowWarningThresholdMbDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigCompactionLargeRowWarningThresholdMbDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigCompactionLargeRowWarningThresholdMbDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
