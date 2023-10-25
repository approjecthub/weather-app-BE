package repository

import "weather-app-BE/model"

type WeatherSearchHistoryRepository interface {
	Save(item model.WeatherSearchHistory) error
	DeleteInBulk(userId uint, wshIds []uint) error
	FindById(userId uint, wshId uint) (result model.WeatherSearchHistory, err error)
	FindAll(userId uint) ([]model.WeatherSearchHistory, error)
}
