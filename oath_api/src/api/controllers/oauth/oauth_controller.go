package oauth

import (
	"go-industry-rest-microservices/oath_api/src/api/domain/oauth"
	"go-industry-rest-microservices/oath_api/src/api/services"
	"go-industry-rest-microservices/src/api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAccessToken function
func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(apiError.GetStatusCode(), apiError)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

// GetAccessToken function
func GetAccessToken(c *gin.Context) {
	token, err := services.OauthService.GetAccessToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
