// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateWebhookByIDReader is a Reader for the UpdateWebhookByID structure.
type UpdateWebhookByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWebhookByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateWebhookByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /webhooks/id/{id}] updateWebhookById", response, response.Code())
	}
}

// NewUpdateWebhookByIDCreated creates a UpdateWebhookByIDCreated with default headers values
func NewUpdateWebhookByIDCreated() *UpdateWebhookByIDCreated {
	return &UpdateWebhookByIDCreated{}
}

/*
UpdateWebhookByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateWebhookByIDCreated struct {
}

// IsSuccess returns true when this update webhook by Id created response has a 2xx status code
func (o *UpdateWebhookByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update webhook by Id created response has a 3xx status code
func (o *UpdateWebhookByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update webhook by Id created response has a 4xx status code
func (o *UpdateWebhookByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update webhook by Id created response has a 5xx status code
func (o *UpdateWebhookByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update webhook by Id created response a status code equal to that given
func (o *UpdateWebhookByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update webhook by Id created response
func (o *UpdateWebhookByIDCreated) Code() int {
	return 201
}

func (o *UpdateWebhookByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /webhooks/id/{id}][%d] updateWebhookByIdCreated ", 201)
}

func (o *UpdateWebhookByIDCreated) String() string {
	return fmt.Sprintf("[PUT /webhooks/id/{id}][%d] updateWebhookByIdCreated ", 201)
}

func (o *UpdateWebhookByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
