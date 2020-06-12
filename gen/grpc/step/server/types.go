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
			if val.Title != nil {
				message.Steps[i].Title = *val.Title
			}
			if val.Target != nil {
				message.Steps[i].Target = *val.Target
			}
			if val.StepNumber != nil {
				message.Steps[i].StepNumber = *val.StepNumber
			}
			if val.Placement != nil {
				message.Steps[i].Placement = *val.Placement
			}
			if val.Content != nil {
				message.Steps[i].Content = *val.Content
			}
			if val.Action != nil {
				message.Steps[i].Action = *val.Action
			}
			if val.Placement == nil {
				message.Steps[i].Placement = "right"
			}
			if val.Action == nil {
				message.Steps[i].Action = "next"
			}
		}
	}
	return message
}

// NewAddPayload builds the payload of the "add" endpoint of the "step" service
// from the gRPC request type.
func NewAddPayload(message *steppb.AddRequest) *step.AddStepPayload {
	v := &step.AddStepPayload{}
	if message.WtId != "" {
		v.WtID = &message.WtId
	}
	if message.Step != nil {
		v.Step = protobufSteppbStep1ToStepStep(message.Step)
	}
	return v
}

// NewAddResponse builds the gRPC response type from the result of the "add"
// endpoint of the "step" service.
func NewAddResponse(result *stepviews.ResultStepView) *steppb.AddResponse {
	message := &steppb.AddResponse{}
	if result.WtID != nil {
		message.WtId = *result.WtID
	}
	if result.Step != nil {
		message.Step = svcStepviewsStoredStepViewToSteppbStoredStep(result.Step)
	}
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
				Title:      val.Title,
				Target:     val.Target,
				StepNumber: val.StepNumber,
				Placement:  val.Placement,
				Content:    val.Content,
				Action:     val.Action,
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
	if !(message.Placement == "left" || message.Placement == "right" || message.Placement == "top" || message.Placement == "buttom") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.placement", message.Placement, []interface{}{"left", "right", "top", "buttom"}))
	}
	if !(message.Action == "click" || message.Action == "next" || message.Action == "end") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError("message.action", message.Action, []interface{}{"click", "next", "end"}))
	}
	return
}

// ValidateAddRequest runs the validations defined on AddRequest.
func ValidateAddRequest(message *steppb.AddRequest) (err error) {
	if message.Step != nil {
		if err2 := ValidateStep1(message.Step); err2 != nil {
			err = goa.MergeErrors(err, err2)
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

// protobufSteppbStep1ToStepStep builds a value of type *step.Step from a value
// of type *steppb.Step1.
func protobufSteppbStep1ToStepStep(v *steppb.Step1) *step.Step {
	if v == nil {
		return nil
	}
	res := &step.Step{
		Title:      v.Title,
		Target:     v.Target,
		StepNumber: v.StepNumber,
		Placement:  v.Placement,
		Content:    v.Content,
		Action:     v.Action,
	}

	return res
}

// svcStepStepToSteppbStep1 builds a value of type *steppb.Step1 from a value
// of type *step.Step.
func svcStepStepToSteppbStep1(v *step.Step) *steppb.Step1 {
	if v == nil {
		return nil
	}
	res := &steppb.Step1{
		Title:      v.Title,
		Target:     v.Target,
		StepNumber: v.StepNumber,
		Placement:  v.Placement,
		Content:    v.Content,
		Action:     v.Action,
	}

	return res
}

// svcStepviewsStoredStepViewToSteppbStoredStep builds a value of type
// *steppb.StoredStep from a value of type *stepviews.StoredStepView.
func svcStepviewsStoredStepViewToSteppbStoredStep(v *stepviews.StoredStepView) *steppb.StoredStep {
	res := &steppb.StoredStep{}
	if v.ID != nil {
		res.Id = *v.ID
	}
	if v.Title != nil {
		res.Title = *v.Title
	}
	if v.Target != nil {
		res.Target = *v.Target
	}
	if v.StepNumber != nil {
		res.StepNumber = *v.StepNumber
	}
	if v.Placement != nil {
		res.Placement = *v.Placement
	}
	if v.Content != nil {
		res.Content = *v.Content
	}
	if v.Action != nil {
		res.Action = *v.Action
	}
	if v.Placement == nil {
		res.Placement = "right"
	}
	if v.Action == nil {
		res.Action = "next"
	}

	return res
}

// protobufSteppbStoredStepToStepviewsStoredStepView builds a value of type
// *stepviews.StoredStepView from a value of type *steppb.StoredStep.
func protobufSteppbStoredStepToStepviewsStoredStepView(v *steppb.StoredStep) *stepviews.StoredStepView {
	res := &stepviews.StoredStepView{
		ID:         &v.Id,
		Title:      &v.Title,
		Target:     &v.Target,
		StepNumber: &v.StepNumber,
		Placement:  &v.Placement,
		Content:    &v.Content,
		Action:     &v.Action,
	}

	return res
}
