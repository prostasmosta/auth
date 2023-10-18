package converter

import (
	"github.com/prostasmosta/auth/grpc/internal/repository/user/model"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpcUser "github.com/prostasmosta/auth/grpc/pkg/user_v1"
)

func ToUserFromRepo(user *model.GetUser) *grpcUser.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &grpcUser.GetResponse{
		Id:        user.ID,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromRepo(info model.UserInfo) *grpcUser.UserInfo {
	return &grpcUser.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  grpcUser.Role(info.Role),
	}
}
