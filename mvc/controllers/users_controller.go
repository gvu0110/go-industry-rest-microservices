package controllers

import (
	"encoding/json"
	"go-industry-rest-microservices/mvc/services"
	"go-industry-rest-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdParam := r.URL.Query().Get("userId")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		// Just return the Bad Request to the client
		apiErr := &utils.ApplicationError{
			Message:    "userId must be a number!",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		w.WriteHeader(apiErr.StatusCode)
		w.Write(jsonValue)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		w.WriteHeader(apiErr.StatusCode)
		w.Write([]byte(apiErr.Message))
		return
	}

	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
