// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateSiteByIDReader is a Reader for the UpdateSiteByID structure.
type UpdateSiteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSiteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateSiteByIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /sites/id/{id}] updateSiteById", response, response.Code())
	}
}

// NewUpdateSiteByIDCreated creates a UpdateSiteByIDCreated with default headers values
func NewUpdateSiteByIDCreated() *UpdateSiteByIDCreated {
	return &UpdateSiteByIDCreated{}
}

/*
UpdateSiteByIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateSiteByIDCreated struct {
}

// IsSuccess returns true when this update site by Id created response has a 2xx status code
func (o *UpdateSiteByIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update site by Id created response has a 3xx status code
func (o *UpdateSiteByIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update site by Id created response has a 4xx status code
func (o *UpdateSiteByIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update site by Id created response has a 5xx status code
func (o *UpdateSiteByIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update site by Id created response a status code equal to that given
func (o *UpdateSiteByIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update site by Id created response
func (o *UpdateSiteByIDCreated) Code() int {
	return 201
}

func (o *UpdateSiteByIDCreated) Error() string {
	return fmt.Sprintf("[PUT /sites/id/{id}][%d] updateSiteByIdCreated ", 201)
}

func (o *UpdateSiteByIDCreated) String() string {
	return fmt.Sprintf("[PUT /sites/id/{id}][%d] updateSiteByIdCreated ", 201)
}

func (o *UpdateSiteByIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
