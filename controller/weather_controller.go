package controller

import (
	"net/http"
	"strconv"
	"weather-app-BE/helper"
	"weather-app-BE/service"

	"github.com/gin-gonic/gin"
)

type WeatherController struct {
	WeatherService service.WeatherService
}

func NewWeatherController(service service.WeatherService) *WeatherController {
	return &WeatherController{
		WeatherService: service,
	}
}

func (controller *WeatherController) GetCities(ctx *gin.Context) {
	cityPrefix := ctx.Query("cityPrefix")

	cities, err := controller.WeatherService.SearchCities(cityPrefix)
	if err != nil {
		helper.SendErrorResponse(err, ctx.Copy())
		return
	}
	ctx.JSON(http.StatusOK, cities)
}

func (controller *WeatherController) GetWeather(ctx *gin.Context) {
	lat := ctx.Query("lat")
	lon := ctx.Query("lon")
	latDecimal, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		helper.SendErrorResponse(err, ctx.Copy())
		return
	}
	lonDecimal, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		helper.SendErrorResponse(err, ctx.Copy())
		return
	}

	weather, err := controller.WeatherService.GetWeather(latDecimal, lonDecimal)
	if err != nil {
		helper.SendErrorResponse(err, ctx.Copy())
		return
	}

	ctx.JSON(http.StatusOK, weather)
}
