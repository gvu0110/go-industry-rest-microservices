package repositories

import (
	"go-industry-rest-microservices/src/api/utils/errors"
	"strings"
)

// CreateRepoRequest struct to create a new GitHub repo with the name
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Validate function
func (input *CreateRepoRequest) Validate() errors.APIError {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return errors.NewBadRequestError("Invalid repository name")
	}
	return nil
}

// CreateRepoResponse struct
type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

// CreateReposResponse struct
type CreateReposResponse struct {
	StatusCode int                 `json:"status"`
	Results    []CreateReposResult `json:"results"`
}

// CreateReposResult struct
type CreateReposResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error    errors.APIError     `json:"error"`
}
