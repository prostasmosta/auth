package repository

import (
	"context"
	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

type UserRepository interface {
	Create(ctx context.Context, params *grpcUser.CreateRequest) (int64, error)
	Get(ctx context.Context, id int64) (*grpcUser.GetResponse, error)
}
