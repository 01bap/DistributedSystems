package main

import (
	"app/handlers"
	"app/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	storage.InitDB()

	// API
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:5173"},
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Use the ping handler from the handlers package
	r.GET("/ping", handlers.PingHandler)

	// Routes for /items
	r.GET("/items", handlers.GetItemsHandler)
	r.POST("/items", handlers.CreateOrUpdateItemHandler)

	// Routes for /items/:itemId
	r.GET("/items/:itemId", handlers.GetItemByIDHandler)
	r.PUT("/items/:itemId", handlers.UpdateItemHandler)
	r.DELETE("/items/:itemId", handlers.DeleteItemHandler)

	// listen and serve on localhost:8080
	r.Run()
}
