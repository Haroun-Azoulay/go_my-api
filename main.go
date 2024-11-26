package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"my-api/database"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // Par défaut, écoute sur 0.0.0.0:8080
}
