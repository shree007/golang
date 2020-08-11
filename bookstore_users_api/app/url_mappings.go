package app

import (
	"bookstore_users_api/controllers/healthCheck"
	"bookstore_users_api/controllers/userController"
)

func mapUrls() {
	router.GET("/ping", healthCheck.HealthCheck)
	router.GET("/users/:user_id", userController.GetUser)
	// router.GET("/user/search", controllers.SearchUser) Implement Me
	router.POST("/users", userController.CreateUser)

}
