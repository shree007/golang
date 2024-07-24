package handlers

import (
	"crud-redis/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func GetItem(c *gin.Context) {
	key := c.Param("key")
	value, err := config.Rdb.Get(config.Ctx, key).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}
