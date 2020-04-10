package controllers

import (
	"go-industry-rest-microservices/mvc/services"
	"go-industry-rest-microservices/mvc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		// Just return the Bad Request to the client
		apiErr := &utils.ApplicationError{
			Message:    "userId must be a number!",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
