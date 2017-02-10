package isos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*GetIsoOK get iso o k

swagger:response getIsoOK
*/
type GetIsoOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewGetIsoOK creates GetIsoOK with default headers values
func NewGetIsoOK() *GetIsoOK {
	return &GetIsoOK{}
}

// WithPayload adds the payload to the get iso o k response
func (o *GetIsoOK) WithPayload(payload io.ReadCloser) *GetIsoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get iso o k response
func (o *GetIsoOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIsoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetIsoUnauthorized get iso unauthorized

swagger:response getIsoUnauthorized
*/
type GetIsoUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetIsoUnauthorized creates GetIsoUnauthorized with default headers values
func NewGetIsoUnauthorized() *GetIsoUnauthorized {
	return &GetIsoUnauthorized{}
}

// WithPayload adds the payload to the get iso unauthorized response
func (o *GetIsoUnauthorized) WithPayload(payload *models.Error) *GetIsoUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get iso unauthorized response
func (o *GetIsoUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIsoUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetIsoNotFound get iso not found

swagger:response getIsoNotFound
*/
type GetIsoNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetIsoNotFound creates GetIsoNotFound with default headers values
func NewGetIsoNotFound() *GetIsoNotFound {
	return &GetIsoNotFound{}
}

// WithPayload adds the payload to the get iso not found response
func (o *GetIsoNotFound) WithPayload(payload *models.Error) *GetIsoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get iso not found response
func (o *GetIsoNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIsoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetIsoInternalServerError get iso internal server error

swagger:response getIsoInternalServerError
*/
type GetIsoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetIsoInternalServerError creates GetIsoInternalServerError with default headers values
func NewGetIsoInternalServerError() *GetIsoInternalServerError {
	return &GetIsoInternalServerError{}
}

// WithPayload adds the payload to the get iso internal server error response
func (o *GetIsoInternalServerError) WithPayload(payload *models.Error) *GetIsoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get iso internal server error response
func (o *GetIsoInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIsoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
