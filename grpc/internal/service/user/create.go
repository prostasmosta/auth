package user

import (
	"context"

	"github.com/prostasmosta/auth/grpc/internal/model"
)

func (s *serv) Create(ctx context.Context, params *model.CreateUser) (int64, error) {
	id, err := s.userRepository.Create(ctx, params)
	if err != nil {
		return 0, err
	}

	return id, nil
}
