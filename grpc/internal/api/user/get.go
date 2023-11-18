package user

import (
	"context"
	"log"

	"github.com/prostasmosta/auth/grpc/internal/converter"
	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

func (s *Server) Get(ctx context.Context, req *grpcUser.GetRequest) (*grpcUser.GetResponse, error) {
	err := req.ValidateAll()
	if err != nil {
		return nil, err
	}

	userObj, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("id: %d, name: %s, email: %s, role: %v, created_at: %v, updated_at: %v\n",
		userObj.Id, userObj.Info.Name, userObj.Info.Email, userObj.Info.Role, userObj.CreatedAt, userObj.UpdatedAt)

	convertedUser := converter.ToGetUserFromService(userObj)

	return &grpcUser.GetResponse{
		Id: convertedUser.Id,
		Info: &grpcUser.UserInfo{
			Name:  convertedUser.Info.Name,
			Email: convertedUser.Info.Email,
			Role:  convertedUser.Info.Role,
		},
		CreatedAt: convertedUser.CreatedAt,
		UpdatedAt: convertedUser.UpdatedAt,
	}, nil
}
