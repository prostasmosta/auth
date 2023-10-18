package user

import (
	"context"
	"github.com/prostasmosta/auth/grpc/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.GetUser, error) {
	note, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return note, nil
}
