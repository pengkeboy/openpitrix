// Code generated by go-swagger; DO NOT EDIT.

package runtime_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRuntimeStatisticsParams creates a new GetRuntimeStatisticsParams object
// with the default values initialized.
func NewGetRuntimeStatisticsParams() *GetRuntimeStatisticsParams {

	return &GetRuntimeStatisticsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRuntimeStatisticsParamsWithTimeout creates a new GetRuntimeStatisticsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRuntimeStatisticsParamsWithTimeout(timeout time.Duration) *GetRuntimeStatisticsParams {

	return &GetRuntimeStatisticsParams{

		timeout: timeout,
	}
}

// NewGetRuntimeStatisticsParamsWithContext creates a new GetRuntimeStatisticsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRuntimeStatisticsParamsWithContext(ctx context.Context) *GetRuntimeStatisticsParams {

	return &GetRuntimeStatisticsParams{

		Context: ctx,
	}
}

// NewGetRuntimeStatisticsParamsWithHTTPClient creates a new GetRuntimeStatisticsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRuntimeStatisticsParamsWithHTTPClient(client *http.Client) *GetRuntimeStatisticsParams {

	return &GetRuntimeStatisticsParams{
		HTTPClient: client,
	}
}

/*GetRuntimeStatisticsParams contains all the parameters to send to the API endpoint
for the get runtime statistics operation typically these are written to a http.Request
*/
type GetRuntimeStatisticsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runtime statistics params
func (o *GetRuntimeStatisticsParams) WithTimeout(timeout time.Duration) *GetRuntimeStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runtime statistics params
func (o *GetRuntimeStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runtime statistics params
func (o *GetRuntimeStatisticsParams) WithContext(ctx context.Context) *GetRuntimeStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runtime statistics params
func (o *GetRuntimeStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runtime statistics params
func (o *GetRuntimeStatisticsParams) WithHTTPClient(client *http.Client) *GetRuntimeStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runtime statistics params
func (o *GetRuntimeStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRuntimeStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
