package router

import (
	"weather-app-BE/controller"
	"weather-app-BE/helper"

	"github.com/gin-gonic/gin"
)

func NewWeatherRouter(baseRouter *gin.RouterGroup, weatherController *controller.WeatherController) {
	weatherRouter := baseRouter.Group("/weather")
	weatherRouter.GET("/city", helper.JWTMiddleware, weatherController.GetCities)
	weatherRouter.GET("/city-weather", helper.JWTMiddleware, weatherController.GetWeather)
}
