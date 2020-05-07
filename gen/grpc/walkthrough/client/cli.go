// Code generated by goa v3.1.1, DO NOT EDIT.
//
// walkthrough gRPC client CLI support package
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"encoding/json"
	"fmt"

	goa "goa.design/goa/v3/pkg"
	walkthroughpb "guide.me/gen/grpc/walkthrough/pb"
	walkthrough "guide.me/gen/walkthrough"
)

// BuildListPayload builds the payload for the walkthrough list endpoint from
// CLI flags.
func BuildListPayload(walkthroughListMessage string) (*walkthrough.ListPayload, error) {
	var err error
	var message walkthroughpb.ListRequest
	{
		if walkthroughListMessage != "" {
			err = json.Unmarshal([]byte(walkthroughListMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Architecto omnis ab accusamus est veniam.\"\n   }'")
			}
		}
	}
	v := &walkthrough.ListPayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildShowPayload builds the payload for the walkthrough show endpoint from
// CLI flags.
func BuildShowPayload(walkthroughShowMessage string, walkthroughShowView string) (*walkthrough.ShowPayload, error) {
	var err error
	var message walkthroughpb.ShowRequest
	{
		if walkthroughShowMessage != "" {
			err = json.Unmarshal([]byte(walkthroughShowMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Repudiandae corporis harum.\"\n   }'")
			}
		}
	}
	var view *string
	{
		if walkthroughShowView != "" {
			view = &walkthroughShowView
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
	v := &walkthrough.ShowPayload{
		ID: message.Id,
	}
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the walkthrough add endpoint from CLI
// flags.
func BuildAddPayload(walkthroughAddMessage string) (*walkthrough.Walkthrough, error) {
	var err error
	var message walkthroughpb.AddRequest
	{
		if walkthroughAddMessage != "" {
			err = json.Unmarshal([]byte(walkthroughAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"baseURL\": \"http://www.google.com/\",\n      \"name\": \"How to create a new process using the exception condition.\",\n      \"organization\": \"Ipsa ipsum ratione eum tenetur.\",\n      \"publishedURL\": \"Quia aspernatur est hic esse aspernatur.\",\n      \"status\": \"draft | published\"\n   }'")
			}
		}
	}
	v := &walkthrough.Walkthrough{
		Name:         message.Name,
		BaseURL:      message.BaseUrl,
		Status:       message.Status,
		Organization: message.Organization,
	}
	if message.PublishedUrl != "" {
		v.PublishedURL = &message.PublishedUrl
	}
	if message.Status == "" {
		v.Status = "draft"
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the walkthrough remove endpoint
// from CLI flags.
func BuildRemovePayload(walkthroughRemoveMessage string) (*walkthrough.RemovePayload, error) {
	var err error
	var message walkthroughpb.RemoveRequest
	{
		if walkthroughRemoveMessage != "" {
			err = json.Unmarshal([]byte(walkthroughRemoveMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Ipsa in ipsum.\"\n   }'")
			}
		}
	}
	v := &walkthrough.RemovePayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the walkthrough update endpoint
// from CLI flags.
func BuildUpdatePayload(walkthroughUpdateMessage string) (*walkthrough.StoredWalkthrough, error) {
	var err error
	var message walkthroughpb.UpdateRequest
	{
		if walkthroughUpdateMessage != "" {
			err = json.Unmarshal([]byte(walkthroughUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"baseURL\": \"http://www.google.com/\",\n      \"id\": \"123abc\",\n      \"name\": \"How to create a new process using the exception condition.\",\n      \"organization\": \"Deleniti qui.\",\n      \"publishedURL\": \"Fugiat et delectus quo quo animi illum.\",\n      \"status\": \"draft | published\"\n   }'")
			}
		}
	}
	v := &walkthrough.StoredWalkthrough{
		ID:           message.Id,
		Name:         message.Name,
		BaseURL:      message.BaseUrl,
		Status:       message.Status,
		Organization: message.Organization,
	}
	if message.PublishedUrl != "" {
		v.PublishedURL = &message.PublishedUrl
	}
	if message.Status == "" {
		v.Status = "draft"
	}

	return v, nil
}

// BuildPublishPayload builds the payload for the walkthrough publish endpoint
// from CLI flags.
func BuildPublishPayload(walkthroughPublishMessage string) (*walkthrough.PublishPayload, error) {
	var err error
	var message walkthroughpb.PublishRequest
	{
		if walkthroughPublishMessage != "" {
			err = json.Unmarshal([]byte(walkthroughPublishMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Eum id vero iste voluptas facere.\"\n   }'")
			}
		}
	}
	v := &walkthrough.PublishPayload{
		ID: message.Id,
	}

	return v, nil
}
