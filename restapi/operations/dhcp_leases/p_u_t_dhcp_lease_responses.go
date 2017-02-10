package dhcp_leases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/rackn/rocket-skates/models"
)

/*PUTDhcpLeaseOK p u t dhcp lease o k

swagger:response pUTDhcpLeaseOK
*/
type PUTDhcpLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.DhcpLeaseInput `json:"body,omitempty"`
}

// NewPUTDhcpLeaseOK creates PUTDhcpLeaseOK with default headers values
func NewPUTDhcpLeaseOK() *PUTDhcpLeaseOK {
	return &PUTDhcpLeaseOK{}
}

// WithPayload adds the payload to the p u t dhcp lease o k response
func (o *PUTDhcpLeaseOK) WithPayload(payload *models.DhcpLeaseInput) *PUTDhcpLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p u t dhcp lease o k response
func (o *PUTDhcpLeaseOK) SetPayload(payload *models.DhcpLeaseInput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PUTDhcpLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
