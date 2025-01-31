package user

import (
	"context"
	"user/internal/convertor"
	"user/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	user, err := i.UserService.Create(ctx, convertor.ToUserInfoFromApi(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	return &user_v1.CreateResponse{User: convertor.ToApiFromUser(user)}, nil
}
