package user

import "user/internal/repository"

type service struct {
	repo repository.UserRepo
}

func NewService(repo repository.UserRepo) repository.UserRepo {
	return &service{repo: repo}
}
