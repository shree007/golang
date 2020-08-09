package app

import (
	"bookstore_users_api/controllers/health_check"
	"bookstore_users_api/controllers/user_controller"
)

func mapUrls() {
	router.GET("/ping", health_check.HealthCheck)
	router.GET("/users/:user_id", user_controller.GetUser)
	// router.GET("/user/search", controllers.SearchUser)
	router.POST("/users", user_controller.CreateUser)

}
