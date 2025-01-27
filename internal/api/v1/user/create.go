package user

import (
	"context"
	"fmt"
	"user/pkg/user_v1"
)

func (i *Implementation) Create(_ context.Context, req *user_v1.CreateRequest) (*user_v1.CreateResponse, error) {
	fmt.Println("log from create")
	return nil, nil
}
