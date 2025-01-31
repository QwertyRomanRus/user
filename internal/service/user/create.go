package user

import (
	"context"
	"user/internal/model"
)

func (s *service) Create(ctx context.Context, userInfo *model.UserInfo) (*model.User, error) {
	user, err := s.repo.Create(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	return user, nil
}
