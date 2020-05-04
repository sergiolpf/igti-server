// Code generated by goa v3.1.1, DO NOT EDIT.
//
// organization gRPC server types
//
// Command:
// $ goa gen guide.me/design

package server

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	organizationpb "guide.me/gen/grpc/organization/pb"
	organization "guide.me/gen/organization"
	organizationviews "guide.me/gen/organization/views"
)

// NewStoredOrganizationCollection builds the gRPC response type from the
// result of the "list" endpoint of the "organization" service.
func NewStoredOrganizationCollection(result organizationviews.StoredOrganizationCollectionView) *organizationpb.StoredOrganizationCollection {
	message := &organizationpb.StoredOrganizationCollection{}
	message.Field = make([]*organizationpb.StoredOrganization, len(result))
	for i, val := range result {
		message.Field[i] = &organizationpb.StoredOrganization{}
		if val.ID != nil {
			message.Field[i].Id = *val.ID
		}
		if val.Name != nil {
			message.Field[i].Name = *val.Name
		}
		if val.URL != nil {
			message.Field[i].Url = *val.URL
		}
	}
	return message
}

// NewShowPayload builds the payload of the "show" endpoint of the
// "organization" service from the gRPC request type.
func NewShowPayload(message *organizationpb.ShowRequest, view *string) *organization.ShowPayload {
	v := &organization.ShowPayload{
		ID: message.Id,
	}
	v.View = view
	return v
}

// NewShowResponse builds the gRPC response type from the result of the "show"
// endpoint of the "organization" service.
func NewShowResponse(result *organizationviews.StoredOrganizationView) *organizationpb.ShowResponse {
	message := &organizationpb.ShowResponse{}
	if result.ID != nil {
		message.Id = *result.ID
	}
	if result.Name != nil {
		message.Name = *result.Name
	}
	if result.URL != nil {
		message.Url = *result.URL
	}
	return message
}

// NewShowNotFoundError builds the gRPC error response type from the error of
// the "show" endpoint of the "organization" service.
func NewShowNotFoundError(er *organization.OrgNotFound) *organizationpb.ShowNotFoundError {
	message := &organizationpb.ShowNotFoundError{
		Message_: er.Message,
		Id:       er.ID,
	}
	return message
}

// NewAddPayload builds the payload of the "add" endpoint of the "organization"
// service from the gRPC request type.
func NewAddPayload(message *organizationpb.AddRequest) *organization.Organization {
	v := &organization.Organization{
		Name: message.Name,
		URL:  message.Url,
	}
	return v
}

// NewAddResponse builds the gRPC response type from the result of the "add"
// endpoint of the "organization" service.
func NewAddResponse(result string) *organizationpb.AddResponse {
	message := &organizationpb.AddResponse{}
	message.Field = result
	return message
}

// NewRemovePayload builds the payload of the "remove" endpoint of the
// "organization" service from the gRPC request type.
func NewRemovePayload(message *organizationpb.RemoveRequest) *organization.RemovePayload {
	v := &organization.RemovePayload{
		ID: message.Id,
	}
	return v
}

// NewRemoveResponse builds the gRPC response type from the result of the
// "remove" endpoint of the "organization" service.
func NewRemoveResponse() *organizationpb.RemoveResponse {
	message := &organizationpb.RemoveResponse{}
	return message
}

// NewMultiAddPayload builds the payload of the "multi_add" endpoint of the
// "organization" service from the gRPC request type.
func NewMultiAddPayload(message *organizationpb.MultiAddRequest) []*organization.Organization {
	v := make([]*organization.Organization, len(message.Field))
	for i, val := range message.Field {
		v[i] = &organization.Organization{
			Name: val.Name,
			URL:  val.Url,
		}
	}
	return v
}

// NewMultiAddResponse builds the gRPC response type from the result of the
// "multi_add" endpoint of the "organization" service.
func NewMultiAddResponse(result []string) *organizationpb.MultiAddResponse {
	message := &organizationpb.MultiAddResponse{}
	message.Field = make([]string, len(result))
	for i, val := range result {
		message.Field[i] = val
	}
	return message
}

// NewMultiUpdatePayload builds the payload of the "multi_update" endpoint of
// the "organization" service from the gRPC request type.
func NewMultiUpdatePayload(message *organizationpb.MultiUpdateRequest) *organization.MultiUpdatePayload {
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
	return v
}

// NewMultiUpdateResponse builds the gRPC response type from the result of the
// "multi_update" endpoint of the "organization" service.
func NewMultiUpdateResponse() *organizationpb.MultiUpdateResponse {
	message := &organizationpb.MultiUpdateResponse{}
	return message
}

// ValidateAddRequest runs the validations defined on AddRequest.
func ValidateAddRequest(message *organizationpb.AddRequest) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	err = goa.MergeErrors(err, goa.ValidatePattern("message.url", message.Url, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	return
}

// ValidateMultiAddRequest runs the validations defined on MultiAddRequest.
func ValidateMultiAddRequest(message *organizationpb.MultiAddRequest) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateOrganization1(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateOrganization1 runs the validations defined on Organization1.
func ValidateOrganization1(message *organizationpb.Organization1) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	err = goa.MergeErrors(err, goa.ValidatePattern("message.url", message.Url, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	return
}

// ValidateMultiUpdateRequest runs the validations defined on
// MultiUpdateRequest.
func ValidateMultiUpdateRequest(message *organizationpb.MultiUpdateRequest) (err error) {
	if message.Ids == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ids", "message"))
	}
	if message.Organizations == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("organizations", "message"))
	}
	for _, e := range message.Organizations {
		if e != nil {
			if err2 := ValidateOrganization1(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
