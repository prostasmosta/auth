package user

import (
	"github.com/prostasmosta/auth/grpc/internal/repository"
	"github.com/prostasmosta/auth/grpc/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) service.UserService {
	return &serv{
		userRepository: userRepository,
	}
}
