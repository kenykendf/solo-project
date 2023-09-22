package user

import (
	"database/sql"
)

type User struct {
	ID        int64        `db:"id" json:"id"`
	Username  string       `db:"username" json:"username"`
	Email     string       `db:"email" json:"email"`
	Password  string       `db:"password" json:"password"`
	Phone     string       `db:"phone" json:"phone"`
	FirstName string       `db:"first_name" json:"first_name"`
	LastName  string       `db:"last_name" json:"last_name"`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}
