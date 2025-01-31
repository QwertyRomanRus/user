package repository

import (
	"context"
	"user/internal/model"
)

type UserRepo interface {
	Create(ctx context.Context, userInfo *model.UserInfo) (*model.User, error)
}
