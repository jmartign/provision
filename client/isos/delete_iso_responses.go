package isos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// DeleteIsoReader is a Reader for the DeleteIso structure.
type DeleteIsoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIsoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteIsoNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteIsoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteIsoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteIsoConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteIsoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteIsoNoContent creates a DeleteIsoNoContent with default headers values
func NewDeleteIsoNoContent() *DeleteIsoNoContent {
	return &DeleteIsoNoContent{}
}

/*DeleteIsoNoContent handles this case with default header values.

DeleteIsoNoContent delete iso no content
*/
type DeleteIsoNoContent struct {
}

func (o *DeleteIsoNoContent) Error() string {
	return fmt.Sprintf("[DELETE /isos/{name}][%d] deleteIsoNoContent ", 204)
}

func (o *DeleteIsoNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIsoUnauthorized creates a DeleteIsoUnauthorized with default headers values
func NewDeleteIsoUnauthorized() *DeleteIsoUnauthorized {
	return &DeleteIsoUnauthorized{}
}

/*DeleteIsoUnauthorized handles this case with default header values.

DeleteIsoUnauthorized delete iso unauthorized
*/
type DeleteIsoUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteIsoUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /isos/{name}][%d] deleteIsoUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIsoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIsoNotFound creates a DeleteIsoNotFound with default headers values
func NewDeleteIsoNotFound() *DeleteIsoNotFound {
	return &DeleteIsoNotFound{}
}

/*DeleteIsoNotFound handles this case with default header values.

DeleteIsoNotFound delete iso not found
*/
type DeleteIsoNotFound struct {
	Payload *models.Error
}

func (o *DeleteIsoNotFound) Error() string {
	return fmt.Sprintf("[DELETE /isos/{name}][%d] deleteIsoNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIsoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIsoConflict creates a DeleteIsoConflict with default headers values
func NewDeleteIsoConflict() *DeleteIsoConflict {
	return &DeleteIsoConflict{}
}

/*DeleteIsoConflict handles this case with default header values.

DeleteIsoConflict delete iso conflict
*/
type DeleteIsoConflict struct {
	Payload *models.Error
}

func (o *DeleteIsoConflict) Error() string {
	return fmt.Sprintf("[DELETE /isos/{name}][%d] deleteIsoConflict  %+v", 409, o.Payload)
}

func (o *DeleteIsoConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIsoInternalServerError creates a DeleteIsoInternalServerError with default headers values
func NewDeleteIsoInternalServerError() *DeleteIsoInternalServerError {
	return &DeleteIsoInternalServerError{}
}

/*DeleteIsoInternalServerError handles this case with default header values.

DeleteIsoInternalServerError delete iso internal server error
*/
type DeleteIsoInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteIsoInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /isos/{name}][%d] deleteIsoInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIsoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
