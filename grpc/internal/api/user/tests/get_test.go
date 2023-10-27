package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/prostasmosta/auth/grpc/internal/api/user"
	"github.com/prostasmosta/auth/grpc/internal/model"
	"github.com/prostasmosta/auth/grpc/internal/service"
	serviceMocks "github.com/prostasmosta/auth/grpc/internal/service/mocks"
	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *grpcUser.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		role      = 0
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		serviceErr = fmt.Errorf("service error")

		req = &grpcUser.GetRequest{
			Id: id,
		}

		serviceRes = &model.GetUser{
			Id: id,
			Info: model.UserInfo{
				Name:  name,
				Email: email,
				Role:  int32(role),
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}

		res = &grpcUser.GetResponse{
			Id: id,
			Info: &grpcUser.UserInfo{
				Name:  name,
				Email: email,
				Role:  grpcUser.Role(role),
			},
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *grpcUser.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(serviceRes, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewServer(userServiceMock)

			resHandler, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
