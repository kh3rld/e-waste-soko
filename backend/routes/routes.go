package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kh3rld/e-waste-soko/controller"
	"github.com/kh3rld/e-waste-soko/middleware"
)

func SetupRoutes(app *gin.RouterGroup) {
	routes := app.Group("/api")
	{
		routes.POST("/register", controller.RegisterUser)
		routes.POST("/login", controller.LoginUser)
	}

	authenticatedRoutes := routes.Group("/")
	authenticatedRoutes.Use(middleware.AuthMiddleware())
	{
		authenticatedRoutes.GET("/profile", func(c *gin.Context) {
			userID, ok := c.Get("user_id")
			if !ok {
				c.JSON(401, gin.H{"error": "Unauthorized"})
				return
			}

			c.JSON(200, gin.H{
				"message": "Welcome!",
				"user_id": userID,
			})
		})
	}
}
