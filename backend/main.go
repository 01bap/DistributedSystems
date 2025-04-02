package main

import (
	"log"
	"net/http"

	"app/handlers"
	"app/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	storage.InitDB()

	_, err := storage.DB.Exec("INSERT INTO items (name, quantity) VALUES ($1, $2)", "Bananas", 5)
	if err != nil {
		log.Println("Test-Daten konnten nicht eingefügt werden:", err)
	}

	//API-Routes
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/items", handlers.GetItemsHandler)

	r.POST("/items", handlers.CreateOrUpdateItemHandler)

	r.Run() // listen and serve on localhost:8080
}
