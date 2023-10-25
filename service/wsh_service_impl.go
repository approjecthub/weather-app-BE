package service

import (
	"time"
	"weather-app-BE/data/response"
	"weather-app-BE/model"
	"weather-app-BE/repository"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type WeatherSearchHistoryServiceImpl struct {
	wshRepository repository.WeatherSearchHistoryRepository
	validate      *validator.Validate
}

// DeleteInBulk implements WeatherSearchHistoryService.
func (w *WeatherSearchHistoryServiceImpl) DeleteInBulk(userId uint, wshIds []uint) error {
	err := w.wshRepository.DeleteInBulk(userId, wshIds)
	return err
}

// FindAll implements WeatherSearchHistoryService.
func (w *WeatherSearchHistoryServiceImpl) FindAll(userId uint) ([]response.WeatherSearchObject, error) {
	results, err := w.wshRepository.FindAll(userId)
	if err != nil {
		return nil, err
	}

	var wshEntries []response.WeatherSearchObject
	for _, wshData := range results {
		result := response.WeatherSearchObject{}
		result.Id = wshData.Id
		result.Place = wshData.Place
		result.Lat = wshData.Lat
		result.Lon = wshData.Lon
		result.Country = wshData.Country
		result.Description = wshData.Description
		result.Temparature = wshData.Temparature
		result.Wind = wshData.Wind
		result.Humidity = wshData.Humidity
		result.Pressure = wshData.Pressure
		result.CreatedAt = wshData.CreatedAt

		wshEntries = append(wshEntries, result)
	}

	return wshEntries, nil
}

// FindById implements WeatherSearchHistoryService.
func (w *WeatherSearchHistoryServiceImpl) FindById(userId uint, wshId uint) (result response.WeatherSearchObject, err error) {
	wshData, err := w.wshRepository.FindById(userId, wshId)
	if err != nil {
		log.Error().Msg(err.Error())
		return response.WeatherSearchObject{}, err
	}

	result.Id = wshData.Id
	result.Place = wshData.Place
	result.Lat = wshData.Lat
	result.Lon = wshData.Lon
	result.Country = wshData.Country
	result.Description = wshData.Description
	result.Temparature = wshData.Temparature
	result.Wind = wshData.Wind
	result.Humidity = wshData.Humidity
	result.Pressure = wshData.Pressure
	result.CreatedAt = wshData.CreatedAt

	return result, nil
}

// Create implements WeatherSearchHistoryService.
func (w *WeatherSearchHistoryServiceImpl) Create(wshReq response.WeatherSearchObject) error {
	err := w.validate.Struct(wshReq)

	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	wshModel := model.WeatherSearchHistory{
		Place:       wshReq.Place,
		Lat:         wshReq.Lat,
		Lon:         wshReq.Lon,
		Country:     wshReq.Country,
		Description: wshReq.Description,
		Temparature: wshReq.Temparature,
		Wind:        wshReq.Wind,
		Humidity:    wshReq.Humidity,
		Pressure:    wshReq.Pressure,
		UserID:      wshReq.UserID,
		CreatedAt:   time.Now(),
	}

	dbOpError := w.wshRepository.Save(wshModel)
	if dbOpError != nil {
		log.Error().Msg(err.Error())
		return dbOpError
	}
	return nil
}

func NewWeatherSearchHistoryServiceImpl(r repository.WeatherSearchHistoryRepository, v *validator.Validate) WeatherSearchHistoryService {
	return &WeatherSearchHistoryServiceImpl{
		wshRepository: r,
		validate:      v,
	}
}
