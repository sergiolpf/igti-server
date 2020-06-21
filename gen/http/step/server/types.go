// Code generated by goa v3.1.1, DO NOT EDIT.
//
// step HTTP server types
//
// Command:
// $ goa gen guide.me/design

package server

import (
	goa "goa.design/goa/v3/pkg"
	step "guide.me/gen/step"
	stepviews "guide.me/gen/step/views"
)

// AddRequestBody is the type of the "step" service "add" endpoint HTTP request
// body.
type AddRequestBody struct {
	// Id of the walkthrough to have a step added to
	WtID *string `form:"wtId,omitempty" json:"wtId,omitempty" xml:"wtId,omitempty"`
	// step to be added
	Step *StepRequestBody `form:"step,omitempty" json:"step,omitempty" xml:"step,omitempty"`
}

// RemoveRequestBody is the type of the "step" service "remove" endpoint HTTP
// request body.
type RemoveRequestBody struct {
	// Id of the Walkthrough
	WtID *string `form:"wtId,omitempty" json:"wtId,omitempty" xml:"wtId,omitempty"`
	// ID of the step to be remove
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ListResponseBody is the type of the "step" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// ID is the unique id of the Walkthrough.
	WtID string `form:"wtId" json:"wtId" xml:"wtId"`
	// List of Stored steps
	Steps []*StoredStepResponseBody `form:"steps" json:"steps" xml:"steps"`
}

// AddResponseBody is the type of the "step" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	// Id of the walkthrough to have a step added to
	WtID string `form:"wtId" json:"wtId" xml:"wtId"`
	// Modified step
	Step *StoredStepResponseBody `form:"step" json:"step" xml:"step"`
}

// AddResponseBodyTiny is the type of the "step" service "add" endpoint HTTP
// response body.
type AddResponseBodyTiny struct {
	// Id of the walkthrough to have a step added to
	WtID string `form:"wtId" json:"wtId" xml:"wtId"`
	// Modified step
	Step *StoredStepResponseBody `form:"step" json:"step" xml:"step"`
}

// StoredStepResponseBody is used to define fields on response body types.
type StoredStepResponseBody struct {
	// Unique id to this step
	ID string `form:"id" json:"id" xml:"id"`
	// Title for the given step
	Title string `form:"title" json:"title" xml:"title"`
	// Unique html if for the target
	Target string `form:"target" json:"target" xml:"target"`
	// The number in the sequence that the step belongs to
	StepNumber int32 `form:"stepNumber" json:"stepNumber" xml:"stepNumber"`
	// Where the popup will be anchored, left, right, top or buttom.
	Placement string `form:"placement,omitempty" json:"placement,omitempty" xml:"placement,omitempty"`
	// The content of the message to be displayed
	Content string `form:"content" json:"content" xml:"content"`
	// What action should trigger the next step
	Action string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
}

// StepRequestBody is used to define fields on request body types.
type StepRequestBody struct {
	// Title for the given step
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// Unique html if for the target
	Target *string `form:"target,omitempty" json:"target,omitempty" xml:"target,omitempty"`
	// The number in the sequence that the step belongs to
	StepNumber *int32 `form:"stepNumber,omitempty" json:"stepNumber,omitempty" xml:"stepNumber,omitempty"`
	// Where the popup will be anchored, left, right, top or buttom.
	Placement *string `form:"placement,omitempty" json:"placement,omitempty" xml:"placement,omitempty"`
	// The content of the message to be displayed
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// What action should trigger the next step
	Action *string `form:"action,omitempty" json:"action,omitempty" xml:"action,omitempty"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "step" service.
func NewListResponseBody(res *stepviews.StoredListOfStepsView) *ListResponseBody {
	body := &ListResponseBody{
		WtID: *res.WtID,
	}
	if res.Steps != nil {
		body.Steps = make([]*StoredStepResponseBody, len(res.Steps))
		for i, val := range res.Steps {
			body.Steps[i] = marshalStepviewsStoredStepViewToStoredStepResponseBody(val)
		}
	}
	return body
}

// NewAddResponseBody builds the HTTP response body from the result of the
// "add" endpoint of the "step" service.
func NewAddResponseBody(res *stepviews.ResultStepView) *AddResponseBody {
	body := &AddResponseBody{
		WtID: *res.WtID,
	}
	if res.Step != nil {
		body.Step = marshalStepviewsStoredStepViewToStoredStepResponseBody(res.Step)
	}
	return body
}

// NewAddResponseBodyTiny builds the HTTP response body from the result of the
// "add" endpoint of the "step" service.
func NewAddResponseBodyTiny(res *stepviews.ResultStepView) *AddResponseBodyTiny {
	body := &AddResponseBodyTiny{
		WtID: *res.WtID,
	}
	if res.Step != nil {
		body.Step = marshalStepviewsStoredStepViewToStoredStepResponseBody(res.Step)
	}
	return body
}

// NewListPayload builds a step service list endpoint payload.
func NewListPayload(id string) *step.ListPayload {
	v := &step.ListPayload{}
	v.ID = id

	return v
}

// NewAddStepPayload builds a step service add endpoint payload.
func NewAddStepPayload(body *AddRequestBody) *step.AddStepPayload {
	v := &step.AddStepPayload{
		WtID: body.WtID,
	}
	if body.Step != nil {
		v.Step = unmarshalStepRequestBodyToStepStep(body.Step)
	}

	return v
}

// NewRemovePayload builds a step service remove endpoint payload.
func NewRemovePayload(body *RemoveRequestBody) *step.RemovePayload {
	v := &step.RemovePayload{
		WtID: *body.WtID,
		ID:   *body.ID,
	}

	return v
}

// ValidateAddRequestBody runs the validations defined on AddRequestBody
func ValidateAddRequestBody(body *AddRequestBody) (err error) {
	if body.Step != nil {
		if err2 := ValidateStepRequestBody(body.Step); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRemoveRequestBody runs the validations defined on RemoveRequestBody
func ValidateRemoveRequestBody(body *RemoveRequestBody) (err error) {
	if body.WtID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("wtId", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateStepRequestBody runs the validations defined on StepRequestBody
func ValidateStepRequestBody(body *StepRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Target == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("target", "body"))
	}
	if body.StepNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("stepNumber", "body"))
	}
	if body.Placement == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("placement", "body"))
	}
	if body.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "body"))
	}
	if body.Action == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action", "body"))
	}
	if body.Placement != nil {
		if !(*body.Placement == "left" || *body.Placement == "right" || *body.Placement == "top" || *body.Placement == "buttom") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.placement", *body.Placement, []interface{}{"left", "right", "top", "buttom"}))
		}
	}
	if body.Action != nil {
		if !(*body.Action == "click" || *body.Action == "next" || *body.Action == "end") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.action", *body.Action, []interface{}{"click", "next", "end"}))
		}
	}
	return
}
