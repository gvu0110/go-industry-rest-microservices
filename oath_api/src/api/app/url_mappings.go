package app

import (
	"go-industry-rest-microservices/oath_api/src/api/controllers/oauth"
	"go-industry-rest-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/macro", polo.Macro)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
