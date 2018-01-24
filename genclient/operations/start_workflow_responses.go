// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/workflow-goclient/models"
)

// StartWorkflowReader is a Reader for the StartWorkflow structure.
type StartWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStartWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewStartWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewStartWorkflowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewStartWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewStartWorkflowDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartWorkflowOK creates a StartWorkflowOK with default headers values
func NewStartWorkflowOK() *StartWorkflowOK {
	return &StartWorkflowOK{}
}

/*StartWorkflowOK handles this case with default header values.

Successfully started the workflow, returns workflow ID
*/
type StartWorkflowOK struct {
	Payload string
}

func (o *StartWorkflowOK) Error() string {
	return fmt.Sprintf("[POST /workflows][%d] startWorkflowOK  %+v", 200, o.Payload)
}

func (o *StartWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartWorkflowUnauthorized creates a StartWorkflowUnauthorized with default headers values
func NewStartWorkflowUnauthorized() *StartWorkflowUnauthorized {
	return &StartWorkflowUnauthorized{}
}

/*StartWorkflowUnauthorized handles this case with default header values.

Not authorized
*/
type StartWorkflowUnauthorized struct {
	Payload *models.Error
}

func (o *StartWorkflowUnauthorized) Error() string {
	return fmt.Sprintf("[POST /workflows][%d] startWorkflowUnauthorized  %+v", 401, o.Payload)
}

func (o *StartWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartWorkflowForbidden creates a StartWorkflowForbidden with default headers values
func NewStartWorkflowForbidden() *StartWorkflowForbidden {
	return &StartWorkflowForbidden{}
}

/*StartWorkflowForbidden handles this case with default header values.

Forbidden
*/
type StartWorkflowForbidden struct {
	Payload *models.Error
}

func (o *StartWorkflowForbidden) Error() string {
	return fmt.Sprintf("[POST /workflows][%d] startWorkflowForbidden  %+v", 403, o.Payload)
}

func (o *StartWorkflowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartWorkflowNotFound creates a StartWorkflowNotFound with default headers values
func NewStartWorkflowNotFound() *StartWorkflowNotFound {
	return &StartWorkflowNotFound{}
}

/*StartWorkflowNotFound handles this case with default header values.

Resource not found
*/
type StartWorkflowNotFound struct {
	Payload *models.Error
}

func (o *StartWorkflowNotFound) Error() string {
	return fmt.Sprintf("[POST /workflows][%d] startWorkflowNotFound  %+v", 404, o.Payload)
}

func (o *StartWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartWorkflowDefault creates a StartWorkflowDefault with default headers values
func NewStartWorkflowDefault(code int) *StartWorkflowDefault {
	return &StartWorkflowDefault{
		_statusCode: code,
	}
}

/*StartWorkflowDefault handles this case with default header values.

error
*/
type StartWorkflowDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the start workflow default response
func (o *StartWorkflowDefault) Code() int {
	return o._statusCode
}

func (o *StartWorkflowDefault) Error() string {
	return fmt.Sprintf("[POST /workflows][%d] startWorkflow default  %+v", o._statusCode, o.Payload)
}

func (o *StartWorkflowDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
