package models

type UserModel struct {
	ID        string `db:"id"`
	Username  string `db:"username" validate:"required"`
	Email     string `db:"email" validate:"required"`
	Password  string `db:"password" validate:"required"`
	CreatedAt string `db:"created_at" validate:"required"`
	UpdatedAt string `db:"updated_at" validate:"required"`
}
