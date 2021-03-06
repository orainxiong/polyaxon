// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CopyRunReader is a Reader for the CopyRun structure.
type CopyRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCopyRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCopyRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCopyRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCopyRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCopyRunOK creates a CopyRunOK with default headers values
func NewCopyRunOK() *CopyRunOK {
	return &CopyRunOK{}
}

/*CopyRunOK handles this case with default header values.

A successful response.
*/
type CopyRunOK struct {
	Payload *service_model.V1Run
}

func (o *CopyRunOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{entity.owner}/{entity.project}/runs/{entity.uuid}/copy][%d] copyRunOK  %+v", 200, o.Payload)
}

func (o *CopyRunOK) GetPayload() *service_model.V1Run {
	return o.Payload
}

func (o *CopyRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Run)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyRunNoContent creates a CopyRunNoContent with default headers values
func NewCopyRunNoContent() *CopyRunNoContent {
	return &CopyRunNoContent{}
}

/*CopyRunNoContent handles this case with default header values.

No content.
*/
type CopyRunNoContent struct {
	Payload interface{}
}

func (o *CopyRunNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{entity.owner}/{entity.project}/runs/{entity.uuid}/copy][%d] copyRunNoContent  %+v", 204, o.Payload)
}

func (o *CopyRunNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CopyRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyRunForbidden creates a CopyRunForbidden with default headers values
func NewCopyRunForbidden() *CopyRunForbidden {
	return &CopyRunForbidden{}
}

/*CopyRunForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CopyRunForbidden struct {
	Payload interface{}
}

func (o *CopyRunForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{entity.owner}/{entity.project}/runs/{entity.uuid}/copy][%d] copyRunForbidden  %+v", 403, o.Payload)
}

func (o *CopyRunForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CopyRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCopyRunNotFound creates a CopyRunNotFound with default headers values
func NewCopyRunNotFound() *CopyRunNotFound {
	return &CopyRunNotFound{}
}

/*CopyRunNotFound handles this case with default header values.

Resource does not exist.
*/
type CopyRunNotFound struct {
	Payload interface{}
}

func (o *CopyRunNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{entity.owner}/{entity.project}/runs/{entity.uuid}/copy][%d] copyRunNotFound  %+v", 404, o.Payload)
}

func (o *CopyRunNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CopyRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
