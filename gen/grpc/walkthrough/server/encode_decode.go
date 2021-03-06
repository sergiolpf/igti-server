// Code generated by goa v3.1.1, DO NOT EDIT.
//
// walkthrough gRPC server encoders and decoders
//
// Command:
// $ goa gen guide.me/design

package server

import (
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/metadata"
	walkthroughpb "guide.me/gen/grpc/walkthrough/pb"
	walkthrough "guide.me/gen/walkthrough"
	walkthroughviews "guide.me/gen/walkthrough/views"
)

// EncodeListResponse encodes responses from the "walkthrough" service "list"
// endpoint.
func EncodeListResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(walkthroughviews.StoredWalkthroughCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("walkthrough", "list", "walkthroughviews.StoredWalkthroughCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewStoredWalkthroughCollection(result)
	return resp, nil
}

// DecodeListRequest decodes requests sent to "walkthrough" service "list"
// endpoint.
func DecodeListRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *walkthroughpb.ListRequest
		ok      bool
	)
	{
		if message, ok = v.(*walkthroughpb.ListRequest); !ok {
			return nil, goagrpc.ErrInvalidType("walkthrough", "list", "*walkthroughpb.ListRequest", v)
		}
	}
	var payload *walkthrough.ListPayload
	{
		payload = NewListPayload(message)
	}
	return payload, nil
}

// EncodeShowResponse encodes responses from the "walkthrough" service "show"
// endpoint.
func EncodeShowResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*walkthroughviews.StoredWalkthrough)
	if !ok {
		return nil, goagrpc.ErrInvalidType("walkthrough", "show", "*walkthroughviews.StoredWalkthrough", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewShowResponse(result)
	return resp, nil
}

// DecodeShowRequest decodes requests sent to "walkthrough" service "show"
// endpoint.
func DecodeShowRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		view *string
		err  error
	)
	{
		if vals := md.Get("view"); len(vals) > 0 {
			view = &vals[0]
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
	}
	if err != nil {
		return nil, err
	}
	var (
		message *walkthroughpb.ShowRequest
		ok      bool
	)
	{
		if message, ok = v.(*walkthroughpb.ShowRequest); !ok {
			return nil, goagrpc.ErrInvalidType("walkthrough", "show", "*walkthroughpb.ShowRequest", v)
		}
	}
	var payload *walkthrough.ShowPayload
	{
		payload = NewShowPayload(message, view)
	}
	return payload, nil
}

// EncodeAddResponse encodes responses from the "walkthrough" service "add"
// endpoint.
func EncodeAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*walkthroughviews.StoredWalkthrough)
	if !ok {
		return nil, goagrpc.ErrInvalidType("walkthrough", "add", "*walkthroughviews.StoredWalkthrough", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewAddResponse(result)
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "walkthrough" service "add"
// endpoint.
func DecodeAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *walkthroughpb.AddRequest
		ok      bool
	)
	{
		if message, ok = v.(*walkthroughpb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("walkthrough", "add", "*walkthroughpb.AddRequest", v)
		}
		if err := ValidateAddRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *walkthrough.Walkthrough
	{
		payload = NewAddPayload(message)
	}
	return payload, nil
}

// EncodeRemoveResponse encodes responses from the "walkthrough" service
// "remove" endpoint.
func EncodeRemoveResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewRemoveResponse()
	return resp, nil
}

// DecodeRemoveRequest decodes requests sent to "walkthrough" service "remove"
// endpoint.
func DecodeRemoveRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *walkthroughpb.RemoveRequest
		ok      bool
	)
	{
		if message, ok = v.(*walkthroughpb.RemoveRequest); !ok {
			return nil, goagrpc.ErrInvalidType("walkthrough", "remove", "*walkthroughpb.RemoveRequest", v)
		}
	}
	var payload *walkthrough.RemovePayload
	{
		payload = NewRemovePayload(message)
	}
	return payload, nil
}

// EncodeUpdateResponse encodes responses from the "walkthrough" service
// "update" endpoint.
func EncodeUpdateResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewUpdateResponse()
	return resp, nil
}

// DecodeUpdateRequest decodes requests sent to "walkthrough" service "update"
// endpoint.
func DecodeUpdateRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *walkthroughpb.UpdateRequest
		ok      bool
	)
	{
		if message, ok = v.(*walkthroughpb.UpdateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("walkthrough", "update", "*walkthroughpb.UpdateRequest", v)
		}
		if err := ValidateUpdateRequest(message); err != nil {
			return nil, err
		}
	}
	var payload *walkthrough.StoredWalkthrough
	{
		payload = NewUpdatePayload(message)
	}
	return payload, nil
}

// EncodeRenameResponse encodes responses from the "walkthrough" service
// "rename" endpoint.
func EncodeRenameResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*walkthroughviews.StoredWalkthrough)
	if !ok {
		return nil, goagrpc.ErrInvalidType("walkthrough", "rename", "*walkthroughviews.StoredWalkthrough", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewRenameResponse(result)
	return resp, nil
}

// DecodeRenameRequest decodes requests sent to "walkthrough" service "rename"
// endpoint.
func DecodeRenameRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *walkthroughpb.RenameRequest
		ok      bool
	)
	{
		if message, ok = v.(*walkthroughpb.RenameRequest); !ok {
			return nil, goagrpc.ErrInvalidType("walkthrough", "rename", "*walkthroughpb.RenameRequest", v)
		}
	}
	var payload *walkthrough.RenamePayload
	{
		payload = NewRenamePayload(message)
	}
	return payload, nil
}

// EncodePublishResponse encodes responses from the "walkthrough" service
// "publish" endpoint.
func EncodePublishResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	resp := NewPublishResponse()
	return resp, nil
}

// DecodePublishRequest decodes requests sent to "walkthrough" service
// "publish" endpoint.
func DecodePublishRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *walkthroughpb.PublishRequest
		ok      bool
	)
	{
		if message, ok = v.(*walkthroughpb.PublishRequest); !ok {
			return nil, goagrpc.ErrInvalidType("walkthrough", "publish", "*walkthroughpb.PublishRequest", v)
		}
	}
	var payload *walkthrough.PublishPayload
	{
		payload = NewPublishPayload(message)
	}
	return payload, nil
}
