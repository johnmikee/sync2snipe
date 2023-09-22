// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateAccountByIDReader is a Reader for the UpdateAccountByID structure.
type UpdateAccountByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccountByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateAccountByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /accounts/userid/{id}] updateAccountById", response, response.Code())
	}
}

// NewUpdateAccountByIDCreated creates a UpdateAccountByIDCreated with default headers values
func NewUpdateAccountByIDCreated() *UpdateAccountByIDCreated {
	return &UpdateAccountByIDCreated{}
}

/*
UpdateAccountByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateAccountByIDCreated struct {
}

// IsSuccess returns true when this update account by Id created response has a 2xx status code
func (o *UpdateAccountByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update account by Id created response has a 3xx status code
func (o *UpdateAccountByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update account by Id created response has a 4xx status code
func (o *UpdateAccountByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update account by Id created response has a 5xx status code
func (o *UpdateAccountByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update account by Id created response a status code equal to that given
func (o *UpdateAccountByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update account by Id created response
func (o *UpdateAccountByIDCreated) Code() int {
	return 201
}

func (o *UpdateAccountByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /accounts/userid/{id}][%d] updateAccountByIdCreated ", 201)
}

func (o *UpdateAccountByIDCreated) String() string {
	return fmt.Sprintf("[PUT /accounts/userid/{id}][%d] updateAccountByIdCreated ", 201)
}

func (o *UpdateAccountByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
