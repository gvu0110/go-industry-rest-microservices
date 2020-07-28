package app

import (
	"go-industry-rest-microservices/src/api/log/log_logrus"
	"go-industry-rest-microservices/src/api/log/log_zap"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

// StartApp function
func StartApp() {
	// Interact with the public interface instead of the Log variable
	//log.Log.Info("About to map URLs")
	log_logrus.Info("About to map URLs", "step:1", "status:pending")
	mapUrls()
	//log_logrus.Info("URLs are successfully mapped", "step:2", "status:created")
	log_zap.Info("URLs are successfully mapped", log_zap.Field("step", 2), log_zap.Field("status", "created"))
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
