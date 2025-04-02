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
