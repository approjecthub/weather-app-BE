package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"weather-app-BE/data/response"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type WeatherServiceImpl struct {
	validate   *validator.Validate
	wshService WeatherSearchHistoryService
}

// SearchCities implements WeatherService.
func (w *WeatherServiceImpl) SearchCities(cityPrefix string) ([]response.CityResponse, error) {
	if cityPrefix == "" {
		return nil, nil
	}

	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s", cityPrefix, os.Getenv("WEATHER_API_KEY"))
	cityResponse, err := http.Get(url)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	defer cityResponse.Body.Close()
	if cityResponse.StatusCode != http.StatusOK {
		return nil, errors.New("weather api failed")
	}

	var cities []map[string]interface{}
	if err := json.NewDecoder(cityResponse.Body).Decode(&cities); err != nil {
		log.Error().Msg(err.Error())
		return nil, errors.New("weather api failed")
	}

	var results []response.CityResponse

	for _, item := range cities {
		result := response.CityResponse{
			Name:    item["name"].(string),
			Lat:     item["lat"].(float64),
			Lon:     item["lon"].(float64),
			Country: item["country"].(string),
			State:   item["state"].(string),
		}
		results = append(results, result)
	}

	return results, nil
}

func (w *WeatherServiceImpl) GetWeather(lat float64, lon float64) (response.WeatherSearchObject, error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric", lat, lon, os.Getenv("WEATHER_API_KEY"))

	WeatherSearchObject, err := http.Get(url)
	if err != nil {
		log.Error().Msg(err.Error())
		return response.WeatherSearchObject{}, err
	}

	defer WeatherSearchObject.Body.Close()
	var weather map[string]interface{}

	if err := json.NewDecoder(WeatherSearchObject.Body).Decode(&weather); err != nil {
		log.Error().Msg(err.Error())
		return response.WeatherSearchObject{}, errors.New("weather api failed")
	}

	result := response.WeatherSearchObject{}
	result.Humidity = weather["main"].(map[string]interface{})["humidity"].(float64)
	result.Pressure = weather["main"].(map[string]interface{})["pressure"].(float64)
	result.Temparature = weather["main"].(map[string]interface{})["temp"].(float64)

	result.Country = weather["sys"].(map[string]interface{})["country"].(string)

	result.Wind = weather["wind"].(map[string]interface{})["speed"].(float64)

	var descriptions []string
	for _, item := range weather["weather"].([]interface{}) {
		descriptions = append(descriptions, item.(map[string]interface{})["description"].(string))
	}

	result.Description = strings.Join(descriptions, ",")
	result.Lat = lat
	result.Lon = lon
	result.Place = weather["name"].(string)
	result.UserID = 1 //this hard coded userid will be replaced later
	result.CreatedAt = time.Now()

	dbOpError := w.wshService.Create(result)
	if dbOpError != nil {
		log.Error().Msg(dbOpError.Error())
		return response.WeatherSearchObject{}, dbOpError
	}

	return result, nil
}

func NewWeatherServiceImpl(v *validator.Validate, w WeatherSearchHistoryService) WeatherService {
	return &WeatherServiceImpl{
		validate:   v,
		wshService: w,
	}
}
