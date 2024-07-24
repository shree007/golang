package main

import (
	"crud-redis/config"
	"crud-redis/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitRedis()
	router := gin.Default()

	router.POST("/db0", handlers.CreateItem)
	router.GET("/db0/:id", handlers.GetItem)
	router.PUT("/db0/:id", handlers.UpdateItem)
	router.DELETE("/db0/:id", handlers.DeleteItem)

	router.Run(":8080")
}
