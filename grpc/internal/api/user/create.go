package user

import (
	"context"
	"log"

	"github.com/prostasmosta/auth/grpc/internal/converter"
	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

func (s *Server) Create(ctx context.Context, req *grpcUser.CreateRequest) (*grpcUser.CreateResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}

	id, err := s.userService.Create(ctx, converter.ToCreateUserFromProto(req))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted user with id: %d", id)

	return &grpcUser.CreateResponse{
		Id: id,
	}, nil
}
