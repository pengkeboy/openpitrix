// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDescribeRepoTasksResponse openpitrix describe repo tasks response
// swagger:model openpitrixDescribeRepoTasksResponse
type OpenpitrixDescribeRepoTasksResponse struct {

	// repo task set
	RepoTaskSet OpenpitrixDescribeRepoTasksResponseRepoTaskSet `json:"repo_task_set"`

	// total count
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this openpitrix describe repo tasks response
func (m *OpenpitrixDescribeRepoTasksResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDescribeRepoTasksResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDescribeRepoTasksResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDescribeRepoTasksResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}