package repository

import (
	"context"
	model "user/internal/model"
)

type UserRepo interface {
	Create(ctx context.Context, userInfo *model.UserInfo) (*model.User, error)
	Get(ctx context.Context, id int) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
}
