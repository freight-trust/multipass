// Code generated by go-swagger; DO NOT EDIT.

package credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/lacchain/credential-server/models"
)

// SendCredentialOKCode is the HTTP code returned for type SendCredentialOK
const SendCredentialOKCode int = 200

/*SendCredentialOK successful operation

swagger:response sendCredentialOK
*/
type SendCredentialOK struct {

	/*
	  In: Body
	*/
	Payload *models.VerifyResponse `json:"body,omitempty"`
}

// NewSendCredentialOK creates SendCredentialOK with default headers values
func NewSendCredentialOK() *SendCredentialOK {

	return &SendCredentialOK{}
}

// WithPayload adds the payload to the send credential o k response
func (o *SendCredentialOK) WithPayload(payload *models.VerifyResponse) *SendCredentialOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the send credential o k response
func (o *SendCredentialOK) SetPayload(payload *models.VerifyResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SendCredentialOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SendCredentialBadRequestCode is the HTTP code returned for type SendCredentialBadRequest
const SendCredentialBadRequestCode int = 400

/*SendCredentialBadRequest Invalid credential supplied

swagger:response sendCredentialBadRequest
*/
type SendCredentialBadRequest struct {
}

// NewSendCredentialBadRequest creates SendCredentialBadRequest with default headers values
func NewSendCredentialBadRequest() *SendCredentialBadRequest {

	return &SendCredentialBadRequest{}
}

// WriteResponse to the client
func (o *SendCredentialBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// SendCredentialNotFoundCode is the HTTP code returned for type SendCredentialNotFound
const SendCredentialNotFoundCode int = 404

/*SendCredentialNotFound DID not found

swagger:response sendCredentialNotFound
*/
type SendCredentialNotFound struct {
}

// NewSendCredentialNotFound creates SendCredentialNotFound with default headers values
func NewSendCredentialNotFound() *SendCredentialNotFound {

	return &SendCredentialNotFound{}
}

// WriteResponse to the client
func (o *SendCredentialNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// SendCredentialInternalServerErrorCode is the HTTP code returned for type SendCredentialInternalServerError
const SendCredentialInternalServerErrorCode int = 500

/*SendCredentialInternalServerError Error Internal Server

swagger:response sendCredentialInternalServerError
*/
type SendCredentialInternalServerError struct {
}

// NewSendCredentialInternalServerError creates SendCredentialInternalServerError with default headers values
func NewSendCredentialInternalServerError() *SendCredentialInternalServerError {

	return &SendCredentialInternalServerError{}
}

// WriteResponse to the client
func (o *SendCredentialInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
