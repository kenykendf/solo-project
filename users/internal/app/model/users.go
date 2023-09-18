package model

type User struct {
	ID        int    `db:"id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Phone     string `db:"phone"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	MemberID  string `db:"member_id"`
	Address   string `db:"address"`
}
