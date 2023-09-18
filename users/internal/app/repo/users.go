package repo

import (
	"database/sql"

	"github.com/kenykendf/solo-project/users/internal/app/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) CreateUser() error {

	return nil
}

func (ur *UserRepo) GetUserLists() ([]model.User, error) {

	return []model.User{}, nil
}

func (ur *UserRepo) GetUserByID() (model.User, error) {

	return model.User{}, nil
}

func (ur *UserRepo) UpdateUser() error {

	return nil
}

func (ur *UserRepo) DeleteUser() error {

	return nil
}
