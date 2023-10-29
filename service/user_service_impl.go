package service

import (
	"errors"
	"weather-app-BE/data/request"
	"weather-app-BE/data/response"
	"weather-app-BE/helper"
	"weather-app-BE/model"
	"weather-app-BE/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type UserServiceImpl struct {
	UsersRepository repository.UsersRepository
	validate        *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UsersRepository: userRepository,
		validate:        validate,
	}
}

// Create implements UserService.
func (t *UserServiceImpl) Create(user request.CreateUserRequest) error {
	err := t.validate.Struct(user)

	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	password, err := helper.GenerateHash(user.Password)

	if err != nil {
		return err
	}
	date, err := helper.ConverToValidDate(user.DateOfBirth)
	if err != nil {
		return err
	}

	userModel := model.User{
		Email:       user.Email,
		Password:    password,
		DateOfBirth: date,
	}

	dbOpError := t.UsersRepository.Save(userModel)
	if dbOpError != nil {
		return dbOpError
	}
	return nil
}

func (t *UserServiceImpl) Login(user request.LoginUserRequest) (string, error) {
	err := t.validate.Struct(user)

	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}

	userFetched, err := t.UsersRepository.FindByEmail(user.Email)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	if !helper.CompareHash(userFetched.Password, user.Password) {
		return "", errors.New("incorrect credentials")
	}
	token, err := helper.GenerateJWTToken(userFetched.Id)
	if err != nil {
		log.Error().Msg(err.Error())
		return "", err
	}
	return token, nil
}

// Delete implements UserService.
func (t *UserServiceImpl) Delete(userId uint) error {
	err := t.UsersRepository.Delete(userId)
	return err
}

// FindAll implements UserService.
func (t *UserServiceImpl) FindAll() ([]response.UserResponse, error) {
	result, err := t.UsersRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var users []response.UserResponse
	for _, value := range result {
		user := response.UserResponse{
			Id:          value.Id,
			Email:       value.Email,
			DateOfBirth: value.DateOfBirth,
		}

		users = append(users, user)
	}

	return users, nil
}

// FindById implements UserService.
func (t *UserServiceImpl) FindById(userId uint) (response.UserResponse, error) {
	userData, err := t.UsersRepository.FindById(userId)

	if err != nil {
		return response.UserResponse{}, err
	}
	userResponse := response.UserResponse{
		Id:          (userData.Id),
		Email:       userData.Email,
		DateOfBirth: userData.DateOfBirth,
	}

	return userResponse, nil
}

// Update implements UserService.
func (t *UserServiceImpl) Update(user request.UpdateUserRequest) error {
	err := t.validate.Struct(user)

	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	password, err := helper.GenerateHash(user.Password)

	if err != nil {
		return err
	}

	date, err := helper.ConverToValidDate(user.DateOfBirth)
	if err != nil {
		return err
	}
	userData, err := t.UsersRepository.FindById(user.Id)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	userData.DateOfBirth = date
	userData.Password = password
	userData.Email = user.Email

	t.UsersRepository.Update(userData)
	return nil
}

func (t *UserServiceImpl) Logout(ctx *gin.Context) {
	for _, cookie := range ctx.Request.Cookies() {
		ctx.SetCookie(cookie.Name, "", -1, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
	}

}
