// Code generated by goa v3.1.1, DO NOT EDIT.
//
// walkthrough HTTP client types
//
// Command:
// $ goa gen guide.me/design

package client

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

// UpdateRequestBody is the type of the "walkthrough" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
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

// RenameRequestBody is the type of the "walkthrough" service "rename" endpoint
// HTTP request body.
type RenameRequestBody struct {
	// ID of Walkthrough to be renamed
	ID string `form:"id" json:"id" xml:"id"`
	// New Name to the walkthrough
	Name string `form:"name" json:"name" xml:"name"`
}

// ListResponseBody is the type of the "walkthrough" service "list" endpoint
// HTTP response body.
type ListResponseBody []*StoredWalkthroughResponse

// ShowResponseBody is the type of the "walkthrough" service "show" endpoint
// HTTP response body.
type ShowResponseBody struct {
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

// AddResponseBody is the type of the "walkthrough" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
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

// RenameResponseBody is the type of the "walkthrough" service "rename"
// endpoint HTTP response body.
type RenameResponseBody struct {
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

// ShowNotFoundResponseBody is the type of the "walkthrough" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing element
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// StoredWalkthroughResponse is used to define fields on response body types.
type StoredWalkthroughResponse struct {
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

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "walkthrough" service.
func NewAddRequestBody(p *walkthrough.Walkthrough) *AddRequestBody {
	body := &AddRequestBody{
		Name:         p.Name,
		BaseURL:      p.BaseURL,
		Status:       p.Status,
		PublishedURL: p.PublishedURL,
		Organization: p.Organization,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "walkthrough" service.
func NewUpdateRequestBody(p *walkthrough.StoredWalkthrough) *UpdateRequestBody {
	body := &UpdateRequestBody{
		ID:           p.ID,
		Name:         p.Name,
		BaseURL:      p.BaseURL,
		Status:       p.Status,
		PublishedURL: p.PublishedURL,
		Organization: p.Organization,
	}
	return body
}

// NewRenameRequestBody builds the HTTP request body from the payload of the
// "rename" endpoint of the "walkthrough" service.
func NewRenameRequestBody(p *walkthrough.RenamePayload) *RenameRequestBody {
	body := &RenameRequestBody{
		ID:   p.ID,
		Name: p.Name,
	}
	return body
}

// NewListStoredWalkthroughCollectionOK builds a "walkthrough" service "list"
// endpoint result from a HTTP "OK" response.
func NewListStoredWalkthroughCollectionOK(body ListResponseBody) walkthroughviews.StoredWalkthroughCollectionView {
	v := make([]*walkthroughviews.StoredWalkthroughView, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredWalkthroughResponseToWalkthroughviewsStoredWalkthroughView(val)
	}
	return v
}

// NewShowStoredWalkthroughOK builds a "walkthrough" service "show" endpoint
// result from a HTTP "OK" response.
func NewShowStoredWalkthroughOK(body *ShowResponseBody) *walkthroughviews.StoredWalkthroughView {
	v := &walkthroughviews.StoredWalkthroughView{
		ID:           body.ID,
		Name:         body.Name,
		BaseURL:      body.BaseURL,
		Status:       body.Status,
		PublishedURL: body.PublishedURL,
		Organization: body.Organization,
	}
	if body.Status == nil {
		var tmp string = "draft"
		v.Status = &tmp
	}

	return v
}

// NewShowNotFound builds a walkthrough service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *walkthrough.ElementNotFound {
	v := &walkthrough.ElementNotFound{
		Message: *body.Message,
		ID:      *body.ID,
	}

	return v
}

// NewAddStoredWalkthroughCreated builds a "walkthrough" service "add" endpoint
// result from a HTTP "Created" response.
func NewAddStoredWalkthroughCreated(body *AddResponseBody) *walkthroughviews.StoredWalkthroughView {
	v := &walkthroughviews.StoredWalkthroughView{
		ID:           body.ID,
		Name:         body.Name,
		BaseURL:      body.BaseURL,
		Status:       body.Status,
		PublishedURL: body.PublishedURL,
		Organization: body.Organization,
	}
	if body.Status == nil {
		var tmp string = "draft"
		v.Status = &tmp
	}

	return v
}

// NewRenameStoredWalkthroughOK builds a "walkthrough" service "rename"
// endpoint result from a HTTP "OK" response.
func NewRenameStoredWalkthroughOK(body *RenameResponseBody) *walkthroughviews.StoredWalkthroughView {
	v := &walkthroughviews.StoredWalkthroughView{
		ID:           body.ID,
		Name:         body.Name,
		BaseURL:      body.BaseURL,
		Status:       body.Status,
		PublishedURL: body.PublishedURL,
		Organization: body.Organization,
	}
	if body.Status == nil {
		var tmp string = "draft"
		v.Status = &tmp
	}

	return v
}

// ValidateShowNotFoundResponseBody runs the validations defined on
// show_not_found_response_body
func ValidateShowNotFoundResponseBody(body *ShowNotFoundResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateStoredWalkthroughResponse runs the validations defined on
// StoredWalkthroughResponse
func ValidateStoredWalkthroughResponse(body *StoredWalkthroughResponse) (err error) {
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
