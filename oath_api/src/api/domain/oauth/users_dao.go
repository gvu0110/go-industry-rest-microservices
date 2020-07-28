package oauth

import (
	"go-industry-rest-microservices/src/api/utils/errors"
)

const (
	queryUserByUsernameAndPassword = ""
)

// Mock database
var (
	users = map[string]*User{
		"gvu": {
			ID:       123,
			Username: "gvu",
		},
	}
)

// GetUserByUsernameAndPassword function
func GetUserByUsernameAndPassword(username string, password string) (*User, errors.APIError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundError("No user found with given parameters")
	}
	return user, nil
}
