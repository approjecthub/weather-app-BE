package service

import (
	"weather-app-BE/data/response"
)

type WeatherSearchHistoryService interface {
	Create(wshReq response.WeatherSearchObject) error
	DeleteInBulk(userId uint, wshIds []uint) error
	FindById(userId uint, wshId uint) (result response.WeatherSearchObject, err error)
	FindAll(userId uint) ([]response.WeatherSearchObject, error)
}
