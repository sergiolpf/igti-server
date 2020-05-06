// Code generated by goa v3.1.1, DO NOT EDIT.
//
// walkthrough gRPC server
//
// Command:
// $ goa gen guide.me/design

package server

import (
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
	walkthroughpb "guide.me/gen/grpc/walkthrough/pb"
	walkthrough "guide.me/gen/walkthrough"
)

// Server implements the walkthroughpb.WalkthroughServer interface.
type Server struct {
	ListH   goagrpc.UnaryHandler
	ShowH   goagrpc.UnaryHandler
	AddH    goagrpc.UnaryHandler
	RemoveH goagrpc.UnaryHandler
	UpdateH goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the walkthrough service endpoints.
func New(e *walkthrough.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		ListH:   NewListHandler(e.List, uh),
		ShowH:   NewShowHandler(e.Show, uh),
		AddH:    NewAddHandler(e.Add, uh),
		RemoveH: NewRemoveHandler(e.Remove, uh),
		UpdateH: NewUpdateHandler(e.Update, uh),
	}
}

// NewListHandler creates a gRPC handler which serves the "walkthrough" service
// "list" endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeListRequest, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in walkthroughpb.WalkthroughServer
// interface.
func (s *Server) List(ctx context.Context, message *walkthroughpb.ListRequest) (*walkthroughpb.StoredWalkthroughCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "walkthrough")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*walkthroughpb.StoredWalkthroughCollection), nil
}

// NewShowHandler creates a gRPC handler which serves the "walkthrough" service
// "show" endpoint.
func NewShowHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeShowRequest, EncodeShowResponse)
	}
	return h
}

// Show implements the "Show" method in walkthroughpb.WalkthroughServer
// interface.
func (s *Server) Show(ctx context.Context, message *walkthroughpb.ShowRequest) (*walkthroughpb.ShowResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "show")
	ctx = context.WithValue(ctx, goa.ServiceKey, "walkthrough")
	resp, err := s.ShowH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				er := err.(*walkthrough.NotFound)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewShowNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*walkthroughpb.ShowResponse), nil
}

// NewAddHandler creates a gRPC handler which serves the "walkthrough" service
// "add" endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in walkthroughpb.WalkthroughServer interface.
func (s *Server) Add(ctx context.Context, message *walkthroughpb.AddRequest) (*walkthroughpb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "walkthrough")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*walkthroughpb.AddResponse), nil
}

// NewRemoveHandler creates a gRPC handler which serves the "walkthrough"
// service "remove" endpoint.
func NewRemoveHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRemoveRequest, EncodeRemoveResponse)
	}
	return h
}

// Remove implements the "Remove" method in walkthroughpb.WalkthroughServer
// interface.
func (s *Server) Remove(ctx context.Context, message *walkthroughpb.RemoveRequest) (*walkthroughpb.RemoveResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "remove")
	ctx = context.WithValue(ctx, goa.ServiceKey, "walkthrough")
	resp, err := s.RemoveH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*walkthroughpb.RemoveResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "walkthrough"
// service "update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in walkthroughpb.WalkthroughServer
// interface.
func (s *Server) Update(ctx context.Context, message *walkthroughpb.UpdateRequest) (*walkthroughpb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "walkthrough")
	resp, err := s.UpdateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*walkthroughpb.UpdateResponse), nil
}