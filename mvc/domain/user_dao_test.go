package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	if user != nil {
		t.Error("We're not expecting a user with ID 0")
	}

	if err == nil {
		t.Error("We're expecting an error when user ID is 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("We're expecting 404 when user is not found")
	}
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(1)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Adam", user.FirstName)
	assert.EqualValues(t, "Vu", user.LastName)
	assert.EqualValues(t, "example@gmail.com", user.Email)
}
