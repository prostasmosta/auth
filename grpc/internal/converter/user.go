package converter

import (
	"github.com/prostasmosta/auth/grpc/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

func ToGetUserFromService(user *model.GetUser) *grpcUser.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &grpcUser.GetResponse{
		Id:        user.Id,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(info model.UserInfo) *grpcUser.UserInfo {
	return &grpcUser.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  grpcUser.Role(info.Role),
	}
}

func ToCreateUserFromProto(params *grpcUser.CreateRequest) *model.CreateUser {
	return &model.CreateUser{
		Info:            ToUserInfoFromProto(params.Info),
		Password:        params.Password,
		PasswordConfirm: params.PasswordConfirm,
	}
}

func ToUserInfoFromProto(info *grpcUser.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  int32(info.Role),
	}
}
