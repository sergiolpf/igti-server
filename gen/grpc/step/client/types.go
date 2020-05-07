// Code generated by goa v3.1.1, DO NOT EDIT.
//
// step gRPC client types
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
	steppb "guide.me/gen/grpc/step/pb"
	step "guide.me/gen/step"
	stepviews "guide.me/gen/step/views"
)

// NewListRequest builds the gRPC request type from the payload of the "list"
// endpoint of the "step" service.
func NewListRequest(payload *step.ListPayload) *steppb.ListRequest {
	message := &steppb.ListRequest{
		Id: payload.ID,
	}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the "step"
// service from the gRPC response type.
func NewListResult(message *steppb.StoredWalkthroughCollection) stepviews.StoredWalkthroughCollectionView {
	result := make([]*stepviews.StoredWalkthroughView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &stepviews.StoredWalkthroughView{
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
// endpoint of the "step" service.
func NewShowRequest(payload *step.ShowPayload) *steppb.ShowRequest {
	message := &steppb.ShowRequest{
		Id: payload.ID,
	}
	return message
}

// NewShowResult builds the result type of the "show" endpoint of the "step"
// service from the gRPC response type.
func NewShowResult(message *steppb.ShowResponse) *stepviews.StoredWalkthroughView {
	result := &stepviews.StoredWalkthroughView{
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
// "step" service from the gRPC error response type.
func NewShowNotFoundError(message *steppb.ShowNotFoundError) *step.ElementNotFound {
	er := &step.ElementNotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "step" service.
func NewAddRequest(payload *step.Walkthrough) *steppb.AddRequest {
	message := &steppb.AddRequest{
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

// NewAddResult builds the result type of the "add" endpoint of the "step"
// service from the gRPC response type.
func NewAddResult(message *steppb.AddResponse) string {
	result := message.Field
	return result
}

// NewRemoveRequest builds the gRPC request type from the payload of the
// "remove" endpoint of the "step" service.
func NewRemoveRequest(payload *step.RemovePayload) *steppb.RemoveRequest {
	message := &steppb.RemoveRequest{
		Id: payload.ID,
	}
	return message
}

// NewUpdateRequest builds the gRPC request type from the payload of the
// "update" endpoint of the "step" service.
func NewUpdateRequest(payload *step.StoredWalkthrough) *steppb.UpdateRequest {
	message := &steppb.UpdateRequest{
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

// NewPublishRequest builds the gRPC request type from the payload of the
// "publish" endpoint of the "step" service.
func NewPublishRequest(payload *step.PublishPayload) *steppb.PublishRequest {
	message := &steppb.PublishRequest{
		Id: payload.ID,
	}
	return message
}

// ValidateStoredWalkthroughCollection runs the validations defined on
// StoredWalkthroughCollection.
func ValidateStoredWalkthroughCollection(message *steppb.StoredWalkthroughCollection) (err error) {
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
func ValidateStoredWalkthrough(message *steppb.StoredWalkthrough) (err error) {
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
func ValidateShowResponse(message *steppb.ShowResponse) (err error) {
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
