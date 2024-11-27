package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"my-api/controller"
	"my-api/security"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	tagRouter := router.Group("/users")
	tagRouter.POST("", userController.Create)
	tagRouter.POST("/login", controller.Login)
	

	protectedRoutes := tagRouter.Group("")
	protectedRoutes.Use(security.ProtectedHandler)


	adminRoutes := tagRouter.Group("")
	adminRoutes.Use(security.IsAdmin)
	adminRoutes.GET("", userController.FindAll)
	adminRoutes.GET("/:userId", userController.FindById)
	adminRoutes.PATCH("/:userId", userController.Update)
	adminRoutes.DELETE("/:userId", userController.Delete)

	return service
}