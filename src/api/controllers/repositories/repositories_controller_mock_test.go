package repositories

import (
	"encoding/json"
	"go-industry-rest-microservices/src/api/domain/repositories"
	"go-industry-rest-microservices/src/api/services"
	"go-industry-rest-microservices/src/api/utils/errors"
	"go-industry-rest-microservices/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	funcCreateRepo  func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
	funcCreateRepos func(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError)
)

type repoServiceMock struct{}

func init() {
	services.RepositoryService = &repoServiceMock{}
}

func (s *repoServiceMock) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	return funcCreateRepo(request)
}

func (s *repoServiceMock) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError) {
	return funcCreateRepos(request)
}

func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	funcCreateRepo = func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
		return &repositories.CreateRepoResponse{
			ID:    1296269,
			Name:  "testing",
			Owner: "gvu0110",
		}, nil
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 1296269, result.ID)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "gvu0110", result.Owner)
}

func TestCreateRepoErrorFromGithubMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
		return nil, errors.NewBadRequestError("Invalid repository name")
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewAPIErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.GetStatusCode())
	assert.EqualValues(t, "Invalid repository name", apiErr.GetMessage())
}
