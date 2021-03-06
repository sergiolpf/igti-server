// Code generated by goa v3.1.1, DO NOT EDIT.
//
// step gRPC server encoders and decoders
//
// Command:
// $ goa gen guide.me/design

package server

import (
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
	steppb "guide.me/gen/grpc/step/pb"
	step "guide.me/gen/step"
	stepviews "guide.me/gen/step/views"
)

// EncodeListResponse encodes responses from the "step" service "list" endpoint.
func EncodeListResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*stepviews.StoredListOfSteps)
	if !ok {
		return nil, goagrpc.ErrInvalidType("step", "list", "*stepviews.StoredListOfSteps", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewListResponse(result)
	return resp, nil
}

// DecodeListRequest decodes requests sent to "step" service "list" endpoint.
func DecodeListRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *steppb.ListRequest
		ok      bool
	)
	{
		if message, ok = v.(*steppb.ListRequest); !ok {
			return nil, goagrpc.ErrInvalidType("step", "list", "*steppb.ListRequest", v)
		}
	}
	var payload *step.ListPayload
	{
		payload = NewListPayload(message)
	}
	return payload, nil
}

// EncodeAddResponse encodes responses from the "step" service "add" endpoint.
func EncodeAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*stepviews.ResultStep)
	if !ok {
		return nil, goagrpc.ErrInvalidType("step", "add", "*stepviews.ResultStep", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewAddResponse(result)
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "step" service "add" endpoint.
func DecodeAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *steppb.AddRequest
		ok      bool
	)
	{
		if message, ok = v.(*steppb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("step", "add", "*steppb.AddRequest", v)
		}
		if err := ValidateAddRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *step.AddStepPayload
	{
		payload = NewAddPayload(message)
	}
	return payload, nil
}

// EncodeRemoveResponse encodes responses from the "step" service "remove"
// endpoint.
func EncodeRemoveResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewRemoveResponse()
	return resp, nil
}

// DecodeRemoveRequest decodes requests sent to "step" service "remove"
// endpoint.
func DecodeRemoveRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *steppb.RemoveRequest
		ok      bool
	)
	{
		if message, ok = v.(*steppb.RemoveRequest); !ok {
			return nil, goagrpc.ErrInvalidType("step", "remove", "*steppb.RemoveRequest", v)
		}
	}
	var payload *step.RemovePayload
	{
		payload = NewRemovePayload(message)
	}
	return payload, nil
}

// EncodeUpdateResponse encodes responses from the "step" service "update"
// endpoint.
func EncodeUpdateResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewUpdateResponse()
	return resp, nil
}

// DecodeUpdateRequest decodes requests sent to "step" service "update"
// endpoint.
func DecodeUpdateRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *steppb.UpdateRequest
		ok      bool
	)
	{
		if message, ok = v.(*steppb.UpdateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("step", "update", "*steppb.UpdateRequest", v)
		}
		if err := ValidateUpdateRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *step.StoredListOfSteps
	{
		payload = NewUpdatePayload(message)
	}
	return payload, nil
}
