package service

import (
	"context"
	model "user/internal/model"
)

type User interface {
	Create(ctx context.Context, userInfo *model.UserInfo) (*model.User, error)
}
