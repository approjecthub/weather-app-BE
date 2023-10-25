package main

import (
	"net/http"

	"weather-app-BE/config"
	"weather-app-BE/controller"
	"weather-app-BE/model"
	"weather-app-BE/repository"
	"weather-app-BE/router"
	"weather-app-BE/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

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
	baseRouter := ginRouter.Group("/api")

	router.NewUserRouter(baseRouter, userController)
	router.NewWeatherRouter(baseRouter, weatherController)
	router.NewWeatherSearchHistoryRouter(baseRouter, wshController)

	server := &http.Server{
		Addr:    "localhost:3333",
		Handler: ginRouter,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
