// Code generated by go-swagger; DO NOT EDIT.

package service_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// ValidateEmailServiceReader is a Reader for the ValidateEmailService structure.
type ValidateEmailServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateEmailServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewValidateEmailServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateEmailServiceOK creates a ValidateEmailServiceOK with default headers values
func NewValidateEmailServiceOK() *ValidateEmailServiceOK {
	return &ValidateEmailServiceOK{}
}

/*ValidateEmailServiceOK handles this case with default header values.

A successful response.
*/
type ValidateEmailServiceOK struct {
	Payload *models.OpenpitrixValidateEmailServiceResponse
}

func (o *ValidateEmailServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/service_configs/validate_email_service][%d] validateEmailServiceOK  %+v", 200, o.Payload)
}

func (o *ValidateEmailServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixValidateEmailServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}