package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// PostMachineReader is a Reader for the PostMachine structure.
type PostMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewPostMachineCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostMachineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostMachineConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostMachineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostMachineOK creates a PostMachineOK with default headers values
func NewPostMachineOK() *PostMachineOK {
	return &PostMachineOK{}
}

/*PostMachineOK handles this case with default header values.

PostMachineOK post machine o k
*/
type PostMachineOK struct {
	Payload *models.MachineOutput
}

func (o *PostMachineOK) Error() string {
	return fmt.Sprintf("[POST /machines][%d] postMachineOK  %+v", 200, o.Payload)
}

func (o *PostMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMachineCreated creates a PostMachineCreated with default headers values
func NewPostMachineCreated() *PostMachineCreated {
	return &PostMachineCreated{}
}

/*PostMachineCreated handles this case with default header values.

PostMachineCreated post machine created
*/
type PostMachineCreated struct {
	Payload *models.MachineOutput
}

func (o *PostMachineCreated) Error() string {
	return fmt.Sprintf("[POST /machines][%d] postMachineCreated  %+v", 201, o.Payload)
}

func (o *PostMachineCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMachineUnauthorized creates a PostMachineUnauthorized with default headers values
func NewPostMachineUnauthorized() *PostMachineUnauthorized {
	return &PostMachineUnauthorized{}
}

/*PostMachineUnauthorized handles this case with default header values.

PostMachineUnauthorized post machine unauthorized
*/
type PostMachineUnauthorized struct {
	Payload *models.Error
}

func (o *PostMachineUnauthorized) Error() string {
	return fmt.Sprintf("[POST /machines][%d] postMachineUnauthorized  %+v", 401, o.Payload)
}

func (o *PostMachineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMachineConflict creates a PostMachineConflict with default headers values
func NewPostMachineConflict() *PostMachineConflict {
	return &PostMachineConflict{}
}

/*PostMachineConflict handles this case with default header values.

PostMachineConflict post machine conflict
*/
type PostMachineConflict struct {
	Payload *models.Error
}

func (o *PostMachineConflict) Error() string {
	return fmt.Sprintf("[POST /machines][%d] postMachineConflict  %+v", 409, o.Payload)
}

func (o *PostMachineConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMachineInternalServerError creates a PostMachineInternalServerError with default headers values
func NewPostMachineInternalServerError() *PostMachineInternalServerError {
	return &PostMachineInternalServerError{}
}

/*PostMachineInternalServerError handles this case with default header values.

PostMachineInternalServerError post machine internal server error
*/
type PostMachineInternalServerError struct {
	Payload *models.Error
}

func (o *PostMachineInternalServerError) Error() string {
	return fmt.Sprintf("[POST /machines][%d] postMachineInternalServerError  %+v", 500, o.Payload)
}

func (o *PostMachineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
