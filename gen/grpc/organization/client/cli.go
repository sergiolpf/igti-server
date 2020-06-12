// Code generated by goa v3.1.1, DO NOT EDIT.
//
// organization gRPC client CLI support package
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"encoding/json"
	"fmt"

	goa "goa.design/goa/v3/pkg"
	organizationpb "guide.me/gen/grpc/organization/pb"
	organization "guide.me/gen/organization"
)

// BuildShowPayload builds the payload for the organization show endpoint from
// CLI flags.
func BuildShowPayload(organizationShowMessage string, organizationShowView string) (*organization.ShowPayload, error) {
	var err error
	var message organizationpb.ShowRequest
	{
		if organizationShowMessage != "" {
			err = json.Unmarshal([]byte(organizationShowMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Ipsum ratione eum.\"\n   }'")
			}
		}
	}
	var view *string
	{
		if organizationShowView != "" {
			view = &organizationShowView
			if view != nil {
				if !(*view == "default" || *view == "tiny") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &organization.ShowPayload{
		ID: message.Id,
	}
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the organization add endpoint from
// CLI flags.
func BuildAddPayload(organizationAddMessage string) (*organization.Organization, error) {
	var err error
	var message organizationpb.AddRequest
	{
		if organizationAddMessage != "" {
			err = json.Unmarshal([]byte(organizationAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"name\": \"Creating a new request in netflix!\",\n      \"url\": \"http://www.google.com/\"\n   }'")
			}
		}
	}
	v := &organization.Organization{
		Name: message.Name,
		URL:  message.Url,
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the organization remove endpoint
// from CLI flags.
func BuildRemovePayload(organizationRemoveMessage string) (*organization.RemovePayload, error) {
	var err error
	var message organizationpb.RemoveRequest
	{
		if organizationRemoveMessage != "" {
			err = json.Unmarshal([]byte(organizationRemoveMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Voluptates occaecati aliquid veniam ipsa in.\"\n   }'")
			}
		}
	}
	v := &organization.RemovePayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the organization update endpoint
// from CLI flags.
func BuildUpdatePayload(organizationUpdateMessage string) (*organization.StoredOrganization, error) {
	var err error
	var message organizationpb.UpdateRequest
	{
		if organizationUpdateMessage != "" {
			err = json.Unmarshal([]byte(organizationUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"123abc\",\n      \"name\": \"Creating a new request in netflix!\",\n      \"url\": \"http://www.google.com/\"\n   }'")
			}
		}
	}
	v := &organization.StoredOrganization{
		ID:   message.Id,
		Name: message.Name,
		URL:  message.Url,
	}

	return v, nil
}
