// Code generated by goa v3.1.1, DO NOT EDIT.
//
// HTTP request path constructors for the walkthrough service.
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"fmt"
)

// ListWalkthroughPath returns the URL path to the walkthrough service list HTTP endpoint.
func ListWalkthroughPath(id string) string {
	return fmt.Sprintf("/walkthrough/%v", id)
}

// ShowWalkthroughPath returns the URL path to the walkthrough service show HTTP endpoint.
func ShowWalkthroughPath(id string) string {
	return fmt.Sprintf("/walkthrough/%v", id)
}

// AddWalkthroughPath returns the URL path to the walkthrough service add HTTP endpoint.
func AddWalkthroughPath() string {
	return "/walkthrough"
}

// RemoveWalkthroughPath returns the URL path to the walkthrough service remove HTTP endpoint.
func RemoveWalkthroughPath(id string) string {
	return fmt.Sprintf("/walkthrough/%v", id)
}

// UpdateWalkthroughPath returns the URL path to the walkthrough service update HTTP endpoint.
func UpdateWalkthroughPath() string {
	return "/walkthrough/update"
}
