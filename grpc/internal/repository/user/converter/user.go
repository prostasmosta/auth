package converter

import (
	"github.com/prostasmosta/auth/grpc/internal/model"
	modelRepo "github.com/prostasmosta/auth/grpc/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.GetUser) *model.GetUser {
	return &model.GetUser{
		Id:        user.Id,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserInfoFromRepo(info modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  info.Role,
	}
}
