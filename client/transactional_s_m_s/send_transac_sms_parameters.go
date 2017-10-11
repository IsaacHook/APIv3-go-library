// Code generated by go-swagger; DO NOT EDIT.

package transactional_s_m_s

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

	"github.com/sendinblue/APIv3-go-library/models"
)

// NewSendTransacSmsParams creates a new SendTransacSmsParams object
// with the default values initialized.
func NewSendTransacSmsParams() *SendTransacSmsParams {
	var ()
	return &SendTransacSmsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendTransacSmsParamsWithTimeout creates a new SendTransacSmsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendTransacSmsParamsWithTimeout(timeout time.Duration) *SendTransacSmsParams {
	var ()
	return &SendTransacSmsParams{

		timeout: timeout,
	}
}

// NewSendTransacSmsParamsWithContext creates a new SendTransacSmsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendTransacSmsParamsWithContext(ctx context.Context) *SendTransacSmsParams {
	var ()
	return &SendTransacSmsParams{

		Context: ctx,
	}
}

// NewSendTransacSmsParamsWithHTTPClient creates a new SendTransacSmsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendTransacSmsParamsWithHTTPClient(client *http.Client) *SendTransacSmsParams {
	var ()
	return &SendTransacSmsParams{
		HTTPClient: client,
	}
}

/*SendTransacSmsParams contains all the parameters to send to the API endpoint
for the send transac sms operation typically these are written to a http.Request
*/
type SendTransacSmsParams struct {

	/*SendTransacSms
	  Values to send a transactional SMS

	*/
	SendTransacSms *models.SendTransacSms

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send transac sms params
func (o *SendTransacSmsParams) WithTimeout(timeout time.Duration) *SendTransacSmsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send transac sms params
func (o *SendTransacSmsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send transac sms params
func (o *SendTransacSmsParams) WithContext(ctx context.Context) *SendTransacSmsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send transac sms params
func (o *SendTransacSmsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send transac sms params
func (o *SendTransacSmsParams) WithHTTPClient(client *http.Client) *SendTransacSmsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send transac sms params
func (o *SendTransacSmsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSendTransacSms adds the sendTransacSms to the send transac sms params
func (o *SendTransacSmsParams) WithSendTransacSms(sendTransacSms *models.SendTransacSms) *SendTransacSmsParams {
	o.SetSendTransacSms(sendTransacSms)
	return o
}

// SetSendTransacSms adds the sendTransacSms to the send transac sms params
func (o *SendTransacSmsParams) SetSendTransacSms(sendTransacSms *models.SendTransacSms) {
	o.SendTransacSms = sendTransacSms
}

// WriteToRequest writes these params to a swagger request
func (o *SendTransacSmsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SendTransacSms == nil {
		o.SendTransacSms = new(models.SendTransacSms)
	}

	if err := r.SetBodyParam(o.SendTransacSms); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}