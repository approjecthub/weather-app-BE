package router

import (
	"weather-app-BE/controller"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(baseRouter *gin.RouterGroup, userController *controller.UserController) {
	userRouter := baseRouter.Group("/user")
	userRouter.POST("/register", userController.Create)
	userRouter.POST("/login", userController.Login)

	userRouter.GET("", userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.PUT("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)

}
