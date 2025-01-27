package user

import "user/pkg/user_v1"

type Implementation struct {
	user_v1.UnimplementedUserServiceV1Server
}
