// Code generated by go-swagger; DO NOT EDIT.

package compliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetIDParams creates a new GetIDParams object
// with the default values initialized.
func NewGetIDParams() *GetIDParams {
	var ()
	return &GetIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIDParamsWithTimeout creates a new GetIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIDParamsWithTimeout(timeout time.Duration) *GetIDParams {
	var ()
	return &GetIDParams{

		timeout: timeout,
	}
}

// NewGetIDParamsWithContext creates a new GetIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIDParamsWithContext(ctx context.Context) *GetIDParams {
	var ()
	return &GetIDParams{

		Context: ctx,
	}
}

// NewGetIDParamsWithHTTPClient creates a new GetIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIDParamsWithHTTPClient(client *http.Client) *GetIDParams {
	var ()
	return &GetIDParams{
		HTTPClient: client,
	}
}

/*GetIDParams contains all the parameters to send to the API endpoint
for the get ID operation typically these are written to a http.Request
*/
type GetIDParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get ID params
func (o *GetIDParams) WithTimeout(timeout time.Duration) *GetIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ID params
func (o *GetIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ID params
func (o *GetIDParams) WithContext(ctx context.Context) *GetIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ID params
func (o *GetIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ID params
func (o *GetIDParams) WithHTTPClient(client *http.Client) *GetIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ID params
func (o *GetIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get ID params
func (o *GetIDParams) WithID(id string) *GetIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ID params
func (o *GetIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
