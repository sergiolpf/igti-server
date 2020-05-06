// Code generated by goa v3.1.1, DO NOT EDIT.
//
// organization gRPC client types
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	organizationpb "guide.me/gen/grpc/organization/pb"
	organization "guide.me/gen/organization"
	organizationviews "guide.me/gen/organization/views"
)

// NewListRequest builds the gRPC request type from the payload of the "list"
// endpoint of the "organization" service.
func NewListRequest() *organizationpb.ListRequest {
	message := &organizationpb.ListRequest{}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the
// "organization" service from the gRPC response type.
func NewListResult(message *organizationpb.StoredOrganizationCollection) organizationviews.StoredOrganizationCollectionView {
	result := make([]*organizationviews.StoredOrganizationView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &organizationviews.StoredOrganizationView{
			ID:   &val.Id,
			Name: &val.Name,
			URL:  &val.Url,
		}
	}
	return result
}

// NewShowRequest builds the gRPC request type from the payload of the "show"
// endpoint of the "organization" service.
func NewShowRequest(payload *organization.ShowPayload) *organizationpb.ShowRequest {
	message := &organizationpb.ShowRequest{
		Id: payload.ID,
	}
	return message
}

// NewShowResult builds the result type of the "show" endpoint of the
// "organization" service from the gRPC response type.
func NewShowResult(message *organizationpb.ShowResponse) *organizationviews.StoredOrganizationView {
	result := &organizationviews.StoredOrganizationView{
		ID:   &message.Id,
		Name: &message.Name,
		URL:  &message.Url,
	}
	return result
}

// NewShowNotFoundError builds the error type of the "show" endpoint of the
// "organization" service from the gRPC error response type.
func NewShowNotFoundError(message *organizationpb.ShowNotFoundError) *organization.NotFound {
	er := &organization.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "organization" service.
func NewAddRequest(payload *organization.Organization) *organizationpb.AddRequest {
	message := &organizationpb.AddRequest{
		Name: payload.Name,
		Url:  payload.URL,
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the
// "organization" service from the gRPC response type.
func NewAddResult(message *organizationpb.AddResponse) string {
	result := message.Field
	return result
}

// NewRemoveRequest builds the gRPC request type from the payload of the
// "remove" endpoint of the "organization" service.
func NewRemoveRequest(payload *organization.RemovePayload) *organizationpb.RemoveRequest {
	message := &organizationpb.RemoveRequest{
		Id: payload.ID,
	}
	return message
}

// NewUpdateRequest builds the gRPC request type from the payload of the
// "update" endpoint of the "organization" service.
func NewUpdateRequest(payload *organization.StoredOrganization) *organizationpb.UpdateRequest {
	message := &organizationpb.UpdateRequest{
		Id:   payload.ID,
		Name: payload.Name,
		Url:  payload.URL,
	}
	return message
}

// ValidateStoredOrganizationCollection runs the validations defined on
// StoredOrganizationCollection.
func ValidateStoredOrganizationCollection(message *organizationpb.StoredOrganizationCollection) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateStoredOrganization(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStoredOrganization runs the validations defined on
// StoredOrganization.
func ValidateStoredOrganization(message *organizationpb.StoredOrganization) (err error) {
	if utf8.RuneCountInString(message.Name) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 200, false))
	}
	err = goa.MergeErrors(err, goa.ValidatePattern("message.url", message.Url, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	return
}

// ValidateShowResponse runs the validations defined on ShowResponse.
func ValidateShowResponse(message *organizationpb.ShowResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 200 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 200, false))
	}
	err = goa.MergeErrors(err, goa.ValidatePattern("message.url", message.Url, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	return
}
