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

	dbOpError := controller.wshService.DeleteInBulk(1, delWshReq.WshIds)
	if dbOpError != nil {
		helper.SendErrorResponse(dbOpError, ctx)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (controller *WeatherSearchHistoryController) Find(ctx *gin.Context) {
	wshIdStr := ctx.Query("wshId")
	if wshIdStr != "" {
		wshId, cnvError := strconv.Atoi(wshIdStr)
		if cnvError != nil {
			helper.SendErrorResponse(cnvError, ctx)
			return
		}
		wshRes, dbOpError := controller.wshService.FindById(1, uint(wshId))
		if dbOpError != nil {
			helper.SendErrorResponse(dbOpError, ctx)
			return
		}
		ctx.JSON(http.StatusOK, wshRes)

	} else {
		wshRes, dbOpError := controller.wshService.FindAll(1)
		if dbOpError != nil {
			helper.SendErrorResponse(dbOpError, ctx)
			return
		}
		ctx.JSON(http.StatusOK, wshRes)
	}
}
