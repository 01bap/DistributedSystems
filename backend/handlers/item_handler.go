package handlers

import (
	"net/http"

	"app/storage"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Get all shopping items
func GetItemsHandler(c *gin.Context) {
	items, err := storage.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}
	c.JSON(http.StatusOK, items)
}

// Create or Update a shopping item
func CreateOrUpdateItemHandler(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON format",
			"details": err.Error(),
		})
		return
	}

	if input.Name == "" || input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing or invalid 'name' or 'quantity'",
		})
		return
	}

	item, created, err := storage.CreateOrUpdateItem(input.Name, input.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create/update item"})
		return
	}

	if created {
		c.JSON(http.StatusCreated, item)
	} else {
		c.JSON(http.StatusOK, item)
	}
}