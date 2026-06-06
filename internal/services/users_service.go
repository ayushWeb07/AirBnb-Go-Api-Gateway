package services

import "github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/repositories"

type UserServiceInterface interface {
	CreateUser() error
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
}

func (u *UserService) CreateUser() error {
	return nil
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	newUserService := &UserService{
		UserRepository: repo,
	}

	return newUserService
}
