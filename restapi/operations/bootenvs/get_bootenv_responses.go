package bootenvs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*GetBootenvOK get bootenv o k

swagger:response getBootenvOK
*/
type GetBootenvOK struct {

	/*
	  In: Body
	*/
	Payload *models.BootenvOutput `json:"body,omitempty"`
}

// NewGetBootenvOK creates GetBootenvOK with default headers values
func NewGetBootenvOK() *GetBootenvOK {
	return &GetBootenvOK{}
}

// WithPayload adds the payload to the get bootenv o k response
func (o *GetBootenvOK) WithPayload(payload *models.BootenvOutput) *GetBootenvOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bootenv o k response
func (o *GetBootenvOK) SetPayload(payload *models.BootenvOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBootenvOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBootenvUnauthorized get bootenv unauthorized

swagger:response getBootenvUnauthorized
*/
type GetBootenvUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBootenvUnauthorized creates GetBootenvUnauthorized with default headers values
func NewGetBootenvUnauthorized() *GetBootenvUnauthorized {
	return &GetBootenvUnauthorized{}
}

// WithPayload adds the payload to the get bootenv unauthorized response
func (o *GetBootenvUnauthorized) WithPayload(payload *models.Error) *GetBootenvUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bootenv unauthorized response
func (o *GetBootenvUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBootenvUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBootenvNotFound get bootenv not found

swagger:response getBootenvNotFound
*/
type GetBootenvNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBootenvNotFound creates GetBootenvNotFound with default headers values
func NewGetBootenvNotFound() *GetBootenvNotFound {
	return &GetBootenvNotFound{}
}

// WithPayload adds the payload to the get bootenv not found response
func (o *GetBootenvNotFound) WithPayload(payload *models.Error) *GetBootenvNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bootenv not found response
func (o *GetBootenvNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBootenvNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetBootenvInternalServerError get bootenv internal server error

swagger:response getBootenvInternalServerError
*/
type GetBootenvInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetBootenvInternalServerError creates GetBootenvInternalServerError with default headers values
func NewGetBootenvInternalServerError() *GetBootenvInternalServerError {
	return &GetBootenvInternalServerError{}
}

// WithPayload adds the payload to the get bootenv internal server error response
func (o *GetBootenvInternalServerError) WithPayload(payload *models.Error) *GetBootenvInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get bootenv internal server error response
func (o *GetBootenvInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBootenvInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
