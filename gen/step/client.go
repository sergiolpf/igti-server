// Code generated by goa v3.1.1, DO NOT EDIT.
//
// step client
//
// Command:
// $ goa gen guide.me/design

package step

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "step" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	AddEndpoint    goa.Endpoint
	RemoveEndpoint goa.Endpoint
	UpdateEndpoint goa.Endpoint
}

// NewClient initializes a "step" service client given the endpoints.
func NewClient(list, add, remove, update goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		AddEndpoint:    add,
		RemoveEndpoint: remove,
		UpdateEndpoint: update,
	}
}

// List calls the "list" endpoint of the "step" service.
func (c *Client) List(ctx context.Context, p *ListPayload) (res *StoredSteps, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredSteps), nil
}

// Add calls the "add" endpoint of the "step" service.
func (c *Client) Add(ctx context.Context, p *Steps) (res string, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Remove calls the "remove" endpoint of the "step" service.
// Remove may return the following errors:
//	- "not_found" (type *ElementNotFound): Steps not found
//	- error: internal error
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}

// Update calls the "update" endpoint of the "step" service.
func (c *Client) Update(ctx context.Context, p *StoredSteps) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}
