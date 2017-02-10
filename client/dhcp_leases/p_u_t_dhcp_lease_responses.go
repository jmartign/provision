package dhcp_leases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// PUTDhcpLeaseReader is a Reader for the PUTDhcpLease structure.
type PUTDhcpLeaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PUTDhcpLeaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPUTDhcpLeaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPUTDhcpLeaseOK creates a PUTDhcpLeaseOK with default headers values
func NewPUTDhcpLeaseOK() *PUTDhcpLeaseOK {
	return &PUTDhcpLeaseOK{}
}

/*PUTDhcpLeaseOK handles this case with default header values.

PUTDhcpLeaseOK p u t dhcp lease o k
*/
type PUTDhcpLeaseOK struct {
	Payload *models.DhcpLeaseInput
}

func (o *PUTDhcpLeaseOK) Error() string {
	return fmt.Sprintf("[PUT /leases/{id}][%d] pUTDhcpLeaseOK  %+v", 200, o.Payload)
}

func (o *PUTDhcpLeaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DhcpLeaseInput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
