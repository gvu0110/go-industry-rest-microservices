package services

import (
	"go-industry-rest-microservices/src/api/clients/restclient"
	"go-industry-rest-microservices/src/api/domain/repositories"
	"go-industry-rest-microservices/src/api/utils/errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.GetStatusCode())
	assert.EqualValues(t, "Invalid repository name", err.GetMessage())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/v3"}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.GetStatusCode())
	assert.EqualValues(t, "Requires authentication", err.GetMessage())
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1296269, "name": "testing", "owner": {"login": "gvu0110"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 1296269, result.ID)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "gvu0110", result.Owner)
}

func TestCreateRepoConcurrentError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication", "documentation_url": "https://developer.github.com/v3"}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}

	output := make(chan repositories.CreateReposResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.GetStatusCode())
	assert.EqualValues(t, "Requires authentication", result.Error.GetMessage())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1296269, "name": "testing", "owner": {"login": "gvu0110"}}`)),
		},
	})

	request := repositories.CreateRepoRequest{Name: "testing"}

	output := make(chan repositories.CreateReposResult)

	service := reposService{}
	go service.createRepoConcurrent(request, output)

	result := <-output
	assert.NotNil(t, result)
	assert.NotNil(t, result.Response)
	assert.Nil(t, result.Error)
	assert.EqualValues(t, 1296269, result.Response.ID)
	assert.EqualValues(t, "testing", result.Response.Name)
	assert.EqualValues(t, "gvu0110", result.Response.Owner)
}

func TestHandleRepoResult(t *testing.T) {
	inputs := make(chan repositories.CreateReposResult)
	outputs := make(chan repositories.CreateReposResponse)
	defer close(outputs)
	var wg sync.WaitGroup

	service := reposService{}
	go service.handleRepoResult(&wg, inputs, outputs)

	wg.Add(1)
	go func() {
		inputs <- repositories.CreateReposResult{
			Error: errors.NewBadRequestError("Invalid repository name"),
		}
	}()
	wg.Wait()
	close(inputs)

	result := <-outputs
	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)
	assert.EqualValues(t, 1, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.GetStatusCode())
	assert.EqualValues(t, "Invalid repository name", result.Results[0].Error.GetMessage())
}

func TestCreateReposInvalidRequests(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: " "},
	}

	results, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.EqualValues(t, 2, len(results.Results))
	assert.EqualValues(t, http.StatusBadRequest, results.StatusCode)

	assert.Nil(t, results.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, results.Results[0].Error.GetStatusCode())
	assert.EqualValues(t, "Invalid repository name", results.Results[0].Error.GetMessage())

	assert.Nil(t, results.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, results.Results[1].Error.GetStatusCode())
	assert.EqualValues(t, "Invalid repository name", results.Results[1].Error.GetMessage())
}

func TestCreateReposPartialSuccess(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 1296269, "name": "testing", "owner": {"login": "gvu0110"}}`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "testing"},
	}

	results, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.EqualValues(t, 2, len(results.Results))
	assert.EqualValues(t, http.StatusPartialContent, results.StatusCode)

	for _, result := range results.Results {
		if result.Error != nil {
			assert.Nil(t, result.Response)
			assert.EqualValues(t, http.StatusBadRequest, result.Error.GetStatusCode())
			assert.EqualValues(t, "Invalid repository name", result.Error.GetMessage())
			continue
		}
		assert.Nil(t, result.Error)
		assert.EqualValues(t, 1296269, result.Response.ID)
		assert.EqualValues(t, "testing", result.Response.Name)
		assert.EqualValues(t, "gvu0110", result.Response.Owner)
	}
}

func TestCreateReposRepoAlreadyExistsFailure(t *testing.T) {
	restclient.FlushMockups()
	restclient.AddMockups(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		BodyText:   `{"id": 1296269, "name": "testing", "owner": {"login": "gvu0110"}}`,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
		},
	})

	requests := []repositories.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}

	results, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.EqualValues(t, 2, len(results.Results))
	assert.EqualValues(t, http.StatusCreated, results.StatusCode)

	assert.Nil(t, results.Results[0].Error)
	assert.EqualValues(t, 1296269, results.Results[0].Response.ID)
	assert.EqualValues(t, "testing", results.Results[0].Response.Name)
	assert.EqualValues(t, "gvu0110", results.Results[0].Response.Owner)

	assert.Nil(t, results.Results[1].Error)
	assert.EqualValues(t, 1296269, results.Results[1].Response.ID)
	assert.EqualValues(t, "testing", results.Results[1].Response.Name)
	assert.EqualValues(t, "gvu0110", results.Results[1].Response.Owner)
}
