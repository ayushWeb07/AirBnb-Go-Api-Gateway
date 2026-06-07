package repositories

import (
	"database/sql"

	"go.uber.org/zap"
)

type UserRepositoryInterface interface {
	CreateUser()
}

type UserRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func (ur *UserRepository) CreateUser() {
	ur.logger.Info("Create user repository called...")
}

func NewUserRepository(logger *zap.Logger) UserRepositoryInterface {
	newUserRepository := &UserRepository{
		db:     nil,
		logger: logger,
	}

	return newUserRepository
}
