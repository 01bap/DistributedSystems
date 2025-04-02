package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"app/storage"
)

func GetItemsHandler(c *gin.Context) {
	items, err := storage.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func CreateOrUpdateItemHandler(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
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