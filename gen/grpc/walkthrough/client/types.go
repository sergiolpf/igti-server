// Code generated by goa v3.1.1, DO NOT EDIT.
//
// walkthrough gRPC client types
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	walkthroughpb "guide.me/gen/grpc/walkthrough/pb"
	walkthrough "guide.me/gen/walkthrough"
	walkthroughviews "guide.me/gen/walkthrough/views"
)

// NewListRequest builds the gRPC request type from the payload of the "list"
// endpoint of the "walkthrough" service.
func NewListRequest(payload *walkthrough.ListPayload) *walkthroughpb.ListRequest {
	message := &walkthroughpb.ListRequest{
		Id: payload.ID,
	}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the
// "walkthrough" service from the gRPC response type.
func NewListResult(message *walkthroughpb.StoredWalkthroughCollection) walkthroughviews.StoredWalkthroughCollectionView {
	result := make([]*walkthroughviews.StoredWalkthroughView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &walkthroughviews.StoredWalkthroughView{
			ID:           &val.Id,
			Name:         &val.Name,
			BaseURL:      &val.BaseUrl,
			Organization: &val.Organization,
		}
		if val.Status != "" {
			result[i].Status = &val.Status
		}
		if val.PublishedUrl != "" {
			result[i].PublishedURL = &val.PublishedUrl
		}
		if val.Status == "" {
			var tmp string = "draft"
			result[i].Status = &tmp
		}
	}
	return result
}

// NewShowRequest builds the gRPC request type from the payload of the "show"
// endpoint of the "walkthrough" service.
func NewShowRequest(payload *walkthrough.ShowPayload) *walkthroughpb.ShowRequest {
	message := &walkthroughpb.ShowRequest{
		Id: payload.ID,
	}
	return message
}

// NewShowResult builds the result type of the "show" endpoint of the
// "walkthrough" service from the gRPC response type.
func NewShowResult(message *walkthroughpb.ShowResponse) *walkthroughviews.StoredWalkthroughView {
	result := &walkthroughviews.StoredWalkthroughView{
		ID:           &message.Id,
		Name:         &message.Name,
		BaseURL:      &message.BaseUrl,
		Organization: &message.Organization,
	}
	if message.Status != "" {
		result.Status = &message.Status
	}
	if message.PublishedUrl != "" {
		result.PublishedURL = &message.PublishedUrl
	}
	if message.Status == "" {
		var tmp string = "draft"
		result.Status = &tmp
	}
	return result
}

// NewShowNotFoundError builds the error type of the "show" endpoint of the
// "walkthrough" service from the gRPC error response type.
func NewShowNotFoundError(message *walkthroughpb.ShowNotFoundError) *walkthrough.NotFound {
	er := &walkthrough.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "walkthrough" service.
func NewAddRequest(payload *walkthrough.Walkthrough) *walkthroughpb.AddRequest {
	message := &walkthroughpb.AddRequest{
		Name:         payload.Name,
		BaseUrl:      payload.BaseURL,
		Status:       payload.Status,
		Organization: payload.Organization,
	}
	if payload.PublishedURL != nil {
		message.PublishedUrl = *payload.PublishedURL
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the
// "walkthrough" service from the gRPC response type.
func NewAddResult(message *walkthroughpb.AddResponse) string {
	result := message.Field
	return result
}

// NewRemoveRequest builds the gRPC request type from the payload of the
// "remove" endpoint of the "walkthrough" service.
func NewRemoveRequest(payload *walkthrough.RemovePayload) *walkthroughpb.RemoveRequest {
	message := &walkthroughpb.RemoveRequest{
		Id: payload.ID,
	}
	return message
}

// NewUpdateRequest builds the gRPC request type from the payload of the
// "update" endpoint of the "walkthrough" service.
func NewUpdateRequest(payload *walkthrough.StoredWalkthrough) *walkthroughpb.UpdateRequest {
	message := &walkthroughpb.UpdateRequest{
		Id:           payload.ID,
		Name:         payload.Name,
		BaseUrl:      payload.BaseURL,
		Status:       payload.Status,
		Organization: payload.Organization,
	}
	if payload.PublishedURL != nil {
		message.PublishedUrl = *payload.PublishedURL
	}
	return message
}

// ValidateStoredWalkthroughCollection runs the validations defined on
// StoredWalkthroughCollection.
func ValidateStoredWalkthroughCollection(message *walkthroughpb.StoredWalkthroughCollection) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateStoredWalkthrough(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStoredWalkthrough runs the validations defined on StoredWalkthrough.
func ValidateStoredWalkthrough(message *walkthroughpb.StoredWalkthrough) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	err = goa.MergeErrors(err, goa.ValidatePattern("message.baseURL", message.BaseUrl, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	if !(message.Status == "draft" || message.Status == "completed" || message.Status == "removed") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.status", message.Status, []interface{}{"draft", "completed", "removed"}))
	}
	return
}

// ValidateShowResponse runs the validations defined on ShowResponse.
func ValidateShowResponse(message *walkthroughpb.ShowResponse) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	err = goa.MergeErrors(err, goa.ValidatePattern("message.baseURL", message.BaseUrl, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	if message.Status != "" {
		if !(message.Status == "draft" || message.Status == "completed" || message.Status == "removed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.status", message.Status, []interface{}{"draft", "completed", "removed"}))
		}
	}
	return
}