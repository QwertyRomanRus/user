package user

import (
	"context"
	"user/internal/convertor"
	"user/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *user_v1.GetRequest) (*user_v1.GetResponse, error) {
	user, err := i.UserService.Get(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &user_v1.GetResponse{User: convertor.ToApiFromUser(user)}, nil
}
