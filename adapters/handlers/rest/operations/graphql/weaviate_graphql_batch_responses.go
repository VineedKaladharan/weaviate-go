//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package graphql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// WeaviateGraphqlBatchOKCode is the HTTP code returned for type WeaviateGraphqlBatchOK
const WeaviateGraphqlBatchOKCode int = 200

/*WeaviateGraphqlBatchOK Successful query (with select).

swagger:response weaviateGraphqlBatchOK
*/
type WeaviateGraphqlBatchOK struct {

	/*
	  In: Body
	*/
	Payload models.GraphQLResponses `json:"body,omitempty"`
}

// NewWeaviateGraphqlBatchOK creates WeaviateGraphqlBatchOK with default headers values
func NewWeaviateGraphqlBatchOK() *WeaviateGraphqlBatchOK {

	return &WeaviateGraphqlBatchOK{}
}

// WithPayload adds the payload to the weaviate graphql batch o k response
func (o *WeaviateGraphqlBatchOK) WithPayload(payload models.GraphQLResponses) *WeaviateGraphqlBatchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate graphql batch o k response
func (o *WeaviateGraphqlBatchOK) SetPayload(payload models.GraphQLResponses) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateGraphqlBatchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.GraphQLResponses, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// WeaviateGraphqlBatchUnauthorizedCode is the HTTP code returned for type WeaviateGraphqlBatchUnauthorized
const WeaviateGraphqlBatchUnauthorizedCode int = 401

/*WeaviateGraphqlBatchUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateGraphqlBatchUnauthorized
*/
type WeaviateGraphqlBatchUnauthorized struct {
}

// NewWeaviateGraphqlBatchUnauthorized creates WeaviateGraphqlBatchUnauthorized with default headers values
func NewWeaviateGraphqlBatchUnauthorized() *WeaviateGraphqlBatchUnauthorized {

	return &WeaviateGraphqlBatchUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateGraphqlBatchUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateGraphqlBatchForbiddenCode is the HTTP code returned for type WeaviateGraphqlBatchForbidden
const WeaviateGraphqlBatchForbiddenCode int = 403

/*WeaviateGraphqlBatchForbidden Forbidden

swagger:response weaviateGraphqlBatchForbidden
*/
type WeaviateGraphqlBatchForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateGraphqlBatchForbidden creates WeaviateGraphqlBatchForbidden with default headers values
func NewWeaviateGraphqlBatchForbidden() *WeaviateGraphqlBatchForbidden {

	return &WeaviateGraphqlBatchForbidden{}
}

// WithPayload adds the payload to the weaviate graphql batch forbidden response
func (o *WeaviateGraphqlBatchForbidden) WithPayload(payload *models.ErrorResponse) *WeaviateGraphqlBatchForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate graphql batch forbidden response
func (o *WeaviateGraphqlBatchForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateGraphqlBatchForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateGraphqlBatchUnprocessableEntityCode is the HTTP code returned for type WeaviateGraphqlBatchUnprocessableEntity
const WeaviateGraphqlBatchUnprocessableEntityCode int = 422

/*WeaviateGraphqlBatchUnprocessableEntity Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateGraphqlBatchUnprocessableEntity
*/
type WeaviateGraphqlBatchUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateGraphqlBatchUnprocessableEntity creates WeaviateGraphqlBatchUnprocessableEntity with default headers values
func NewWeaviateGraphqlBatchUnprocessableEntity() *WeaviateGraphqlBatchUnprocessableEntity {

	return &WeaviateGraphqlBatchUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate graphql batch unprocessable entity response
func (o *WeaviateGraphqlBatchUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateGraphqlBatchUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate graphql batch unprocessable entity response
func (o *WeaviateGraphqlBatchUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateGraphqlBatchUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateGraphqlBatchInternalServerErrorCode is the HTTP code returned for type WeaviateGraphqlBatchInternalServerError
const WeaviateGraphqlBatchInternalServerErrorCode int = 500

/*WeaviateGraphqlBatchInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response weaviateGraphqlBatchInternalServerError
*/
type WeaviateGraphqlBatchInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateGraphqlBatchInternalServerError creates WeaviateGraphqlBatchInternalServerError with default headers values
func NewWeaviateGraphqlBatchInternalServerError() *WeaviateGraphqlBatchInternalServerError {

	return &WeaviateGraphqlBatchInternalServerError{}
}

// WithPayload adds the payload to the weaviate graphql batch internal server error response
func (o *WeaviateGraphqlBatchInternalServerError) WithPayload(payload *models.ErrorResponse) *WeaviateGraphqlBatchInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate graphql batch internal server error response
func (o *WeaviateGraphqlBatchInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateGraphqlBatchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
