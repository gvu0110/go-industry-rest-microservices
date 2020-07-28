package oauth

import (
	"go-industry-rest-microservices/src/api/utils/errors"
	"strings"
)

// AccessTokenRequest struct
type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate function
func (r *AccessTokenRequest) Validate() errors.APIError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.NewBadRequestError("Invalid username")
	}

	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.NewBadRequestError("Invalid password")
	}
	return nil
}
