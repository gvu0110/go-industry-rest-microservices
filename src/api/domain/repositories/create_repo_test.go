package repositories

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidRequest(t *testing.T) {
	input := CreateRepoRequest{}
	err := input.Validate()
	assert.EqualValues(t, http.StatusBadRequest, err.GetStatusCode())
	assert.EqualValues(t, "Invalid repository name", err.GetMessage())
}

func TestValidRequest(t *testing.T) {
	input := CreateRepoRequest{
		Name: "testing",
	}
	err := input.Validate()
	assert.Nil(t, err)
}
