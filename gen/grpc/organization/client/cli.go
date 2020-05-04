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
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Velit est culpa.\"\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"name\": \"Blue\\'s Cuvee\",\n      \"url\": \"http://www.google.com/\"\n   }'")
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
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Rem delectus.\"\n   }'")
			}
		}
	}
	v := &organization.RemovePayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildMultiAddPayload builds the payload for the organization multi_add
// endpoint from CLI flags.
func BuildMultiAddPayload(organizationMultiAddMessage string) ([]*organization.Organization, error) {
	var err error
	var message organizationpb.MultiAddRequest
	{
		if organizationMultiAddMessage != "" {
			err = json.Unmarshal([]byte(organizationMultiAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"field\": [\n         {\n            \"name\": \"Blue\\'s Cuvee\",\n            \"url\": \"http://www.google.com/\"\n         },\n         {\n            \"name\": \"Blue\\'s Cuvee\",\n            \"url\": \"http://www.google.com/\"\n         },\n         {\n            \"name\": \"Blue\\'s Cuvee\",\n            \"url\": \"http://www.google.com/\"\n         },\n         {\n            \"name\": \"Blue\\'s Cuvee\",\n            \"url\": \"http://www.google.com/\"\n         }\n      ]\n   }'")
			}
		}
	}
	v := make([]*organization.Organization, len(message.Field))
	for i, val := range message.Field {
		v[i] = &organization.Organization{
			Name: val.Name,
			URL:  val.Url,
		}
	}
	return v, nil
}

// BuildMultiUpdatePayload builds the payload for the organization multi_update
// endpoint from CLI flags.
func BuildMultiUpdatePayload(organizationMultiUpdateMessage string) (*organization.MultiUpdatePayload, error) {
	var err error
	var message organizationpb.MultiUpdateRequest
	{
		if organizationMultiUpdateMessage != "" {
			err = json.Unmarshal([]byte(organizationMultiUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"ids\": [\n         \"Recusandae corporis fugit non.\",\n         \"Asperiores maxime.\"\n      ],\n      \"organizations\": [\n         {\n            \"name\": \"Blue\\'s Cuvee\",\n            \"url\": \"http://www.google.com/\"\n         },\n         {\n            \"name\": \"Blue\\'s Cuvee\",\n            \"url\": \"http://www.google.com/\"\n         },\n         {\n            \"name\": \"Blue\\'s Cuvee\",\n            \"url\": \"http://www.google.com/\"\n         },\n         {\n            \"name\": \"Blue\\'s Cuvee\",\n            \"url\": \"http://www.google.com/\"\n         }\n      ]\n   }'")
			}
		}
	}
	v := &organization.MultiUpdatePayload{}
	if message.Ids != nil {
		v.Ids = make([]string, len(message.Ids))
		for i, val := range message.Ids {
			v.Ids[i] = val
		}
	}
	if message.Organizations != nil {
		v.Organizations = make([]*organization.Organization, len(message.Organizations))
		for i, val := range message.Organizations {
			v.Organizations[i] = &organization.Organization{
				Name: val.Name,
				URL:  val.Url,
			}
		}
	}

	return v, nil
}
