// Code generated by goa v3.1.1, DO NOT EDIT.
//
// organization HTTP client types
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	organization "guide.me/gen/organization"
	organizationviews "guide.me/gen/organization/views"
)

// AddRequestBody is the type of the "organization" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// Name of Organization
	Name string `form:"name" json:"name" xml:"name"`
	// Company website URL
	URL string `form:"url" json:"url" xml:"url"`
}

// UpdateRequestBody is the type of the "organization" service "update"
// endpoint HTTP request body.
type UpdateRequestBody struct {
	// ID is the unique id of the Organization.
	ID string `form:"id" json:"id" xml:"id"`
	// Name of Organization
	Name string `form:"name" json:"name" xml:"name"`
	// Company website URL
	URL string `form:"url" json:"url" xml:"url"`
}

// ListResponseBody is the type of the "organization" service "list" endpoint
// HTTP response body.
type ListResponseBody []*StoredOrganizationResponse

// ShowResponseBody is the type of the "organization" service "show" endpoint
// HTTP response body.
type ShowResponseBody struct {
	// ID is the unique id of the Organization.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of Organization
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Company website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "organization" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// ID of missing element
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// StoredOrganizationResponse is used to define fields on response body types.
type StoredOrganizationResponse struct {
	// ID is the unique id of the Organization.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of Organization
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Company website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "organization" service.
func NewAddRequestBody(p *organization.Organization) *AddRequestBody {
	body := &AddRequestBody{
		Name: p.Name,
		URL:  p.URL,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "organization" service.
func NewUpdateRequestBody(p *organization.StoredOrganization) *UpdateRequestBody {
	body := &UpdateRequestBody{
		ID:   p.ID,
		Name: p.Name,
		URL:  p.URL,
	}
	return body
}

// NewListStoredOrganizationCollectionOK builds a "organization" service "list"
// endpoint result from a HTTP "OK" response.
func NewListStoredOrganizationCollectionOK(body ListResponseBody) organizationviews.StoredOrganizationCollectionView {
	v := make([]*organizationviews.StoredOrganizationView, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredOrganizationResponseToOrganizationviewsStoredOrganizationView(val)
	}
	return v
}

// NewShowStoredOrganizationOK builds a "organization" service "show" endpoint
// result from a HTTP "OK" response.
func NewShowStoredOrganizationOK(body *ShowResponseBody) *organizationviews.StoredOrganizationView {
	v := &organizationviews.StoredOrganizationView{
		ID:   body.ID,
		Name: body.Name,
		URL:  body.URL,
	}

	return v
}

// NewShowNotFound builds a organization service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *organization.NotFound {
	v := &organization.NotFound{
		Message: *body.Message,
		ID:      *body.ID,
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

// ValidateStoredOrganizationResponse runs the validations defined on
// StoredOrganizationResponse
func ValidateStoredOrganizationResponse(body *StoredOrganizationResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 200 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 200, false))
		}
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}
