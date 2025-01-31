package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	ID        int `db:"id"`
	Info      UserInfo
	CreatedAt timestamp.Timestamp `db:"created_at"`
	UpdatedAt timestamp.Timestamp `db:"updated_at"`
}

type UserInfo struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Age       int    `db:"age"`
}
