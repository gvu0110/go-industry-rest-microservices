package app

import (
	"go-industry-rest-microservices/src/api/controllers/polo"
	"go-industry-rest-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/macro", polo.Macro)
	router.POST("/repositories", repositories.CreateRepo)
}
