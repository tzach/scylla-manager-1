// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-manager/pkg/scyllaclient/internal/agent/models"
)

// NewOperationsCopyfileParams creates a new OperationsCopyfileParams object
// with the default values initialized.
func NewOperationsCopyfileParams() *OperationsCopyfileParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsCopyfileParams{
		Async: asyncDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewOperationsCopyfileParamsWithTimeout creates a new OperationsCopyfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperationsCopyfileParamsWithTimeout(timeout time.Duration) *OperationsCopyfileParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsCopyfileParams{
		Async: asyncDefault,

		timeout: timeout,
	}
}

// NewOperationsCopyfileParamsWithContext creates a new OperationsCopyfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperationsCopyfileParamsWithContext(ctx context.Context) *OperationsCopyfileParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsCopyfileParams{
		Async: asyncDefault,

		Context: ctx,
	}
}

// NewOperationsCopyfileParamsWithHTTPClient creates a new OperationsCopyfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperationsCopyfileParamsWithHTTPClient(client *http.Client) *OperationsCopyfileParams {
	var (
		asyncDefault = bool(true)
	)
	return &OperationsCopyfileParams{
		Async:      asyncDefault,
		HTTPClient: client,
	}
}

/*OperationsCopyfileParams contains all the parameters to send to the API endpoint
for the operations copyfile operation typically these are written to a http.Request
*/
type OperationsCopyfileParams struct {

	/*Async
	  Async request

	*/
	Async bool
	/*Group
	  Place this operation under this stat group

	*/
	Group string
	/*Copyfile
	  copyfile

	*/
	Copyfile *models.MoveOrCopyFileOptions

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operations copyfile params
func (o *OperationsCopyfileParams) WithTimeout(timeout time.Duration) *OperationsCopyfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operations copyfile params
func (o *OperationsCopyfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operations copyfile params
func (o *OperationsCopyfileParams) WithContext(ctx context.Context) *OperationsCopyfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operations copyfile params
func (o *OperationsCopyfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operations copyfile params
func (o *OperationsCopyfileParams) WithHTTPClient(client *http.Client) *OperationsCopyfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operations copyfile params
func (o *OperationsCopyfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsync adds the async to the operations copyfile params
func (o *OperationsCopyfileParams) WithAsync(async bool) *OperationsCopyfileParams {
	o.SetAsync(async)
	return o
}

// SetAsync adds the async to the operations copyfile params
func (o *OperationsCopyfileParams) SetAsync(async bool) {
	o.Async = async
}

// WithGroup adds the group to the operations copyfile params
func (o *OperationsCopyfileParams) WithGroup(group string) *OperationsCopyfileParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the operations copyfile params
func (o *OperationsCopyfileParams) SetGroup(group string) {
	o.Group = group
}

// WithCopyfile adds the copyfile to the operations copyfile params
func (o *OperationsCopyfileParams) WithCopyfile(copyfile *models.MoveOrCopyFileOptions) *OperationsCopyfileParams {
	o.SetCopyfile(copyfile)
	return o
}

// SetCopyfile adds the copyfile to the operations copyfile params
func (o *OperationsCopyfileParams) SetCopyfile(copyfile *models.MoveOrCopyFileOptions) {
	o.Copyfile = copyfile
}

// WriteToRequest writes these params to a swagger request
func (o *OperationsCopyfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param _async
	qrAsync := o.Async
	qAsync := swag.FormatBool(qrAsync)
	if qAsync != "" {
		if err := r.SetQueryParam("_async", qAsync); err != nil {
			return err
		}
	}

	// query param _group
	qrGroup := o.Group
	qGroup := qrGroup
	if qGroup != "" {
		if err := r.SetQueryParam("_group", qGroup); err != nil {
			return err
		}
	}

	if o.Copyfile != nil {
		if err := r.SetBodyParam(o.Copyfile); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
