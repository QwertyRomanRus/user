package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	ID        int
	Info      UserInfo
	CreatedAt timestamp.Timestamp
	UpdatedAt timestamp.Timestamp
}

type UserInfo struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}
