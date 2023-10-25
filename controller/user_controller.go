package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"weather-app-BE/data/request"
	"weather-app-BE/data/response"
	"weather-app-BE/helper"
	"weather-app-BE/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (controller *UserController) Create(ctx *gin.Context) {
	createUserReq := request.CreateUserRequest{}
	jsonValidationError := ctx.ShouldBindJSON(&createUserReq)

	if jsonValidationError != nil {
		helper.SendErrorResponse(jsonValidationError, ctx)
		return
	}

	dbOpError := controller.userService.Create(createUserReq)
	if dbOpError != nil {
		helper.SendErrorResponse(dbOpError, ctx)
		return
	}
	ctx.JSON(http.StatusCreated, response.Response{})
}

func (controller *UserController) Login(ctx *gin.Context) {
	loginUserReq := request.LoginUserRequest{}
	jsonValidationError := ctx.ShouldBindJSON(&loginUserReq)

	if jsonValidationError != nil {
		helper.SendErrorResponse(jsonValidationError, ctx)
		return
	}

	token, loginErr := controller.userService.Login(loginUserReq)
	if loginErr != nil {
		helper.SendErrorResponse(loginErr, ctx)
		return
	}

	domain := "localhost"
	ctx.SetCookie("Authorization", fmt.Sprintf("bearer %s", token), 3600, "/", domain, false, true)
}

func (controller *UserController) Update(ctx *gin.Context) {
	updateUserReq := request.UpdateUserRequest{}
	jsonValidationError := ctx.ShouldBindJSON(&updateUserReq)
	if jsonValidationError != nil {
		helper.SendErrorResponse(jsonValidationError, ctx)
		return
	}

	userId := ctx.Param("userId")
	id, cnvError := strconv.Atoi(userId)
	if cnvError != nil {
		helper.SendErrorResponse(cnvError, ctx)
		return
	}
	updateUserReq.Id = uint(id)

	dbOpError := controller.userService.Update(updateUserReq)
	if dbOpError != nil {
		helper.SendErrorResponse(dbOpError, ctx)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{})
}

func (controller *UserController) Delete(ctx *gin.Context) {

	userId := ctx.Param("userId")

	id, cnvError := strconv.Atoi(userId)
	if cnvError != nil {
		helper.SendErrorResponse(cnvError, ctx)
		return
	}
	dbOpError := controller.userService.Delete(uint(id))
	if dbOpError != nil {
		helper.SendErrorResponse(dbOpError, ctx)
		return
	}
	ctx.JSON(http.StatusOK, response.Response{})
}

func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, cnvError := strconv.Atoi(userId)
	if cnvError != nil {
		helper.SendErrorResponse(cnvError, ctx)
		return
	}
	userRes, dbOpError := controller.userService.FindById(uint(id))
	if dbOpError != nil {
		helper.SendErrorResponse(dbOpError, ctx)
		return
	}
	webRes := response.Response{
		Data: userRes,
	}
	ctx.JSON(http.StatusOK, webRes)
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	userRes, dbOpError := controller.userService.FindAll()
	if dbOpError != nil {
		helper.SendErrorResponse(dbOpError, ctx)
		return
	}
	webRes := response.Response{
		Data: userRes,
	}
	ctx.JSON(http.StatusOK, webRes)
}
