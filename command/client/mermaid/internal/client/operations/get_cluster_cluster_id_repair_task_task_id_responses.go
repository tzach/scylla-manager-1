// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/scylladb/mermaid/command/client/mermaid/internal/models"
)

// GetClusterClusterIDRepairTaskTaskIDReader is a Reader for the GetClusterClusterIDRepairTaskTaskID structure.
type GetClusterClusterIDRepairTaskTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDRepairTaskTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterClusterIDRepairTaskTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		body := response.Body()
		if body != nil {
			defer body.Close()

			b, err := ioutil.ReadAll(body)
			if err != nil {
				return nil, err
			}

			buf := new(bytes.Buffer)
			json.Indent(buf, b, "", "  ")

			return nil, errors.New(int32(response.Code()), buf.String())
		}
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClusterClusterIDRepairTaskTaskIDOK creates a GetClusterClusterIDRepairTaskTaskIDOK with default headers values
func NewGetClusterClusterIDRepairTaskTaskIDOK() *GetClusterClusterIDRepairTaskTaskIDOK {
	return &GetClusterClusterIDRepairTaskTaskIDOK{}
}

/*GetClusterClusterIDRepairTaskTaskIDOK handles this case with default header values.

percentage of run completed in total and details
*/
type GetClusterClusterIDRepairTaskTaskIDOK struct {
	Payload *models.RepairProgress
}

func (o *GetClusterClusterIDRepairTaskTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/repair/task/{task_id}][%d] getClusterClusterIdRepairTaskTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDRepairTaskTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepairProgress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
