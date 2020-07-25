package repositories

import (
	"go-industry-rest-microservices/src/api/domain/repositories"
	"go-industry-rest-microservices/src/api/services"
	"go-industry-rest-microservices/src/api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRepo function
func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(apiError.GetStatusCode(), apiError)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
