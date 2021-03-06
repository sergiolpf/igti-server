// Code generated by goa v3.1.1, DO NOT EDIT.
//
// organization HTTP client CLI support package
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	organization "guide.me/gen/organization"
)

// BuildShowPayload builds the payload for the organization show endpoint from
// CLI flags.
func BuildShowPayload(organizationShowID string, organizationShowView string) (*organization.ShowPayload, error) {
	var err error
	var id string
	{
		id = organizationShowID
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
	v := &organization.ShowPayload{}
	v.ID = id
	v.View = view

	return v, nil
}

// BuildAddPayload builds the payload for the organization add endpoint from
// CLI flags.
func BuildAddPayload(organizationAddBody string) (*organization.Organization, error) {
	var err error
	var body AddRequestBody
	{
		err = json.Unmarshal([]byte(organizationAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"name\": \"Creating a new request in netflix!\",\n      \"url\": \"http://www.google.com/\"\n   }'")
		}
		if utf8.RuneCountInString(body.Name) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 200, false))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", body.URL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
		if err != nil {
			return nil, err
		}
	}
	v := &organization.Organization{
		Name: body.Name,
		URL:  body.URL,
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the organization remove endpoint
// from CLI flags.
func BuildRemovePayload(organizationRemoveID string) (*organization.RemovePayload, error) {
	var id string
	{
		id = organizationRemoveID
	}
	v := &organization.RemovePayload{}
	v.ID = id

	return v, nil
}

// BuildUpdatePayload builds the payload for the organization update endpoint
// from CLI flags.
func BuildUpdatePayload(organizationUpdateBody string) (*organization.StoredOrganization, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(organizationUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"id\": \"123abc\",\n      \"name\": \"Creating a new request in netflix!\",\n      \"url\": \"http://www.google.com/\"\n   }'")
		}
		if utf8.RuneCountInString(body.Name) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", body.Name, utf8.RuneCountInString(body.Name), 200, false))
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", body.URL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
		if err != nil {
			return nil, err
		}
	}
	v := &organization.StoredOrganization{
		ID:   body.ID,
		Name: body.Name,
		URL:  body.URL,
	}

	return v, nil
}
