package handlers

import (
	"crud-redis/config"
	"crud-redis/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateItem(c *gin.Context) {
	config.Rdb.Do(config.Ctx, "SELECT", "0").Err()
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := config.Rdb.Set(config.Ctx, item.Key, item.Value, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Item updated"})
}
