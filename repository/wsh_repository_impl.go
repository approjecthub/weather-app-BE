package repository

import (
	"errors"
	"weather-app-BE/model"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

type WeatherSearchHistoryRepositoryImpl struct {
	Db *gorm.DB
}

// DeleteInBulk implements WeatherSearchHistoryRepository.
func (wsh *WeatherSearchHistoryRepositoryImpl) DeleteInBulk(userId uint, wshIds []uint) error {
	result := wsh.Db.Where("user_id = ? AND id IN (?)", userId, wshIds).Delete(&model.WeatherSearchHistory{})
	if err := result.Error; err != nil {
		log.Error().
			Str("file", "wsh_repository_impl.go").
			Str("method", "DeleteInBulk").Msg(err.Error())
		return err
	}
	if result.RowsAffected == 0 {
		return errors.New("history is not found")
	}
	return nil
}

// FindAll implements WeatherSearchHistoryRepository.
func (wsh *WeatherSearchHistoryRepositoryImpl) FindAll(userId uint) ([]model.WeatherSearchHistory, error) {
	var result []model.WeatherSearchHistory

	queryResult := wsh.Db.Where("user_id = ?", userId).Find(&result)
	if err := queryResult.Error; err != nil {
		log.Error().
			Str("file", "wsh_repository_impl.go").
			Str("method", "FindAll").Msg(err.Error())
		return nil, err
	}

	return result, nil
}

// FindById implements WeatherSearchHistoryRepository.
func (wsh *WeatherSearchHistoryRepositoryImpl) FindById(userId uint, wshId uint) (result model.WeatherSearchHistory, err error) {
	queryResult := wsh.Db.Where("user_id = ? AND id = ?", userId, wshId).First(&result)

	if err := queryResult.Error; err != nil {
		log.Error().
			Str("file", "wsh_repository_impl.go").
			Str("method", "FindById").Msg(err.Error())

		return model.WeatherSearchHistory{}, err
	}

	if queryResult.RowsAffected == 0 {
		return model.WeatherSearchHistory{}, errors.New("history is not found")
	}

	return result, nil
}

// Save implements WeatherSearchHistoryRepository.
func (wsh *WeatherSearchHistoryRepositoryImpl) Save(item model.WeatherSearchHistory) error {
	result := wsh.Db.Create(&item)
	if err := result.Error; err != nil {
		log.Error().
			Str("file", "wsh_repository_impl.go").
			Str("method", "Save").Msg(err.Error())

		return err
	}

	return nil
}

func NewWeatherSearchHistoryRepositoryImpl(db *gorm.DB) WeatherSearchHistoryRepository {
	return &WeatherSearchHistoryRepositoryImpl{Db: db}
}
