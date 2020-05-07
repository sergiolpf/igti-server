// Code generated by goa v3.1.1, DO NOT EDIT.
//
// walkthrough endpoints
//
// Command:
// $ goa gen guide.me/design

package walkthrough

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "walkthrough" service endpoints.
type Endpoints struct {
	List    goa.Endpoint
	Show    goa.Endpoint
	Add     goa.Endpoint
	Remove  goa.Endpoint
	Update  goa.Endpoint
	Publish goa.Endpoint
}

// NewEndpoints wraps the methods of the "walkthrough" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:    NewListEndpoint(s),
		Show:    NewShowEndpoint(s),
		Add:     NewAddEndpoint(s),
		Remove:  NewRemoveEndpoint(s),
		Update:  NewUpdateEndpoint(s),
		Publish: NewPublishEndpoint(s),
	}
}

// Use applies the given middleware to all the "walkthrough" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Show = m(e.Show)
	e.Add = m(e.Add)
	e.Remove = m(e.Remove)
	e.Update = m(e.Update)
	e.Publish = m(e.Publish)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "walkthrough".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListPayload)
		res, err := s.List(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredWalkthroughCollection(res, "tiny")
		return vres, nil
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "walkthrough".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ShowPayload)
		res, view, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStoredWalkthrough(res, view)
		return vres, nil
	}
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "walkthrough".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Walkthrough)
		return s.Add(ctx, p)
	}
}

// NewRemoveEndpoint returns an endpoint function that calls the method
// "remove" of service "walkthrough".
func NewRemoveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RemovePayload)
		return nil, s.Remove(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "walkthrough".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StoredWalkthrough)
		return nil, s.Update(ctx, p)
	}
}

// NewPublishEndpoint returns an endpoint function that calls the method
// "publish" of service "walkthrough".
func NewPublishEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PublishPayload)
		return nil, s.Publish(ctx, p)
	}
}
