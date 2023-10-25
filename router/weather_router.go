package router

import (
	"weather-app-BE/controller"

	"github.com/gin-gonic/gin"
)

func NewWeatherRouter(baseRouter *gin.RouterGroup, weatherController *controller.WeatherController) {
	weatherRouter := baseRouter.Group("/weather")
	weatherRouter.GET("/city", weatherController.GetCities)
	weatherRouter.GET("/city-weather", weatherController.GetWeather)
}
