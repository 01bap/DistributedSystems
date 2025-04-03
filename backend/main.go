package main

import (
	"app/handlers"
	"app/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	storage.InitDB()


	// API
	r := gin.Default()

	// Use the ping handler from the handlers package
	r.GET("/ping", handlers.PingHandler)

	// Other routes
	r.GET("/items", handlers.GetItemsHandler)
	r.POST("/items", handlers.CreateOrUpdateItemHandler)

	r.Run() // listen and serve on localhost:8080
}
