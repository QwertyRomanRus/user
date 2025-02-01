package user

import (
	"context"
	model "user/internal/model"
)

func (s *service) Get(ctx context.Context, id int) (*model.User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
