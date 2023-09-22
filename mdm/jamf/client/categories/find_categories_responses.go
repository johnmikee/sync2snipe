// Code generated by go-swagger; DO NOT EDIT.

package categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/johnmikee/sync2snipe/mdm/jamf/models"
)

// FindCategoriesReader is a Reader for the FindCategories structure.
type FindCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /categories] findCategories", response, response.Code())
	}
}

// NewFindCategoriesOK creates a FindCategoriesOK with default headers values
func NewFindCategoriesOK() *FindCategoriesOK {
	return &FindCategoriesOK{}
}

/*
FindCategoriesOK describes a response with status code 200, with default header values.

OK
*/
type FindCategoriesOK struct {
	Payload models.Categories
}

// IsSuccess returns true when this find categories o k response has a 2xx status code
func (o *FindCategoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find categories o k response has a 3xx status code
func (o *FindCategoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find categories o k response has a 4xx status code
func (o *FindCategoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find categories o k response has a 5xx status code
func (o *FindCategoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find categories o k response a status code equal to that given
func (o *FindCategoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the find categories o k response
func (o *FindCategoriesOK) Code() int {
	return 200
}

func (o *FindCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /categories][%d] findCategoriesOK  %+v", 200, o.Payload)
}

func (o *FindCategoriesOK) String() string {
	return fmt.Sprintf("[GET /categories][%d] findCategoriesOK  %+v", 200, o.Payload)
}

func (o *FindCategoriesOK) GetPayload() models.Categories {
	return o.Payload
}

func (o *FindCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
