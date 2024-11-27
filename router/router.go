package router

import (
	"github.com/gin-gonic/gin"
	"my-api/controller"
	"my-api/security"
	"net/http"
)

func NewRouter(userController *controller.UserController, bookController *controller.BookController) *gin.Engine {
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

	apiGroup := service.Group("/api")

	// user group
	userRouter := apiGroup.Group("/users")
	userRouter.POST("", userController.Create)
	userRouter.POST("/login", controller.Login)

	// admin group user
	adminRoutes := userRouter.Group("")
	adminRoutes.Use(security.IsAdmin)
	adminRoutes.GET("", userController.FindAll)
	adminRoutes.GET("/:userId", userController.FindById)
	adminRoutes.PATCH("/:userId", userController.Update)
	adminRoutes.DELETE("/:userId", userController.Delete)

	// book group
	bookRouter := apiGroup.Group("/books")

	// admin group book
	bookAdminRoutes := bookRouter.Group("")
	bookAdminRoutes.Use(security.IsAdmin)
	bookAdminRoutes.POST("", bookController.Create)
	bookAdminRoutes.GET("", bookController.FindAll)
	bookAdminRoutes.GET("/:bookId", bookController.FindById)
	bookAdminRoutes.PATCH("/:bookId", bookController.Update)
	bookAdminRoutes.DELETE("/:bookId", bookController.Delete)
	bookAdminRoutes.GET("/books/search", bookController.FindByTitle)

	// user group book
	bookGuessRouter := apiGroup.Group("/books/guess")
	bookGuessRouter.Use(security.ProtectedHandler)
	bookGuessRouter.GET("/search", bookController.FindByTitleGuess)
	bookGuessRouter.GET("/search/all", bookController.FindAllGuess)

	return service
}
