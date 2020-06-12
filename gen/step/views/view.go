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

// ResultStep is the viewed result type that is projected based on a view.
type ResultStep struct {
	// Type to project
	Projected *ResultStepView
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
	// Title for the given step
	Title *string
	// Unique html if for the target
	Target *string
	// The number in the sequence that the step belongs to
	StepNumber *int32
	// Where the popup will be anchored, left, right, top or buttom.
	Placement *string
	// The content of the message to be displayed
	Content *string
	// What action should trigger the next step
	Action *string
}

// ResultStepView is a type that runs validations on a projected type.
type ResultStepView struct {
	// Id of the walkthrough to have a step added to
	WtID *string
	// Modified step
	Step *StoredStepView
}

// StoredStepView is a type that runs validations on a projected type.
type StoredStepView struct {
	// Unique id to this step
	ID *string
	// Title for the given step
	Title *string
	// Unique html if for the target
	Target *string
	// The number in the sequence that the step belongs to
	StepNumber *int32
	// Where the popup will be anchored, left, right, top or buttom.
	Placement *string
	// The content of the message to be displayed
	Content *string
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
	// ResultStepMap is a map of attribute names in result type ResultStep indexed
	// by view name.
	ResultStepMap = map[string][]string{
		"default": []string{
			"wtId",
			"step",
		},
		"tiny": []string{
			"wtId",
			"step",
		},
	}
	// StoredStepMap is a map of attribute names in result type StoredStep indexed
	// by view name.
	StoredStepMap = map[string][]string{
		"default": []string{
			"id",
			"title",
			"target",
			"stepNumber",
			"placement",
			"content",
			"action",
		},
		"tiny": []string{
			"id",
			"title",
			"target",
			"stepNumber",
			"content",
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

// ValidateResultStep runs the validations defined on the viewed result type
// ResultStep.
func ValidateResultStep(result *ResultStep) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResultStepView(result.Projected)
	case "tiny":
		err = ValidateResultStepViewTiny(result.Projected)
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
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.Target == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("target", "result"))
	}
	if result.StepNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("stepNumber", "result"))
	}
	if result.Placement == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("placement", "result"))
	}
	if result.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "result"))
	}
	if result.Action == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action", "result"))
	}
	if result.Placement != nil {
		if !(*result.Placement == "left" || *result.Placement == "right" || *result.Placement == "top" || *result.Placement == "buttom") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.placement", *result.Placement, []interface{}{"left", "right", "top", "buttom"}))
		}
	}
	if result.Action != nil {
		if !(*result.Action == "click" || *result.Action == "next" || *result.Action == "end") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.action", *result.Action, []interface{}{"click", "next", "end"}))
		}
	}
	return
}

// ValidateResultStepView runs the validations defined on ResultStepView using
// the "default" view.
func ValidateResultStepView(result *ResultStepView) (err error) {
	if result.WtID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("wtId", "result"))
	}
	if result.Step != nil {
		if err2 := ValidateStoredStepView(result.Step); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResultStepViewTiny runs the validations defined on ResultStepView
// using the "tiny" view.
func ValidateResultStepViewTiny(result *ResultStepView) (err error) {
	if result.WtID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("wtId", "result"))
	}
	if result.Step != nil {
		if err2 := ValidateStoredStepView(result.Step); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredStepView runs the validations defined on StoredStepView using
// the "default" view.
func ValidateStoredStepView(result *StoredStepView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.Target == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("target", "result"))
	}
	if result.StepNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("stepNumber", "result"))
	}
	if result.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "result"))
	}
	if result.Placement == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("placement", "result"))
	}
	if result.Action == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action", "result"))
	}
	if result.Placement != nil {
		if !(*result.Placement == "left" || *result.Placement == "right" || *result.Placement == "top" || *result.Placement == "buttom") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.placement", *result.Placement, []interface{}{"left", "right", "top", "buttom"}))
		}
	}
	if result.Action != nil {
		if !(*result.Action == "click" || *result.Action == "next" || *result.Action == "end") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.action", *result.Action, []interface{}{"click", "next", "end"}))
		}
	}
	return
}

// ValidateStoredStepViewTiny runs the validations defined on StoredStepView
// using the "tiny" view.
func ValidateStoredStepViewTiny(result *StoredStepView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.Target == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("target", "result"))
	}
	if result.StepNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("stepNumber", "result"))
	}
	if result.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "result"))
	}
	return
}
