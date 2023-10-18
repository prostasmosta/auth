package service

import (
	"context"
	"github.com/prostasmosta/auth/grpc/internal/model"
)

type UserService interface {
	Create(ctx context.Context, params *model.CreateUser) (int64, error)
	Get(ctx context.Context, id int64) (*model.GetUser, error)
}
