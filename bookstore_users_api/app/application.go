package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	// starting the application on 8080 port
	mapUrls()
	router.Run(":8080")
}
