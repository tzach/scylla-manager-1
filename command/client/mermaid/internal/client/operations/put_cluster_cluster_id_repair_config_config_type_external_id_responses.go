// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutClusterClusterIDRepairConfigConfigTypeExternalIDReader is a Reader for the PutClusterClusterIDRepairConfigConfigTypeExternalID structure.
type PutClusterClusterIDRepairConfigConfigTypeExternalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutClusterClusterIDRepairConfigConfigTypeExternalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutClusterClusterIDRepairConfigConfigTypeExternalIDOK()
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

// NewPutClusterClusterIDRepairConfigConfigTypeExternalIDOK creates a PutClusterClusterIDRepairConfigConfigTypeExternalIDOK with default headers values
func NewPutClusterClusterIDRepairConfigConfigTypeExternalIDOK() *PutClusterClusterIDRepairConfigConfigTypeExternalIDOK {
	return &PutClusterClusterIDRepairConfigConfigTypeExternalIDOK{}
}

/*PutClusterClusterIDRepairConfigConfigTypeExternalIDOK handles this case with default header values.

OK
*/
type PutClusterClusterIDRepairConfigConfigTypeExternalIDOK struct {
}

func (o *PutClusterClusterIDRepairConfigConfigTypeExternalIDOK) Error() string {
	return fmt.Sprintf("[PUT /cluster/{cluster_id}/repair/config/{config_type}/{external_id}][%d] putClusterClusterIdRepairConfigConfigTypeExternalIdOK ", 200)
}

func (o *PutClusterClusterIDRepairConfigConfigTypeExternalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
