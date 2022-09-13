package app

import (
	"go_bookstore-api/controllers"
)

// Map All Url at this file.
func mapUrls() {
	// Ping Controller
	router.GET("/ping", controllers.Ping)

	//User Controller
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUser)
	router.PUT("/users/:user_id", controllers.UpdateUser)
	router.PATCH("/users/:user_id", controllers.UpdateUser)
}
