package config

import (
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel             = "info"
	goEnvironment        = "GO_ENVIRONMENT"
	production           = "production"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

// GetGithubAccessToken function
func GetGithubAccessToken() string {
	return githubAccessToken
}

// IsProduction function
func IsProduction() bool {
	return os.Getenv(goEnvironment) == production
}

// Put in Dockerfile
// ENV GO_ENVIRONMENT=production
// ENV LOG_LEVEL=info
