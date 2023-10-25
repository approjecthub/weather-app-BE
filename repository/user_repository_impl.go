package repository

import (
	"errors"
	"weather-app-BE/model"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: db}
}

// Delete implements UsersRepository.
func (t *UsersRepositoryImpl) Delete(userId uint) error {
	// Create a struct with the condition
	condition := model.User{Id: userId}

	// Use the condition to delete the user
	result := t.Db.Delete(&condition)

	if result.Error != nil {
		err := result.Error
		log.Error().
			Str("file", "user_repository_impl.go").
			Str("method", "Delete").
			Msg(err.Error())
		return err
	}

	// Check if no rows were affected, which means the user was not found
	if result.RowsAffected == 0 {
		return errors.New("user is not found")
	}

	return nil
}

// FindAll implements UsersRepository.
func (t *UsersRepositoryImpl) FindAll() ([]model.User, error) {
	var users []model.User

	result := t.Db.Select("Id", "Email", "DateOfBirth").Find(&users)
	err := result.Error
	if err != nil {
		log.Error().
			Str("file", "user_repository_impl.go").
			Str("method", "FindAll").Msg(err.Error())
		return nil, err
	}
	return users, nil
}

// FindById implements UsersRepository.
func (t *UsersRepositoryImpl) FindById(userId uint) (user model.User, err error) {
	result := t.Db.First(&user, userId)

	if result.Error != nil {
		err = result.Error
		log.Error().
			Str("file", "user_repository_impl.go").
			Str("method", "FindById").Msg(err.Error())
		return model.User{}, err
	}

	if result.RowsAffected == 0 {
		return model.User{}, errors.New("user is not found")
	}

	return user, nil
}

func (t *UsersRepositoryImpl) FindByEmail(email string) (user model.User, err error) {
	result := t.Db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		err = result.Error
		log.Error().
			Str("file", "user_repository_impl.go").
			Str("method", "FindByEmail").Msg(err.Error())
		return model.User{}, err
	}

	if result.RowsAffected == 0 {
		return model.User{}, errors.New("user is not found")
	}

	return user, nil
}

// Save implements UsersRepository.
func (t *UsersRepositoryImpl) Save(user model.User) error {
	result := t.Db.Create(&user)
	err := result.Error
	if err != nil {
		log.Error().
			Str("file", "user_repository_impl.go").
			Str("method", "Save").Msg(err.Error())

		return err
	}
	return nil
}

// Update implements UsersRepository.
func (t *UsersRepositoryImpl) Update(user model.User) error {
	var updateUser = model.User{
		Id:          user.Id,
		Email:       user.Email,
		Password:    user.Password,
		DateOfBirth: user.DateOfBirth,
	}

	result := t.Db.Model(&user).Updates(updateUser)
	err := result.Error
	if err != nil {
		log.Error().
			Str("file", "user_repository_impl.go").
			Str("method", "Update").Msg(err.Error())
		return err
	}
	return nil
}
