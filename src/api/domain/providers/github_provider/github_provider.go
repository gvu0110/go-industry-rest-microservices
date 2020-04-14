package github_provider

import (
	"encoding/json"
	"fmt"
	"go-industry-rest-microservices/src/api/clients/restclient"
	"go-industry-rest-microservices/src/api/domain/github"
	"io/ioutil"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateaRepo            = "https://api.github.com/users/repos"
)

func getAuthorizationHeader(token string) string {
	return fmt.Sprintf(headerAuthorizationFormat, token)
}

func CreateRepo(token string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(token))
	response, err := restclient.Post(urlCreateaRepo, request, headers)
	if err != nil {
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GitHubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GitHubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid JSON response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when trying to unmarshal GitHub JSON repo creation response",
		}
	}
	return &result, nil
}
