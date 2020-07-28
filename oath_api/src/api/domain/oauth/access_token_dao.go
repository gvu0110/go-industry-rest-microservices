package oauth

import (
	"fmt"
	"go-industry-rest-microservices/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

// Save function
func (at *AccessToken) Save() errors.APIError {
	at.Token = fmt.Sprintf("USER_%d", at.UserID)
	tokens[at.Token] = at
	return nil
}

// GetAccessToken function
func GetAccessToken(accessToken string) (*AccessToken, errors.APIError) {
	token := tokens[accessToken]
	if token == nil || token.IsExpired() {
		return nil, errors.NewNotFoundError("No access token found with given parameters")
	}
	return token, nil
}
