package repository

import "weather-app-BE/model"

type UsersRepository interface {
	Save(user model.User) error
	Update(user model.User) error
	Delete(userId uint) error
	FindById(userId uint) (user model.User, err error)
	FindAll() ([]model.User, error)
	FindByEmail(email string) (user model.User, err error)
}
