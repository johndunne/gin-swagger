// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterStatusProblemsItems cluster status problems items
// swagger:model clusterStatusProblemsItems
type ClusterStatusProblemsItems struct {

	// A human-readable explanation specific to this occurrence of
	// the problem.
	//
	Detail string `json:"detail,omitempty"`

	// A URI reference that identifies the specific occurrence of
	// the problem.
	//
	Instance string `json:"instance,omitempty"`

	// The HTTP status code generated by the origin server for this
	// occurence of the problem.
	//
	Status int32 `json:"status,omitempty"`

	// A short, human-readable summary of the problem type.
	//
	// Required: true
	Title *string `json:"title"`

	// A URI reference the indentifies the problem type.
	// Required: true
	Type *string `json:"type"`

	// cluster status problems items
	ClusterStatusProblemsItems map[string]string `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *ClusterStatusProblemsItems) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// A human-readable explanation specific to this occurrence of
		// the problem.
		//
		Detail string `json:"detail,omitempty"`

		// A URI reference that identifies the specific occurrence of
		// the problem.
		//
		Instance string `json:"instance,omitempty"`

		// The HTTP status code generated by the origin server for this
		// occurence of the problem.
		//
		Status int32 `json:"status,omitempty"`

		// A short, human-readable summary of the problem type.
		//
		// Required: true
		Title *string `json:"title"`

		// A URI reference the indentifies the problem type.
		// Required: true
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv ClusterStatusProblemsItems

	rcv.Detail = stage1.Detail

	rcv.Instance = stage1.Instance

	rcv.Status = stage1.Status

	rcv.Title = stage1.Title

	rcv.Type = stage1.Type

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "detail")

	delete(stage2, "instance")

	delete(stage2, "status")

	delete(stage2, "title")

	delete(stage2, "type")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]string)
		for k, v := range stage2 {
			var toadd string
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.ClusterStatusProblemsItems = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m ClusterStatusProblemsItems) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// A human-readable explanation specific to this occurrence of
		// the problem.
		//
		Detail string `json:"detail,omitempty"`

		// A URI reference that identifies the specific occurrence of
		// the problem.
		//
		Instance string `json:"instance,omitempty"`

		// The HTTP status code generated by the origin server for this
		// occurence of the problem.
		//
		Status int32 `json:"status,omitempty"`

		// A short, human-readable summary of the problem type.
		//
		// Required: true
		Title *string `json:"title"`

		// A URI reference the indentifies the problem type.
		// Required: true
		Type *string `json:"type"`
	}

	stage1.Detail = m.Detail

	stage1.Instance = m.Instance

	stage1.Status = m.Status

	stage1.Title = m.Title

	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.ClusterStatusProblemsItems) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.ClusterStatusProblemsItems)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this cluster status problems items
func (m *ClusterStatusProblemsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStatusProblemsItems) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *ClusterStatusProblemsItems) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatusProblemsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatusProblemsItems) UnmarshalBinary(b []byte) error {
	var res ClusterStatusProblemsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
