package router

import (
	"weather-app-BE/controller"

	"github.com/gin-gonic/gin"
)

func NewWeatherSearchHistoryRouter(baseRouter *gin.RouterGroup, wshController *controller.WeatherSearchHistoryController) {
	wshRouter := baseRouter.Group("/weather-search-history")
	wshRouter.DELETE("", wshController.DeleteInBulk)
	wshRouter.GET("", wshController.Find)
}
