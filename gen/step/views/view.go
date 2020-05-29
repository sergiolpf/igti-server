// Code generated by goa v3.1.1, DO NOT EDIT.
//
// step views
//
// Command:
// $ goa gen guide.me/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// StoredSteps is the viewed result type that is projected based on a view.
type StoredSteps struct {
	// Type to project
	Projected *StoredStepsView
	// View to render
	View string
}

// StoredStepsView is a type that runs validations on a projected type.
type StoredStepsView struct {
	// ID is the unique id of the Step.
	ID *string
	// The id of the Walkthrough those steps belong to.
	WtID *string
	// List of steps for a given walkthrough.
	Steps []*StepView
}

// StepView is a type that runs validations on a projected type.
type StepView struct {
	// A string representing the HTML ID of an element
	Targetid *string
	// The type of step to be used
	Type *string
	// The content of the message to be displayed
	Value *string
	// The number in the sequence that the step belongs to.
	Sequence *int32
	// What action should trigger the next step
	Action *string
}

var (
	// StoredStepsMap is a map of attribute names in result type StoredSteps
	// indexed by view name.
	StoredStepsMap = map[string][]string{
		"default": []string{
			"id",
			"wtId",
			"steps",
		},
		"tiny": []string{
			"id",
			"wtId",
			"steps",
		},
	}
)

// ValidateStoredSteps runs the validations defined on the viewed result type
// StoredSteps.
func ValidateStoredSteps(result *StoredSteps) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateStoredStepsView(result.Projected)
	case "tiny":
		err = ValidateStoredStepsViewTiny(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny"})
	}
	return
}

// ValidateStoredStepsView runs the validations defined on StoredStepsView
// using the "default" view.
func ValidateStoredStepsView(result *StoredStepsView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.WtID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("wtId", "result"))
	}
	if result.Steps == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("steps", "result"))
	}
	for _, e := range result.Steps {
		if e != nil {
			if err2 := ValidateStepView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStoredStepsViewTiny runs the validations defined on StoredStepsView
// using the "tiny" view.
func ValidateStoredStepsViewTiny(result *StoredStepsView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.WtID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("wtId", "result"))
	}
	if result.Steps == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("steps", "result"))
	}
	for _, e := range result.Steps {
		if e != nil {
			if err2 := ValidateStepView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStepView runs the validations defined on StepView.
func ValidateStepView(result *StepView) (err error) {
	if result.Targetid == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("targetid", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "result"))
	}
	if result.Sequence == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sequence", "result"))
	}
	if result.Action == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "text" || *result.Type == "picture") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"text", "picture"}))
		}
	}
	if result.Action != nil {
		if !(*result.Action == "click" || *result.Action == "next" || *result.Action == "end") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.action", *result.Action, []interface{}{"click", "next", "end"}))
		}
	}
	return
}
