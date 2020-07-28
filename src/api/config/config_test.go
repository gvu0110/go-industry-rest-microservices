package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", apiGithubAccessToken)
	assert.EqualValues(t, "info", LogLevel)
	assert.EqualValues(t, "GO_ENVIRONMENT", goEnvironment)
	assert.EqualValues(t, "production", production)
}

// func TestGetGithubAccessToken(t *testing.T) {
// 	assert.EqualValues(t, "", GetGithubAccessToken())
// }
