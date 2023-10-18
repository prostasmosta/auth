package user

import (
	"github.com/prostasmosta/auth/grpc/internal/service"
	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

type Server struct {
	grpcUser.UnimplementedUserV1Server
	userService service.UserService
}

func NewServer(userService service.UserService) *Server {
	return &Server{
		userService: userService,
	}
}
