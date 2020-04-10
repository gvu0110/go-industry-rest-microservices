package app

import (
	"go-industry-rest-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:userId", controllers.GetUser)
}
