package user

import (
	"context"
	model "user/internal/model"
)

func (s *service) List(ctx context.Context) ([]*model.User, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
