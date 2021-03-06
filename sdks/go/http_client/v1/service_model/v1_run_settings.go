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

package service_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1RunSettings v1 run settings
// swagger:model v1RunSettings
type V1RunSettings struct {

	// Agent
	Agent *V1RunSettingsCatalog `json:"agent,omitempty"`

	// K8S config maps
	ConfigResources []*V1RunSettingsCatalog `json:"config_resources"`

	// Connections
	Connections []*V1RunSettingsCatalog `json:"connections"`

	// Git Accesses
	GitAccesses []*V1RunSettingsCatalog `json:"git_accesses"`

	// Init connections
	InitConnections []*V1RunSettingsCatalog `json:"init_connections"`

	// Logs Store
	LogsStore *V1RunSettingsCatalog `json:"logs_store,omitempty"`

	// Namespace
	Namespace string `json:"namespace,omitempty"`

	// Outputs Store
	OutputsStore *V1RunSettingsCatalog `json:"outputs_store,omitempty"`

	// Queue
	Queue *V1RunSettingsCatalog `json:"queue,omitempty"`

	// Registry Access
	RegistryAccess *V1RunSettingsCatalog `json:"registry_access,omitempty"`
}

// Validate validates this v1 run settings
func (m *V1RunSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGitAccesses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogsStore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputsStore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RunSettings) validateAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSettings) validateConfigResources(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigResources) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigResources); i++ {
		if swag.IsZero(m.ConfigResources[i]) { // not required
			continue
		}

		if m.ConfigResources[i] != nil {
			if err := m.ConfigResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("config_resources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1RunSettings) validateConnections(formats strfmt.Registry) error {

	if swag.IsZero(m.Connections) { // not required
		return nil
	}

	for i := 0; i < len(m.Connections); i++ {
		if swag.IsZero(m.Connections[i]) { // not required
			continue
		}

		if m.Connections[i] != nil {
			if err := m.Connections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1RunSettings) validateGitAccesses(formats strfmt.Registry) error {

	if swag.IsZero(m.GitAccesses) { // not required
		return nil
	}

	for i := 0; i < len(m.GitAccesses); i++ {
		if swag.IsZero(m.GitAccesses[i]) { // not required
			continue
		}

		if m.GitAccesses[i] != nil {
			if err := m.GitAccesses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("git_accesses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1RunSettings) validateInitConnections(formats strfmt.Registry) error {

	if swag.IsZero(m.InitConnections) { // not required
		return nil
	}

	for i := 0; i < len(m.InitConnections); i++ {
		if swag.IsZero(m.InitConnections[i]) { // not required
			continue
		}

		if m.InitConnections[i] != nil {
			if err := m.InitConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("init_connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1RunSettings) validateLogsStore(formats strfmt.Registry) error {

	if swag.IsZero(m.LogsStore) { // not required
		return nil
	}

	if m.LogsStore != nil {
		if err := m.LogsStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logs_store")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSettings) validateOutputsStore(formats strfmt.Registry) error {

	if swag.IsZero(m.OutputsStore) { // not required
		return nil
	}

	if m.OutputsStore != nil {
		if err := m.OutputsStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outputs_store")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSettings) validateQueue(formats strfmt.Registry) error {

	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

func (m *V1RunSettings) validateRegistryAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistryAccess) { // not required
		return nil
	}

	if m.RegistryAccess != nil {
		if err := m.RegistryAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry_access")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RunSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RunSettings) UnmarshalBinary(b []byte) error {
	var res V1RunSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
