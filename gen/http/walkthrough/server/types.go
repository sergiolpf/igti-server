// Code generated by goa v3.1.1, DO NOT EDIT.
//
// walkthrough HTTP server types
//
// Command:
// $ goa gen guide.me/design

package server

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	walkthrough "guide.me/gen/walkthrough"
	walkthroughviews "guide.me/gen/walkthrough/views"
)

// AddRequestBody is the type of the "walkthrough" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of the Tutorial
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// base url for your tutorial to start from
	BaseURL *string `form:"baseURL,omitempty" json:"baseURL,omitempty" xml:"baseURL,omitempty"`
	// Status of the walkthrough [draft|published]
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Code to be added into an existing page to make it visible locally
	PublishedURL *string `form:"publishedURL,omitempty" json:"publishedURL,omitempty" xml:"publishedURL,omitempty"`
	// ID of the organization this tutorial belongs to
	Organization *string `form:"organization,omitempty" json:"organization,omitempty" xml:"organization,omitempty"`
}

// UpdateRequestBody is the type of the "walkthrough" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// ID is the unique id of the Walkthrough.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the Tutorial
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// base url for your tutorial to start from
	BaseURL *string `form:"baseURL,omitempty" json:"baseURL,omitempty" xml:"baseURL,omitempty"`
	// Status of the walkthrough [draft|published]
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Code to be added into an existing page to make it visible locally
	PublishedURL *string `form:"publishedURL,omitempty" json:"publishedURL,omitempty" xml:"publishedURL,omitempty"`
	// ID of the organization this tutorial belongs to
	Organization *string `form:"organization,omitempty" json:"organization,omitempty" xml:"organization,omitempty"`
}

// StoredWalkthroughResponseTinyCollection is the type of the "walkthrough"
// service "list" endpoint HTTP response body.
type StoredWalkthroughResponseTinyCollection []*StoredWalkthroughResponseTiny

// ShowResponseBody is the type of the "walkthrough" service "show" endpoint
// HTTP response body.
type ShowResponseBody struct {
	// ID is the unique id of the Walkthrough.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the Tutorial
	Name string `form:"name" json:"name" xml:"name"`
	// base url for your tutorial to start from
	BaseURL string `form:"baseURL" json:"baseURL" xml:"baseURL"`
	// Status of the walkthrough [draft|published]
	Status string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Code to be added into an existing page to make it visible locally
	PublishedURL *string `form:"publishedURL,omitempty" json:"publishedURL,omitempty" xml:"publishedURL,omitempty"`
	// ID of the organization this tutorial belongs to
	Organization string `form:"organization" json:"organization" xml:"organization"`
}

// ShowResponseBodyTiny is the type of the "walkthrough" service "show"
// endpoint HTTP response body.
type ShowResponseBodyTiny struct {
	// ID is the unique id of the Walkthrough.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the Tutorial
	Name string `form:"name" json:"name" xml:"name"`
	// base url for your tutorial to start from
	BaseURL string `form:"baseURL" json:"baseURL" xml:"baseURL"`
	// ID of the organization this tutorial belongs to
	Organization string `form:"organization" json:"organization" xml:"organization"`
}

// ShowNotFoundResponseBody is the type of the "walkthrough" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// ID of missing element
	ID string `form:"id" json:"id" xml:"id"`
}

// StoredWalkthroughResponseTiny is used to define fields on response body
// types.
type StoredWalkthroughResponseTiny struct {
	// ID is the unique id of the Walkthrough.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the Tutorial
	Name string `form:"name" json:"name" xml:"name"`
	// base url for your tutorial to start from
	BaseURL string `form:"baseURL" json:"baseURL" xml:"baseURL"`
	// ID of the organization this tutorial belongs to
	Organization string `form:"organization" json:"organization" xml:"organization"`
}

