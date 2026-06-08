package services

import (
	"fmt"
	"time"

	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/config"
	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/database/models"
	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/dtos"
	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	CreateUser(userPayload *dtos.CreateUser) error
	LoginUser(userPayload *dtos.LoginUser) (string, error)
	GetAllUsers() ([]*models.UserModel, error)
	GetUserById(userPayload *dtos.GetUserById) (*models.UserModel, error)
	DeleteUserById(userPayload *dtos.DeleteUserById) error
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
	logger         *zap.Logger
	serverConfig   *config.ServerConfig
}

func (us *UserService) CreateUser(userPayload *dtos.CreateUser) error {
	us.logger.Info("Create user service called...")

	// check if the user already exists
	_, err := us.UserRepository.GetUserByUsernameAndEmail(&dtos.GetUserByUsernameAndEmail{
		Username: userPayload.Username,
		Email:    userPayload.Email,
	})

	if err == nil {
		return fmt.Errorf("User with such username and email, already exists")
	}

	// hash the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(userPayload.Password), bcrypt.DefaultCost)

	if err != nil {
		us.logger.Fatal("Something went wrong while hashing the password",
			zap.String("error", err.Error()))

		return err
	}

	userPayload.Password = string(bytes)

	// call the create user repository
	err = us.UserRepository.CreateUser(userPayload)

	if err != nil {
		return err
	}

	us.logger.Info("Create user service was successful")

	return nil
}

func (us *UserService) LoginUser(userPayload *dtos.LoginUser) (string, error) {
	us.logger.Info("Login user service called...")

	// fetch the user by username and email repository
	existingUserModel, err := us.UserRepository.GetUserByUsernameAndEmail(&dtos.GetUserByUsernameAndEmail{
		Username: userPayload.Username,
		Email:    userPayload.Email,
	})

	if err != nil {
		return "", err
	}

	// check if passwords match
	err = bcrypt.CompareHashAndPassword([]byte(existingUserModel.Password), []byte(userPayload.Password))

	if err != nil {
		us.logger.Error("Invalid password has been provided",
			zap.String("error", err.Error()))

		return "", err
	}

	// generate the jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_name":  existingUserModel.Username,
		"user_email": existingUserModel.Email,
		"user_id":    existingUserModel.ID,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(us.serverConfig.JwtSecretKey))

	if err != nil {
		us.logger.Fatal("Something went wrong while generating the token",
			zap.String("error", err.Error()))

		return "", err
	}

	us.logger.Info("Login user service was successful",
		zap.String("token", tokenString))

	return tokenString, nil
}

func (us *UserService) GetAllUsers() ([]*models.UserModel, error) {
	us.logger.Info("Get all users service called...")

	// call the fetch all users repository
	userModels, err := us.UserRepository.GetAllUsers()
	return userModels, err
}

func (us *UserService) GetUserById(userPayload *dtos.GetUserById) (*models.UserModel, error) {
	us.logger.Info("Get by id user service called...")

	// call the fetch user by id repository
	userModel, err := us.UserRepository.GetUserById(userPayload)
	return userModel, err
}

func (us *UserService) DeleteUserById(userPayload *dtos.DeleteUserById) error {
	us.logger.Info("Delete user service called...")

	// call the delete user by id repository
	err := us.UserRepository.DeleteUserById(userPayload)
	return err
}

func NewUserService(repo repositories.UserRepositoryInterface, logger *zap.Logger, serverConfig *config.ServerConfig) UserServiceInterface {
	newUserService := &UserService{
		UserRepository: repo,
		logger:         logger,
		serverConfig:   serverConfig,
	}

	return newUserService
}
