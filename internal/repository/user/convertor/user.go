package convertor

import (
	model "user/internal/model"
	modelRepo "user/internal/repository/user/model"
)

func ToUserFromRepo(u *modelRepo.User) *model.User {
	return &model.User{
		ID:   u.ID,
		Info: ToUserInfoFromRepoInfo(u.Info),
	}
}

func ToUserInfoFromRepoInfo(info modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		FirstName: info.FirstName,
		LastName:  info.LastName,
		Email:     info.Email,
		Age:       info.Age,
	}
}
