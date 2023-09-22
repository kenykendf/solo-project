package user

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
    username
    ,email
    ,password
    ,phone
    ,first_name
    ,last_name
    ,created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
`

type CreateUserParams struct {
	Username  string       `db:"username" json:"username"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	Phone     string       `db:"phone" json:"phone"`
	FirstName string       `db:"first_name" json:"first_name"`
	LastName  string       `db:"last_name" json:"last_name"`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Phone,
		arg.FirstName,
		arg.LastName,
		arg.CreatedAt,
	)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password, phone, first_name, last_name, created_at, updated_at, deleted_at FROM users
WHERE email = $1 LIMIT 1
`

type GetUserByEmailParams struct {
	Email string `db:"email" json:"email"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, arg *GetUserByEmailParams) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email, password, phone, first_name, last_name, created_at, updated_at, deleted_at FROM users
WHERE id = $1 LIMIT 1
`

type GetUserByIDParams struct {
	ID int64 `db:"id" json:"id"`
}

func (q *Queries) GetUserByID(ctx context.Context, arg *GetUserByIDParams) (User, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT id, username, email, password, phone, first_name, last_name, created_at, updated_at, deleted_at FROM users
WHERE phone = $1 LIMIT 1
`

type GetUserByPhoneParams struct {
	Phone string `db:"phone" json:"phone"`
}

func (q *Queries) GetUserByPhone(ctx context.Context, arg *GetUserByPhoneParams) (User, error) {
	row := q.queryRow(ctx, q.getUserByPhoneStmt, getUserByPhone, arg.Phone)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, password, phone, first_name, last_name, created_at, updated_at, deleted_at FROM users
WHERE username = $1 LIMIT 1
`

type GetUserByUsernameParams struct {
	Username string `db:"username" json:"username"`
}

func (q *Queries) GetUserByUsername(ctx context.Context, arg *GetUserByUsernameParams) (User, error) {
	row := q.queryRow(ctx, q.getUserByUsernameStmt, getUserByUsername, arg.Username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Phone,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const hardDeleteUser = `-- name: HardDeleteUser :exec
DELETE FROM users
WHERE id = $1
`

type HardDeleteUserParams struct {
	ID int64 `db:"id" json:"id"`
}

func (q *Queries) HardDeleteUser(ctx context.Context, arg *HardDeleteUserParams) error {
	_, err := q.exec(ctx, q.hardDeleteUserStmt, hardDeleteUser, arg.ID)
	return err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, password, phone, first_name, last_name, created_at, updated_at, deleted_at FROM users
ORDER BY id
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.Phone,
			&i.FirstName,
			&i.LastName,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const softDeleteUser = `-- name: SoftDeleteUser :exec
UPDATE users SET updated_at = $2, deleted_at = $3
WHERE id = $1
`

type SoftDeleteUserParams struct {
	ID        int64        `db:"id" json:"id"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

func (q *Queries) SoftDeleteUser(ctx context.Context, arg *SoftDeleteUserParams) error {
	_, err := q.exec(ctx, q.softDeleteUserStmt, softDeleteUser, arg.ID, arg.UpdatedAt, arg.DeletedAt)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET
    username = $2
    ,email = $3
    ,password = $4
    ,phone = $5
    ,first_name = $6
    ,last_name = $7
    ,updated_at = $8
WHERE id = $1
`

type UpdateUserParams struct {
	ID        int64        `db:"id" json:"id"`
	Username  string       `db:"username" json:"username"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	Phone     string       `db:"phone" json:"phone"`
	FirstName string       `db:"first_name" json:"first_name"`
	LastName  string       `db:"last_name" json:"last_name"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg *UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Phone,
		arg.FirstName,
		arg.LastName,
		arg.UpdatedAt,
	)
	return err
}
