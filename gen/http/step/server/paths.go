// Code generated by goa v3.1.1, DO NOT EDIT.
//
// HTTP request path constructors for the step service.
//
// Command:
// $ goa gen guide.me/design

package server

import (
	"fmt"
)

// ListStepPath returns the URL path to the step service list HTTP endpoint.
func ListStepPath(id string) string {
	return fmt.Sprintf("/steps/%v", id)
}

// AddStepPath returns the URL path to the step service add HTTP endpoint.
func AddStepPath() string {
	return "/steps"
}

// RemoveStepPath returns the URL path to the step service remove HTTP endpoint.
func RemoveStepPath() string {
	return "/steps"
}

// UpdateStepPath returns the URL path to the step service update HTTP endpoint.
func UpdateStepPath() string {
	return "/steps/update"
}
