// Code generated by goa v3.1.1, DO NOT EDIT.
//
// step gRPC client
//
// Command:
// $ goa gen guide.me/design

package client

import (
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
	steppb "guide.me/gen/grpc/step/pb"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli steppb.StepClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the step service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: steppb.NewStepClient(cc),
		opts:    opts,
	}
}

// List calls the "List" function in steppb.StepClient interface.
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

// Add calls the "Add" function in steppb.StepClient interface.
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

// Remove calls the "Remove" function in steppb.StepClient interface.
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

// Update calls the "Update" function in steppb.StepClient interface.
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
