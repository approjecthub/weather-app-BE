package main

import (
	"fmt"
	"os"

	"weather-app-BE/config"
	"weather-app-BE/controller"
	"weather-app-BE/model"
	"weather-app-BE/repository"
	"weather-app-BE/router"
	"weather-app-BE/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTION"}
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}

	return cors.New(config)
}
func main() {
	log.Info().Msg("Server started...")
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	db := config.DatabaseConnection()

	validate := validator.New()
	db.Table("users").AutoMigrate(&model.User{})
	db.Table("weather_search_histories").AutoMigrate(&model.WeatherSearchHistory{})

	//repository
	userRepository := repository.NewUserRepositoryImpl(db)
	wshRepository := repository.NewWeatherSearchHistoryRepositoryImpl(db)
	//service
	userService := service.NewUserServiceImpl(userRepository, validate)
	wshService := service.NewWeatherSearchHistoryServiceImpl(wshRepository, validate)
	weatherService := service.NewWeatherServiceImpl(validate, wshService)
	//controller
	userController := controller.NewUserController(userService)
	weatherController := controller.NewWeatherController(weatherService)
	wshController := controller.NewWeatherSearchHistoryController(wshService)

	ginRouter := gin.Default()
	ginRouter.Use(CORSMiddleware())
	baseRouter := ginRouter.Group("/api")
	router.NewUserRouter(baseRouter, userController)
	router.NewWeatherRouter(baseRouter, weatherController)
	router.NewWeatherSearchHistoryRouter(baseRouter, wshController)

	ginRouter.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
