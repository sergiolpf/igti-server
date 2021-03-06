// Code generated by goa v3.1.1, DO NOT EDIT.
//
// walkthrough gRPC client
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
	walkthroughpb "guide.me/gen/grpc/walkthrough/pb"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli walkthroughpb.WalkthroughClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the walkthrough service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: walkthroughpb.NewWalkthroughClient(cc),
		opts:    opts,
	}
}

// List calls the "List" function in walkthroughpb.WalkthroughClient interface.
func (c *Client) List() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildListFunc(c.grpccli, c.opts...),
			EncodeListRequest,
			DecodeListResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Show calls the "Show" function in walkthroughpb.WalkthroughClient interface.
func (c *Client) Show() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildShowFunc(c.grpccli, c.opts...),
			EncodeShowRequest,
			DecodeShowResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *walkthroughpb.ShowNotFoundError:
				return nil, NewShowNotFoundError(message)
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Add calls the "Add" function in walkthroughpb.WalkthroughClient interface.
func (c *Client) Add() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildAddFunc(c.grpccli, c.opts...),
			EncodeAddRequest,
			DecodeAddResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Remove calls the "Remove" function in walkthroughpb.WalkthroughClient
// interface.
func (c *Client) Remove() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildRemoveFunc(c.grpccli, c.opts...),
			EncodeRemoveRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Update calls the "Update" function in walkthroughpb.WalkthroughClient
// interface.
func (c *Client) Update() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildUpdateFunc(c.grpccli, c.opts...),
			EncodeUpdateRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Rename calls the "Rename" function in walkthroughpb.WalkthroughClient
// interface.
func (c *Client) Rename() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildRenameFunc(c.grpccli, c.opts...),
			EncodeRenameRequest,
			DecodeRenameResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Publish calls the "Publish" function in walkthroughpb.WalkthroughClient
// interface.
func (c *Client) Publish() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildPublishFunc(c.grpccli, c.opts...),
			EncodePublishRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
