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

package teams_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateTeamMemberReader is a Reader for the UpdateTeamMember structure.
type UpdateTeamMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTeamMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTeamMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateTeamMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateTeamMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTeamMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateTeamMemberOK creates a UpdateTeamMemberOK with default headers values
func NewUpdateTeamMemberOK() *UpdateTeamMemberOK {
	return &UpdateTeamMemberOK{}
}

/*UpdateTeamMemberOK handles this case with default header values.

A successful response.
*/
type UpdateTeamMemberOK struct {
	Payload *service_model.V1TeamMember
}

func (o *UpdateTeamMemberOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamMemberOK) GetPayload() *service_model.V1TeamMember {
	return o.Payload
}

func (o *UpdateTeamMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1TeamMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberNoContent creates a UpdateTeamMemberNoContent with default headers values
func NewUpdateTeamMemberNoContent() *UpdateTeamMemberNoContent {
	return &UpdateTeamMemberNoContent{}
}

/*UpdateTeamMemberNoContent handles this case with default header values.

No content.
*/
type UpdateTeamMemberNoContent struct {
	Payload interface{}
}

func (o *UpdateTeamMemberNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberNoContent  %+v", 204, o.Payload)
}

func (o *UpdateTeamMemberNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTeamMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberForbidden creates a UpdateTeamMemberForbidden with default headers values
func NewUpdateTeamMemberForbidden() *UpdateTeamMemberForbidden {
	return &UpdateTeamMemberForbidden{}
}

/*UpdateTeamMemberForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdateTeamMemberForbidden struct {
	Payload interface{}
}

func (o *UpdateTeamMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTeamMemberForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTeamMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamMemberNotFound creates a UpdateTeamMemberNotFound with default headers values
func NewUpdateTeamMemberNotFound() *UpdateTeamMemberNotFound {
	return &UpdateTeamMemberNotFound{}
}

/*UpdateTeamMemberNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdateTeamMemberNotFound struct {
	Payload interface{}
}

func (o *UpdateTeamMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/teams/{team}/members/{member.user}][%d] updateTeamMemberNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTeamMemberNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateTeamMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
