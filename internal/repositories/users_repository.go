package repositories

import "database/sql"

type UserRepositoryInterface interface {
	CreateUser() error
}

type UserRepository struct {
	db *sql.DB
}

func (u *UserRepository) CreateUser() error {
	return nil
}
