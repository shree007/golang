package handlers

import (
	"crud-redis/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteItem(c *gin.Context) {
	config.Rdb.Do(config.Ctx, "SELECT", "0").Err()
	key := c.Param("id")

	err := config.Rdb.Del(config.Ctx, key).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Item deleted"})
}
