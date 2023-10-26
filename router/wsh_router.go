package router

import (
	"weather-app-BE/controller"
	"weather-app-BE/helper"

	"github.com/gin-gonic/gin"
)

func NewWeatherSearchHistoryRouter(baseRouter *gin.RouterGroup, wshController *controller.WeatherSearchHistoryController) {
	wshRouter := baseRouter.Group("/weather-search-history")
	wshRouter.DELETE("", helper.JWTMiddleware, wshController.DeleteInBulk)
	wshRouter.GET("", helper.JWTMiddleware, wshController.Find)
}
