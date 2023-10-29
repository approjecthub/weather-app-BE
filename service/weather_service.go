package service

import "weather-app-BE/data/response"

type WeatherService interface {
	SearchCities(cityPrefix string) ([]response.CityResponse, error)
	GetWeather(lat float64, lon float64, userId uint) (response.WeatherSearchObject, error)
}
