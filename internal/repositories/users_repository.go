package repositories

import (
	"database/sql"

	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/database/models"
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
	// create a dummy user model instance
	userModel := &models.UserModel{
		Username: "ayush",
		Email:    "ayush@gmail.com",
	}

	// insert into the db
	query := "INSERT INTO users (username, email) VALUES (?, ?)"
	result, err := ur.db.Exec(query, userModel.Username, userModel.Email)

	if err != nil {
		ur.logger.Error("Failed to insert user into the database",
			zap.String("error", err.Error()))

		return
	}

	id, err := result.LastInsertId()
	
	if err != nil {
		ur.logger.Error("Failed to insert user into the database",
			zap.String("error", err.Error()))

		return
	}

	ur.logger.Info("Successfully inserted user into the database",
		zap.Int64("user_id", id))
}

func NewUserRepository(logger *zap.Logger, db *sql.DB) UserRepositoryInterface {
	newUserRepository := &UserRepository{
		db:     db,
		logger: logger,
	}

	return newUserRepository
}
