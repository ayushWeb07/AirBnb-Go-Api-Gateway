package controllers

import (
	"net/http"

	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/services"
	"go.uber.org/zap"
)

type UserControllerInterface interface {
	CreateUser(resWriter http.ResponseWriter, req *http.Request)
}

type UserController struct {
	UserService services.UserServiceInterface
	logger      *zap.Logger
}

func (uc *UserController) CreateUser(resWriter http.ResponseWriter, req *http.Request) {
	resWriter.Write([]byte("Create user endpoint working fine!"))
	uc.UserService.CreateUser()
}

func NewUserController(service services.UserServiceInterface, logger *zap.Logger) UserControllerInterface {
	newUserController := &UserController{
		UserService: service,
		logger:      logger,
	}

	return newUserController
}
