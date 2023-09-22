// Code generated by go-swagger; DO NOT EDIT.

package categories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteCategoryByNameReader is a Reader for the DeleteCategoryByName structure.
type DeleteCategoryByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCategoryByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCategoryByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /categories/name/{name}] deleteCategoryByName", response, response.Code())
	}
}

// NewDeleteCategoryByNameOK creates a DeleteCategoryByNameOK with default headers values
func NewDeleteCategoryByNameOK() *DeleteCategoryByNameOK {
	return &DeleteCategoryByNameOK{}
}

/*
DeleteCategoryByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeleteCategoryByNameOK struct {
}

// IsSuccess returns true when this delete category by name o k response has a 2xx status code
func (o *DeleteCategoryByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete category by name o k response has a 3xx status code
func (o *DeleteCategoryByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete category by name o k response has a 4xx status code
func (o *DeleteCategoryByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete category by name o k response has a 5xx status code
func (o *DeleteCategoryByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete category by name o k response a status code equal to that given
func (o *DeleteCategoryByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete category by name o k response
func (o *DeleteCategoryByNameOK) Code() int {
	return 200
}

func (o *DeleteCategoryByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /categories/name/{name}][%d] deleteCategoryByNameOK ", 200)
}

func (o *DeleteCategoryByNameOK) String() string {
	return fmt.Sprintf("[DELETE /categories/name/{name}][%d] deleteCategoryByNameOK ", 200)
}

func (o *DeleteCategoryByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}