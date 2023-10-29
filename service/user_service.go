package service

import (
	"weather-app-BE/data/request"
	"weather-app-BE/data/response"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(user request.CreateUserRequest) error
	Login(user request.LoginUserRequest) (string, error)
	Update(user request.UpdateUserRequest) error
	Delete(userId uint) error
	FindById(userId uint) (response.UserResponse, error)
	FindAll() ([]response.UserResponse, error)
	Logout(ctx *gin.Context)
}
