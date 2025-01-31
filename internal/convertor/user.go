package convertor

import (
	"user/internal/model"
	"user/pkg/user_v1"
)

func ToUserInfoFromApi(u *user_v1.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       int(u.Age),
	}
}

func ToApiFromUser(model *model.User) *user_v1.User {
	return &user_v1.User{
		Id:        int64(model.ID),
		Info:      ToApiFromUserInfo(model.Info),
		CreatedAt: &model.CreatedAt,
		UpdatedAt: &model.UpdatedAt,
	}
}

func ToApiFromUserInfo(model model.UserInfo) *user_v1.UserInfo {
	return &user_v1.UserInfo{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
		Age:       int64(model.Age),
	}
}
