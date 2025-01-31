package user

import (
	"user/internal/service"
	"user/pkg/user_v1"
)

type Implementation struct {
	user_v1.UnimplementedUserServiceV1Server
	UserService service.User
}
