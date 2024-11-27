package main

import (
	"log"
	"my-api/controller"
	"my-api/database"
	"my-api/helper"
	"my-api/model"
	"my-api/repository"
	"my-api/router"
	"my-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	// Database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	validate := validator.New()

	// AutoMigrate
	err = db.Table("user").AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	err = db.Table("book").AutoMigrate(&model.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Repositories
	userRepository := repository.NewUserRepositoryImpl(db)
	bookRepository := repository.NewBookRepositoryImpl(db)

	// Services
	userService := service.NewUserServiceImpl(userRepository, validate)
	bookService := service.NewBookServiceImpl(bookRepository, validate)

	// Controllers
	userController := controller.NewUserController(userService)
	bookController := controller.NewBookController(bookService)

	// Router
	mainRouter := router.NewRouter(userController, bookController)

	// Start HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mainRouter,
	}

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
