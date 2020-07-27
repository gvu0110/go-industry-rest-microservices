package services

import (
	"go-industry-rest-microservices/src/api/config"
	"go-industry-rest-microservices/src/api/domain/github"
	"go-industry-rest-microservices/src/api/domain/providers/github_provider"
	"go-industry-rest-microservices/src/api/domain/repositories"
	"go-industry-rest-microservices/src/api/utils/errors"
	"net/http"
	"sync"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
	CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError)
}

// RepositoryService variable
var RepositoryService reposServiceInterface

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewAPIError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APIError) {
	inputs := make(chan repositories.CreateReposResult)
	outputs := make(chan repositories.CreateReposResponse)
	defer close(outputs)

	var wg sync.WaitGroup
	go s.handleRepoResult(&wg, inputs, outputs)

	for _, request := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(request, inputs)
	}

	wg.Wait()
	close(inputs)

	results := <-outputs
	successfulCreations := 0
	for _, result := range results.Results {
		if result.Response != nil {
			successfulCreations++
		}
	}
	if successfulCreations == 0 {
		results.StatusCode = results.Results[0].Error.GetStatusCode()
	} else if successfulCreations == len(requests) {
		results.StatusCode = http.StatusCreated
	} else {
		results.StatusCode = http.StatusPartialContent
	}
	return results, nil
}

func (s *reposService) handleRepoResult(wg *sync.WaitGroup, inputs chan repositories.CreateReposResult, outputs chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for input := range inputs {
		results.Results = append(results.Results, input)
		wg.Done()
	}

	outputs <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, inputs chan repositories.CreateReposResult) {
	result, err := s.CreateRepo(input)
	if err != nil {
		inputs <- repositories.CreateReposResult{
			Error: err,
		}
		return
	}
	inputs <- repositories.CreateReposResult{
		Response: result,
	}
}