// NewStoredWalkthroughResponseTinyCollection builds the HTTP response body
// from the result of the "list" endpoint of the "walkthrough" service.
func NewStoredWalkthroughResponseTinyCollection(res walkthroughviews.StoredWalkthroughCollectionView) StoredWalkthroughResponseTinyCollection {
	body := make([]*StoredWalkthroughResponseTiny, len(res))
	for i, val := range res {
		body[i] = marshalWalkthroughviewsStoredWalkthroughViewToStoredWalkthroughResponseTiny(val)
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "walkthrough" service.
func NewShowResponseBody(res *walkthroughviews.StoredWalkthroughView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:           *res.ID,
		Name:         *res.Name,
		BaseURL:      *res.BaseURL,
		PublishedURL: res.PublishedURL,
		Organization: *res.Organization,
	}
	if res.Status != nil {
		body.Status = *res.Status
	}
	if res.Status == nil {
		body.Status = "draft"
	}
	return body
}

// NewShowResponseBodyTiny builds the HTTP response body from the result of the
// "show" endpoint of the "walkthrough" service.
func NewShowResponseBodyTiny(res *walkthroughviews.StoredWalkthroughView) *ShowResponseBodyTiny {
	body := &ShowResponseBodyTiny{
		ID:           *res.ID,
		Name:         *res.Name,
		BaseURL:      *res.BaseURL,
		Organization: *res.Organization,
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "walkthrough" service.
func NewShowNotFoundResponseBody(res *walkthrough.ElementNotFound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewListPayload builds a walkthrough service list endpoint payload.
func NewListPayload(id string) *walkthrough.ListPayload {
	v := &walkthrough.ListPayload{}
	v.ID = id

	return v
}

// NewShowPayload builds a walkthrough service show endpoint payload.
func NewShowPayload(id string, view *string) *walkthrough.ShowPayload {
	v := &walkthrough.ShowPayload{}
	v.ID = id
	v.View = view

	return v
}

// NewAddWalkthrough builds a walkthrough service add endpoint payload.
func NewAddWalkthrough(body *AddRequestBody) *walkthrough.Walkthrough {
	v := &walkthrough.Walkthrough{
		Name:         *body.Name,
		BaseURL:      *body.BaseURL,
		PublishedURL: body.PublishedURL,
		Organization: *body.Organization,
	}
	if body.Status != nil {
		v.Status = *body.Status
	}
	if body.Status == nil {
		v.Status = "draft"
	}

	return v
}

// NewRemovePayload builds a walkthrough service remove endpoint payload.
func NewRemovePayload(id string) *walkthrough.RemovePayload {
	v := &walkthrough.RemovePayload{}
	v.ID = id

	return v
}

// NewUpdateStoredWalkthrough builds a walkthrough service update endpoint
// payload.
func NewUpdateStoredWalkthrough(body *UpdateRequestBody) *walkthrough.StoredWalkthrough {
	v := &walkthrough.StoredWalkthrough{
		ID:           *body.ID,
		Name:         *body.Name,
		BaseURL:      *body.BaseURL,
		PublishedURL: body.PublishedURL,
		Organization: *body.Organization,
	}
	if body.Status != nil {
		v.Status = *body.Status
	}
	if body.Status == nil {
		v.Status = "draft"
	}

	return v
}

// NewPublishPayload builds a walkthrough service publish endpoint payload.
func NewPublishPayload(id string) *walkthrough.PublishPayload {
	v := &walkthrough.PublishPayload{}
	v.ID = id

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.BaseURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("baseURL", "body"))
	}
	if body.Organization == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("organization", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.BaseURL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.baseURL", *body.BaseURL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	if body.Status != nil {
		if !(*body.Status == "draft" || *body.Status == "completed" || *body.Status == "removed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"draft", "completed", "removed"}))
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.BaseURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("baseURL", "body"))
	}
	if body.Organization == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("organization", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.BaseURL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.baseURL", *body.BaseURL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	if body.Status != nil {
		if !(*body.Status == "draft" || *body.Status == "completed" || *body.Status == "removed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"draft", "completed", "removed"}))
		}
	}
	return
}
