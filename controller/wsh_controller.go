package controller

import (
	"net/http"
	"strconv"
	"weather-app-BE/data/request"
	"weather-app-BE/helper"
	"weather-app-BE/service"

	"github.com/gin-gonic/gin"
)

type WeatherSearchHistoryController struct {
	wshService service.WeatherSearchHistoryService
}

func NewWeatherSearchHistoryController(service service.WeatherSearchHistoryService) *WeatherSearchHistoryController {
	return &WeatherSearchHistoryController{
		wshService: service,
	}
}

func (controller *WeatherSearchHistoryController) DeleteInBulk(ctx *gin.Context) {
	delWshReq := request.DeleteWeatherSearchHistoryRequest{}

	if err := ctx.ShouldBindJSON(&delWshReq); err != nil {
		helper.SendErrorResponse(err, ctx.Copy())
		return
	}

	userIdStr, _ := ctx.Get("userId")
	dbOpError := controller.wshService.DeleteInBulk(uint(userIdStr.(float64)), delWshReq.WshIds)
	if dbOpError != nil {
		helper.SendErrorResponse(dbOpError, ctx)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (controller *WeatherSearchHistoryController) Find(ctx *gin.Context) {
	wshIdStr := ctx.Query("wshId")
	userIdStr, _ := ctx.Get("userId")
	if wshIdStr != "" {
		wshId, cnvError := strconv.Atoi(wshIdStr)
		if cnvError != nil {
			helper.SendErrorResponse(cnvError, ctx)
			return
		}

		wshRes, dbOpError := controller.wshService.FindById(uint(userIdStr.(float64)), uint(wshId))
		if dbOpError != nil {
			helper.SendErrorResponse(dbOpError, ctx)
			return
		}
		ctx.JSON(http.StatusOK, wshRes)

	} else {

		wshRes, dbOpError := controller.wshService.FindAll(uint(userIdStr.(float64)))
		if dbOpError != nil {
			helper.SendErrorResponse(dbOpError, ctx)
			return
		}
		ctx.JSON(http.StatusOK, wshRes)
	}
}
