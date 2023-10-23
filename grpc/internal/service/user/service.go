package user

import (
	"github.com/prostasmosta/auth/grpc/internal/client/db"
	"github.com/prostasmosta/auth/grpc/internal/repository"
	"github.com/prostasmosta/auth/grpc/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
