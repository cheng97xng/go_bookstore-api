package app

import (
	"go_bookstore-api/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
