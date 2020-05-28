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
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Optio et architecto omnis ab accusamus est.\"\n   }'")
			}
		}
	}
	v := &step.ListPayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildAddPayload builds the payload for the step add endpoint from CLI flags.
func BuildAddPayload(stepAddMessage string) (*step.Steps, error) {
	var err error
	var message steppb.AddRequest
	{
		if stepAddMessage != "" {
			err = json.Unmarshal([]byte(stepAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"steps\": [\n         {\n            \"action\": \"next\",\n            \"sequence\": 1093668274,\n            \"targetid\": \"\",\n            \"type\": \"text\",\n            \"value\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\"\n         },\n         {\n            \"action\": \"next\",\n            \"sequence\": 1093668274,\n            \"targetid\": \"\",\n            \"type\": \"text\",\n            \"value\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\"\n         },\n         {\n            \"action\": \"next\",\n            \"sequence\": 1093668274,\n            \"targetid\": \"\",\n            \"type\": \"text\",\n            \"value\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\"\n         },\n         {\n            \"action\": \"next\",\n            \"sequence\": 1093668274,\n            \"targetid\": \"\",\n            \"type\": \"text\",\n            \"value\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\"\n         }\n      ],\n      \"wtId\": \"abc234235\"\n   }'")
			}
		}
	}
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
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"Facere non et.\"\n   }'")
			}
		}
	}
	v := &step.RemovePayload{
		ID: message.Id,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the step update endpoint from CLI
// flags.
func BuildUpdatePayload(stepUpdateMessage string) (*step.StoredSteps, error) {
	var err error
	var message steppb.UpdateRequest
	{
		if stepUpdateMessage != "" {
			err = json.Unmarshal([]byte(stepUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"id\": \"123abc\",\n      \"steps\": [\n         {\n            \"action\": \"next\",\n            \"sequence\": 1093668274,\n            \"targetid\": \"\",\n            \"type\": \"text\",\n            \"value\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\"\n         },\n         {\n            \"action\": \"next\",\n            \"sequence\": 1093668274,\n            \"targetid\": \"\",\n            \"type\": \"text\",\n            \"value\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\"\n         },\n         {\n            \"action\": \"next\",\n            \"sequence\": 1093668274,\n            \"targetid\": \"\",\n            \"type\": \"text\",\n            \"value\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\"\n         },\n         {\n            \"action\": \"next\",\n            \"sequence\": 1093668274,\n            \"targetid\": \"\",\n            \"type\": \"text\",\n            \"value\": \"This dropdown contains values from the list of status, for our scenario we want to chose \\'active\\'\"\n         }\n      ],\n      \"wtId\": \"abc234235\"\n   }'")
			}
		}
	}
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

	return v, nil
}
