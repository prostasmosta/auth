package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/prostasmosta/auth/grpc/internal/api/user"
	"github.com/prostasmosta/auth/grpc/internal/model"
	"github.com/prostasmosta/auth/grpc/internal/service"
	serviceMocks "github.com/prostasmosta/auth/grpc/internal/service/mocks"
	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *grpcUser.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id              = gofakeit.Int64()
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		role            = 0
		password        = gofakeit.Password(true, true, true, true, false, 5)
		passwordConfirm = password

		serviceErr = fmt.Errorf("service error")

		req = &grpcUser.CreateRequest{
			Info: &grpcUser.UserInfo{
				Name:  name,
				Email: email,
				Role:  grpcUser.Role(role),
			},
			Password:        password,
			PasswordConfirm: passwordConfirm,
		}

		createUser = &model.CreateUser{
			Info: model.UserInfo{
				Name:  name,
				Email: email,
				Role:  int32(role),
			},
			Password:        password,
			PasswordConfirm: passwordConfirm,
		}

		res = &grpcUser.CreateResponse{
			Id: id,
		}
	)

	t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *grpcUser.CreateResponse
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
				mock.CreateMock.Expect(ctx, createUser).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, createUser).Return(0, serviceErr)
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

			resHandler, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resHandler)
		})
	}
}
