package user

import (
	"context"
	"user/internal/convertor"
	"user/pkg/user_v1"
)

func (i *Implementation) List(ctx context.Context, req *user_v1.ListRequest) (*user_v1.ListResponse, error) {
	allUsers, err := i.UserService.List(ctx)
	if err != nil {
		return nil, err
	}
	var users []*user_v1.User
	for _, user := range allUsers {
		users = append(users, convertor.ToApiFromUser(user))
	}

	return &user_v1.ListResponse{User: users}, nil
}
