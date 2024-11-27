package main

import (
	"log"
	"net/http"
	"my-api/database"
	"my-api/controller"
	"my-api/helper"
	"my-api/model"
	"my-api/repository"
	"my-api/router"
	"my-api/service"
	"github.com/go-playground/validator/v10"
)

func main() {

	// Database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	
	validate := validator.New()

	db.Table("user").AutoMigrate(&model.User{})

	// Repository
	userRepository := repository.NewUserRepositoryImpl(db)

	// Service
	userService := service.NewUserServiceImpl(userRepository, validate)

	// Controller
	userController := controller.NewUserController(userService)

	// Router
	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
