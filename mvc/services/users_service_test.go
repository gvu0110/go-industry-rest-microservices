package services

import (
	"go-industry-rest-microservices/mvc/domain"
	"go-industry-rest-microservices/mvc/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type usersDaoMock struct{}

var userDaoMock usersDaoMock
var getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)

func init() {
	domain.UserDao = &userDaoMock
}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "User 0 does not exist!",
		}
	}

	user, err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 does not exist!", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}

	user, err := UserService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
