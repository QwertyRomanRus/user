package user

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"user/internal/repository"
)

const (
	tableName = "users"

	idColumn        = "id"
	firstNameColumn = "first_name"
	lastNameColumn  = "last_name"
	emailColumn     = "email"
	ageColumn       = "age"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepo {
	return &repo{db: db}
}
