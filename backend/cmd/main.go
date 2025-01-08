package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kh3rld/e-waste-soko/database"
	"github.com/kh3rld/e-waste-soko/middleware"
	"github.com/kh3rld/e-waste-soko/routes"
)

func main() {
	app := gin.New()

	database.ConnectDatabase()

	routes.SetupRoutes(app.Group("/"))

	app.Use(middleware.AuthMiddleware())

	port := ":3000"
	log.Printf("Server running on http://localhost%s", port)
	if err := app.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
