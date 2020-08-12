// Code generated by goa v3.1.1, DO NOT EDIT.
//
// walkthrough HTTP client CLI support package
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	walkthrough "guide.me/gen/walkthrough"
)

// BuildListPayload builds the payload for the walkthrough list endpoint from
// CLI flags.
func BuildListPayload(walkthroughListID string) (*walkthrough.ListPayload, error) {
	var id string
	{
		id = walkthroughListID
	}
	v := &walkthrough.ListPayload{}
	v.ID = id

	return v, nil
}

// BuildShowPayload builds the payload for the walkthrough show endpoint from
// CLI flags.
func BuildShowPayload(walkthroughShowID string, walkthroughShowView string) (*walkthrough.ShowPayload, error) {
	var err error
	var id string
	{
		id = walkthroughShowID
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
	v := &walkthrough.ShowPayload{}
	v.ID = id
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the walkthrough add endpoint from CLI
// flags.
func BuildAddPayload(walkthroughAddBody string) (*walkthrough.Walkthrough, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(walkthroughAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"baseURL\": \"http://www.google.com/\",\n      \"name\": \"How to create a new process using the exception condition.\",\n      \"organization\": \"Et non quis neque sit non.\",\n      \"publishedURL\": \"Quas aut quia dolores.\",\n      \"status\": \"draft | published\"\n   }'")
		}
		if utf8.RuneCountInString(body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 100, false))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.baseURL", body.BaseURL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
		if !(body.Status == "draft" || body.Status == "completed" || body.Status == "removed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", body.Status, []interface{}{"draft", "completed", "removed"}))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &walkthrough.Walkthrough{
		Name:         body.Name,
		BaseURL:      body.BaseURL,
		Status:       body.Status,
		PublishedURL: body.PublishedURL,
		Organization: body.Organization,
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the walkthrough remove endpoint
// from CLI flags.
func BuildRemovePayload(walkthroughRemoveID string) (*walkthrough.RemovePayload, error) {
	var id string
	{
		id = walkthroughRemoveID
	}
	v := &walkthrough.RemovePayload{}
	v.ID = id

	return v, nil
}

// BuildUpdatePayload builds the payload for the walkthrough update endpoint
// from CLI flags.
func BuildUpdatePayload(walkthroughUpdateBody string) (*walkthrough.StoredWalkthrough, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(walkthroughUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"baseURL\": \"http://www.google.com/\",\n      \"id\": \"123abc\",\n      \"name\": \"How to create a new process using the exception condition.\",\n      \"organization\": \"Quod repudiandae sed dolor optio et architecto.\",\n      \"publishedURL\": \"Maxime ea qui molestiae adipisci sint aut.\",\n      \"status\": \"draft | published\"\n   }'")
		}
		if utf8.RuneCountInString(body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 100, false))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.baseURL", body.BaseURL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
		if !(body.Status == "draft" || body.Status == "completed" || body.Status == "removed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", body.Status, []interface{}{"draft", "completed", "removed"}))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &walkthrough.StoredWalkthrough{
		ID:           body.ID,
		Name:         body.Name,
		BaseURL:      body.BaseURL,
		Status:       body.Status,
		PublishedURL: body.PublishedURL,
		Organization: body.Organization,
	}

	return v, nil
}

// BuildRenamePayload builds the payload for the walkthrough rename endpoint
// from CLI flags.
func BuildRenamePayload(walkthroughRenameBody string) (*walkthrough.RenamePayload, error) {
	var err error
	var body RenameRequestBody
	{
		err = json.Unmarshal([]byte(walkthroughRenameBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"id\": \"Ab accusamus.\",\n      \"name\": \"Veniam cupiditate quia aperiam illo enim.\"\n   }'")
		}
	}
	v := &walkthrough.RenamePayload{
		ID:   body.ID,
		Name: body.Name,
	}

	return v, nil
}

// BuildPublishPayload builds the payload for the walkthrough publish endpoint
// from CLI flags.
func BuildPublishPayload(walkthroughPublishID string) (*walkthrough.PublishPayload, error) {
	var id string
	{
		id = walkthroughPublishID
	}
	v := &walkthrough.PublishPayload{}
	v.ID = id

	return v, nil
}
