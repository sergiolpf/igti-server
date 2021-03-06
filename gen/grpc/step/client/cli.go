// Code generated by goa v3.1.1, DO NOT EDIT.
//
// step gRPC client CLI support package
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"encoding/json"
	"fmt"

	steppb "guide.me/gen/grpc/step/pb"
	step "guide.me/gen/step"
)

// BuildListPayload builds the payload for the step list endpoint from CLI
// flags.
func BuildListPayload(stepListMessage string) (*step.ListPayload, error) {
	var err error
	var message steppb.ListRequest
	{
		if stepListMessage != "" {
			err = json.Unmarshal([]byte(stepListMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Deleniti qui.\"\n   }'")
			}
		}
	}
	v := &step.ListPayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildAddPayload builds the payload for the step add endpoint from CLI flags.
func BuildAddPayload(stepAddMessage string) (*step.AddStepPayload, error) {
	var err error
	var message steppb.AddRequest
	{
		if stepAddMessage != "" {
			err = json.Unmarshal([]byte(stepAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"step\": {\n         \"action\": \"click\",\n         \"content\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\",\n         \"placement\": \"left\",\n         \"stepNumber\": 588513995,\n         \"target\": \"Ullam minima ut et aspernatur impedit iste.\",\n         \"title\": \"Click here to make it work!\"\n      },\n      \"wtId\": \"Id vero iste voluptas.\"\n   }'")
			}
		}
	}
	v := &step.AddStepPayload{}
	if message.WtId != "" {
		v.WtID = &message.WtId
	}
	if message.Step != nil {
		v.Step = protobufSteppbStep1ToStepStep(message.Step)
	}

	return v, nil
}

// BuildRemovePayload builds the payload for the step remove endpoint from CLI
// flags.
func BuildRemovePayload(stepRemoveMessage string) (*step.RemovePayload, error) {
	var err error
	var message steppb.RemoveRequest
	{
		if stepRemoveMessage != "" {
			err = json.Unmarshal([]byte(stepRemoveMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Rerum harum.\",\n      \"wtId\": \"Dolor incidunt.\"\n   }'")
			}
		}
	}
	v := &step.RemovePayload{
		WtID: message.WtId,
		ID:   message.Id,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the step update endpoint from CLI
// flags.
func BuildUpdatePayload(stepUpdateMessage string) (*step.StoredListOfSteps, error) {
	var err error
	var message steppb.UpdateRequest
	{
		if stepUpdateMessage != "" {
			err = json.Unmarshal([]byte(stepUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"steps\": [\n         {\n            \"action\": \"next\",\n            \"content\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\",\n            \"id\": \"Sed et sit dolor aut voluptas.\",\n            \"placement\": \"top\",\n            \"stepNumber\": 68608253,\n            \"target\": \"Impedit voluptatem dolor et voluptatem.\",\n            \"title\": \"Click here to make it work!\"\n         },\n         {\n            \"action\": \"next\",\n            \"content\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\",\n            \"id\": \"Sed et sit dolor aut voluptas.\",\n            \"placement\": \"top\",\n            \"stepNumber\": 68608253,\n            \"target\": \"Impedit voluptatem dolor et voluptatem.\",\n            \"title\": \"Click here to make it work!\"\n         },\n         {\n            \"action\": \"next\",\n            \"content\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\",\n            \"id\": \"Sed et sit dolor aut voluptas.\",\n            \"placement\": \"top\",\n            \"stepNumber\": 68608253,\n            \"target\": \"Impedit voluptatem dolor et voluptatem.\",\n            \"title\": \"Click here to make it work!\"\n         },\n         {\n            \"action\": \"next\",\n            \"content\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\",\n            \"id\": \"Sed et sit dolor aut voluptas.\",\n            \"placement\": \"top\",\n            \"stepNumber\": 68608253,\n            \"target\": \"Impedit voluptatem dolor et voluptatem.\",\n            \"title\": \"Click here to make it work!\"\n         }\n      ],\n      \"wtId\": \"123abc\"\n   }'")
			}
		}
	}
	v := &step.StoredListOfSteps{
		WtID: message.WtId,
	}
	if message.Steps != nil {
		v.Steps = make([]*step.StoredStep, len(message.Steps))
		for i, val := range message.Steps {
			v.Steps[i] = &step.StoredStep{
				ID:         val.Id,
				Title:      val.Title,
				Target:     val.Target,
				StepNumber: val.StepNumber,
				Placement:  val.Placement,
				Content:    val.Content,
				Action:     val.Action,
			}
		}
	}

	return v, nil
}
