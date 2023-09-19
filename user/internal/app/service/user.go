package service

import (
	"context"

	"github.com/kenykendf/solo-project/users/internal/app/repository/user"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	repo user.Querier
}

func NewUserService(repo user.Querier) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) CreateUser(req user.CreateUserParams) error {
	if err := us.repo.CreateUser(context.Background(), &req); err != nil {
		logrus.Error("error CreateUser : ", err)
		return err
	}

	return nil
}
