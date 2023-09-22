package user

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg *CreateUserParams) error
	GetUserByEmail(ctx context.Context, arg *GetUserByEmailParams) (User, error)
	GetUserByID(ctx context.Context, arg *GetUserByIDParams) (User, error)
	GetUserByPhone(ctx context.Context, arg *GetUserByPhoneParams) (User, error)
	GetUserByUsername(ctx context.Context, arg *GetUserByUsernameParams) (User, error)
	HardDeleteUser(ctx context.Context, arg *HardDeleteUserParams) error
	ListUsers(ctx context.Context) ([]User, error)
	SoftDeleteUser(ctx context.Context, arg *SoftDeleteUserParams) error
	UpdateUser(ctx context.Context, arg *UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
