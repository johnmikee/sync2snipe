// Code generated by go-swagger; DO NOT EDIT.

package printers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePrinterByNameReader is a Reader for the DeletePrinterByName structure.
type DeletePrinterByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePrinterByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePrinterByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /printers/name/{name}] deletePrinterByName", response, response.Code())
	}
}

// NewDeletePrinterByNameOK creates a DeletePrinterByNameOK with default headers values
func NewDeletePrinterByNameOK() *DeletePrinterByNameOK {
	return &DeletePrinterByNameOK{}
}

/*
DeletePrinterByNameOK describes a response with status code 200, with default header values.

OK
*/
type DeletePrinterByNameOK struct {
}

// IsSuccess returns true when this delete printer by name o k response has a 2xx status code
func (o *DeletePrinterByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete printer by name o k response has a 3xx status code
func (o *DeletePrinterByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete printer by name o k response has a 4xx status code
func (o *DeletePrinterByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete printer by name o k response has a 5xx status code
func (o *DeletePrinterByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete printer by name o k response a status code equal to that given
func (o *DeletePrinterByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete printer by name o k response
func (o *DeletePrinterByNameOK) Code() int {
	return 200
}

func (o *DeletePrinterByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /printers/name/{name}][%d] deletePrinterByNameOK ", 200)
}

func (o *DeletePrinterByNameOK) String() string {
	return fmt.Sprintf("[DELETE /printers/name/{name}][%d] deletePrinterByNameOK ", 200)
}

func (o *DeletePrinterByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}