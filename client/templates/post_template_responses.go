package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// PostTemplateReader is a Reader for the PostTemplate structure.
type PostTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewPostTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTemplateOK creates a PostTemplateOK with default headers values
func NewPostTemplateOK() *PostTemplateOK {
	return &PostTemplateOK{}
}

/*PostTemplateOK handles this case with default header values.

PostTemplateOK post template o k
*/
type PostTemplateOK struct {
	Payload *models.TemplateOutput
}

func (o *PostTemplateOK) Error() string {
	return fmt.Sprintf("[POST /templates][%d] postTemplateOK  %+v", 200, o.Payload)
}

func (o *PostTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTemplateCreated creates a PostTemplateCreated with default headers values
func NewPostTemplateCreated() *PostTemplateCreated {
	return &PostTemplateCreated{}
}

/*PostTemplateCreated handles this case with default header values.

PostTemplateCreated post template created
*/
type PostTemplateCreated struct {
	Payload *models.TemplateOutput
}

func (o *PostTemplateCreated) Error() string {
	return fmt.Sprintf("[POST /templates][%d] postTemplateCreated  %+v", 201, o.Payload)
}

func (o *PostTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTemplateUnauthorized creates a PostTemplateUnauthorized with default headers values
func NewPostTemplateUnauthorized() *PostTemplateUnauthorized {
	return &PostTemplateUnauthorized{}
}

/*PostTemplateUnauthorized handles this case with default header values.

PostTemplateUnauthorized post template unauthorized
*/
type PostTemplateUnauthorized struct {
	Payload *models.Error
}

func (o *PostTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /templates][%d] postTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTemplateConflict creates a PostTemplateConflict with default headers values
func NewPostTemplateConflict() *PostTemplateConflict {
	return &PostTemplateConflict{}
}

/*PostTemplateConflict handles this case with default header values.

PostTemplateConflict post template conflict
*/
type PostTemplateConflict struct {
	Payload *models.Error
}

func (o *PostTemplateConflict) Error() string {
	return fmt.Sprintf("[POST /templates][%d] postTemplateConflict  %+v", 409, o.Payload)
}

func (o *PostTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTemplateInternalServerError creates a PostTemplateInternalServerError with default headers values
func NewPostTemplateInternalServerError() *PostTemplateInternalServerError {
	return &PostTemplateInternalServerError{}
}

/*PostTemplateInternalServerError handles this case with default header values.

PostTemplateInternalServerError post template internal server error
*/
type PostTemplateInternalServerError struct {
	Payload *models.Error
}

func (o *PostTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /templates][%d] postTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
