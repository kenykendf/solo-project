package service

import (
	"context"
	"database/sql"
	"time"

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

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	data := user.User{}
	params := user.GetUserByEmailParams{Email: email}

	user, err := us.repo.GetUserByEmail(ctx, &params)
	if err != nil {
		logrus.Error("error GetUserByEmail : ", err)
		return data, err
	}

	data = user

	return data, err
}

func (us *UserService) GetUserByID(ctx context.Context, ID int64) (user.User, error) {
	data := user.User{}
	params := user.GetUserByIDParams{ID: ID}

	user, err := us.repo.GetUserByID(ctx, &params)
	if err != nil {
		logrus.Error("error GetUserByID : ", err)
		return data, err
	}

	data = user

	return data, err
}

func (us *UserService) GetUserByPhone(ctx context.Context, phone string) (user.User, error) {
	data := user.User{}
	params := user.GetUserByPhoneParams{Phone: phone}

	user, err := us.repo.GetUserByPhone(ctx, &params)
	if err != nil {
		logrus.Error("error GetUserByPhone : ", err)
		return data, err
	}

	data = user

	return data, err
}

func (us *UserService) GetUserByUsername(ctx context.Context, username string) (user.User, error) {
	data := user.User{}
	params := user.GetUserByUsernameParams{Username: username}

	user, err := us.repo.GetUserByUsername(ctx, &params)
	if err != nil {
		logrus.Error("error GetUserByUsername : ", err)
		return data, err
	}

	data = user

	return data, err
}

func (us *UserService) HardDeleteUser(ctx context.Context, ID int64) error {
	param := user.HardDeleteUserParams{
		ID: ID,
	}

	if err := us.repo.HardDeleteUser(ctx, &param); err != nil {
		logrus.Error("error SoftDeleteUser : ", err)
		return err
	}

	return nil
}

func (us *UserService) ListUsers(ctx context.Context) ([]user.User, error) {
	data := []user.User{}

	users, err := us.repo.ListUsers(ctx)
	if err != nil {
		logrus.Error("error SoftDeleteUser : ", err)
		return data, err
	}

	data = append(data, users...)

	return data, nil
}

func (us *UserService) SoftDeleteUser(ctx context.Context, ID int64) error {
	param := user.SoftDeleteUserParams{
		ID:        ID,
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		DeletedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	if err := us.repo.SoftDeleteUser(ctx, &param); err != nil {
		logrus.Error("error SoftDeleteUser : ", err)
		return err
	}

	return nil
}

func (us *UserService) UpdateUser(ctx context.Context, req *user.UpdateUserParams) error {
	param := user.UpdateUserParams{}
	data, err := us.GetUserByID(ctx, req.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			logrus.Error("error GetUserByID : ", err)
			return err
		}
		logrus.Error("error occurs : ", err)
		return err
	}

	if req.ID == 0 {
		param.ID = data.ID
	} else {
		param.ID = req.ID
	}

	if req.Username == "" {
		param.Username = data.Username
	} else {
		param.Username = req.Username
	}

	if req.Email == "" {
		param.Email = data.Email
	} else {
		param.Email = req.Email
	}

	if req.Password == "" {
		param.Password = data.Password
	} else {
		param.Password = req.Password
	}

	if req.Phone == "" {
		param.Phone = data.Phone
	} else {
		param.Phone = req.Phone
	}

	if req.FirstName == "" {
		param.FirstName = data.FirstName
	} else {
		param.FirstName = req.FirstName
	}

	if req.LastName == "" {
		param.LastName = data.LastName
	} else {
		param.LastName = req.LastName
	}

	param = user.UpdateUserParams{
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	if err := us.repo.UpdateUser(ctx, &param); err != nil {
		logrus.Error("error UpdateUser : ", err)
		return err
	}

	return nil
}
