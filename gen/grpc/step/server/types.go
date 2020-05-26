// Code generated by goa v3.1.1, DO NOT EDIT.
//
// step gRPC server types
//
// Command:
// $ goa gen guide.me/design

package server

import (
	goa "goa.design/goa/v3/pkg"
	steppb "guide.me/gen/grpc/step/pb"
	step "guide.me/gen/step"
	stepviews "guide.me/gen/step/views"
)

// NewListPayload builds the payload of the "list" endpoint of the "step"
// service from the gRPC request type.
func NewListPayload(message *steppb.ListRequest) *step.ListPayload {
	v := &step.ListPayload{
		ID: message.Id,
	}
	return v
}

// NewListResponse builds the gRPC response type from the result of the "list"
// endpoint of the "step" service.
func NewListResponse(result *stepviews.StoredStepsView) *steppb.ListResponse {
	message := &steppb.ListResponse{}
	if result.ID != nil {
		message.Id = *result.ID
	}
	if result.WtID != nil {
		message.WtId = *result.WtID
	}
	if result.Steps != nil {
		message.Steps = make([]*steppb.Step1, len(result.Steps))
		for i, val := range result.Steps {
			message.Steps[i] = &steppb.Step1{}
			if val.Targetid != nil {
				message.Steps[i].Targetid = *val.Targetid
			}
			if val.Type != nil {
				message.Steps[i].Type = *val.Type
			}
			if val.Value != nil {
				message.Steps[i].Value = *val.Value
			}
			if val.Sequence != nil {
				message.Steps[i].Sequence = *val.Sequence
			}
			if val.Action != nil {
				message.Steps[i].Action = *val.Action
			}
			if val.Type == nil {
				message.Steps[i].Type = "text"
			}
		}
	}
	return message
}

// NewAddPayload builds the payload of the "add" endpoint of the "step" service
// from the gRPC request type.
func NewAddPayload(message *steppb.AddRequest) *step.Steps {
	v := &step.Steps{}
	if message.WtId != "" {
		v.WtID = &message.WtId
	}
	if message.Steps != nil {
		v.Steps = make([]*step.Step, len(message.Steps))
		for i, val := range message.Steps {
			v.Steps[i] = &step.Step{
				Targetid: val.Targetid,
				Type:     val.Type,
				Value:    val.Value,
				Sequence: val.Sequence,
				Action:   val.Action,
			}
		}
	}
	return v
}

// NewAddResponse builds the gRPC response type from the result of the "add"
// endpoint of the "step" service.
func NewAddResponse(result string) *steppb.AddResponse {
	message := &steppb.AddResponse{}
	message.Field = result
	return message
}

// NewRemovePayload builds the payload of the "remove" endpoint of the "step"
// service from the gRPC request type.
func NewRemovePayload(message *steppb.RemoveRequest) *step.RemovePayload {
	v := &step.RemovePayload{
		ID: message.Id,
	}
	return v
}

// NewRemoveResponse builds the gRPC response type from the result of the
// "remove" endpoint of the "step" service.
func NewRemoveResponse() *steppb.RemoveResponse {
	message := &steppb.RemoveResponse{}
	return message
}

// NewUpdatePayload builds the payload of the "update" endpoint of the "step"
// service from the gRPC request type.
func NewUpdatePayload(message *steppb.UpdateRequest) *step.StoredSteps {
	v := &step.StoredSteps{
		ID:   message.Id,
		WtID: message.WtId,
	}
	if message.Steps != nil {
		v.Steps = make([]*step.Step, len(message.Steps))
		for i, val := range message.Steps {
			v.Steps[i] = &step.Step{
				Targetid: val.Targetid,
				Type:     val.Type,
				Value:    val.Value,
				Sequence: val.Sequence,
				Action:   val.Action,
			}
		}
	}
	return v
}

// NewUpdateResponse builds the gRPC response type from the result of the
// "update" endpoint of the "step" service.
func NewUpdateResponse() *steppb.UpdateResponse {
	message := &steppb.UpdateResponse{}
	return message
}

// ValidateStep1 runs the validations defined on Step1.
func ValidateStep1(message *steppb.Step1) (err error) {
	if !(message.Type == "text" || message.Type == "picture") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.type", message.Type, []interface{}{"text", "picture"}))
	}
	if !(message.Action == "click" || message.Action == "next" || message.Action == "end") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.action", message.Action, []interface{}{"click", "next", "end"}))
	}
	return
}

// ValidateAddRequest runs the validations defined on AddRequest.
func ValidateAddRequest(message *steppb.AddRequest) (err error) {
	for _, e := range message.Steps {
		if e != nil {
			if err2 := ValidateStep1(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateUpdateRequest runs the validations defined on UpdateRequest.
func ValidateUpdateRequest(message *steppb.UpdateRequest) (err error) {
	if message.Steps == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("steps", "message"))
	}
	for _, e := range message.Steps {
		if e != nil {
			if err2 := ValidateStep1(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}