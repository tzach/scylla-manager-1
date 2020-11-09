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

// CollectdByPluginidPostReader is a Reader for the CollectdByPluginidPost structure.
type CollectdByPluginidPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectdByPluginidPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectdByPluginidPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCollectdByPluginidPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCollectdByPluginidPostOK creates a CollectdByPluginidPostOK with default headers values
func NewCollectdByPluginidPostOK() *CollectdByPluginidPostOK {
	return &CollectdByPluginidPostOK{}
}

/*CollectdByPluginidPostOK handles this case with default header values.

CollectdByPluginidPostOK collectd by pluginid post o k
*/
type CollectdByPluginidPostOK struct {
}

func (o *CollectdByPluginidPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCollectdByPluginidPostDefault creates a CollectdByPluginidPostDefault with default headers values
func NewCollectdByPluginidPostDefault(code int) *CollectdByPluginidPostDefault {
	return &CollectdByPluginidPostDefault{
		_statusCode: code,
	}
}

/*CollectdByPluginidPostDefault handles this case with default header values.

internal server error
*/
type CollectdByPluginidPostDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the collectd by pluginid post default response
func (o *CollectdByPluginidPostDefault) Code() int {
	return o._statusCode
}

func (o *CollectdByPluginidPostDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *CollectdByPluginidPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *CollectdByPluginidPostDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
