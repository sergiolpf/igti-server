// Code generated by goa v3.1.1, DO NOT EDIT.
//
// organization gRPC server
//
// Command:
// $ goa gen guide.me/design

package server

import (
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
	organizationpb "guide.me/gen/grpc/organization/pb"
	organization "guide.me/gen/organization"
)

// Server implements the organizationpb.OrganizationServer interface.
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

// New instantiates the server struct with the organization service endpoints.
func New(e *organization.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		ListH:   NewListHandler(e.List, uh),
		ShowH:   NewShowHandler(e.Show, uh),
		AddH:    NewAddHandler(e.Add, uh),
		RemoveH: NewRemoveHandler(e.Remove, uh),
		UpdateH: NewUpdateHandler(e.Update, uh),
	}
}

// NewListHandler creates a gRPC handler which serves the "organization"
// service "list" endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in organizationpb.OrganizationServer
// interface.
func (s *Server) List(ctx context.Context, message *organizationpb.ListRequest) (*organizationpb.StoredOrganizationCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "organization")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*organizationpb.StoredOrganizationCollection), nil
}

// NewShowHandler creates a gRPC handler which serves the "organization"
// service "show" endpoint.
func NewShowHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeShowRequest, EncodeShowResponse)
	}
	return h
}

// Show implements the "Show" method in organizationpb.OrganizationServer
// interface.
func (s *Server) Show(ctx context.Context, message *organizationpb.ShowRequest) (*organizationpb.ShowResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "show")
	ctx = context.WithValue(ctx, goa.ServiceKey, "organization")
	resp, err := s.ShowH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				er := err.(*organization.ElementNotFound)
				return nil, goagrpc.NewStatusError(codes.NotFound, err, NewShowNotFoundError(er))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*organizationpb.ShowResponse), nil
}

// NewAddHandler creates a gRPC handler which serves the "organization" service
// "add" endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in organizationpb.OrganizationServer
// interface.
func (s *Server) Add(ctx context.Context, message *organizationpb.AddRequest) (*organizationpb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "organization")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*organizationpb.AddResponse), nil
}

// NewRemoveHandler creates a gRPC handler which serves the "organization"
// service "remove" endpoint.
func NewRemoveHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeRemoveRequest, EncodeRemoveResponse)
	}
	return h
}

// Remove implements the "Remove" method in organizationpb.OrganizationServer
// interface.
func (s *Server) Remove(ctx context.Context, message *organizationpb.RemoveRequest) (*organizationpb.RemoveResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "remove")
	ctx = context.WithValue(ctx, goa.ServiceKey, "organization")
	resp, err := s.RemoveH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*organizationpb.RemoveResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "organization"
// service "update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in organizationpb.OrganizationServer
// interface.
func (s *Server) Update(ctx context.Context, message *organizationpb.UpdateRequest) (*organizationpb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "organization")
	resp, err := s.UpdateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*organizationpb.UpdateResponse), nil
}
