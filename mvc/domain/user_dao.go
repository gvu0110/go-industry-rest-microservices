package domain

import (
	"fmt"
	"go-industry-rest-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		1: {Id: 1, FirstName: "Adam", LastName: "Vu", Email: "example@gmail.com"},
	}
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

func init() {
	UserDao = &userDao{}
}

type userDao struct{}

var UserDao userDaoInterface

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("We're accessing the actual database!")
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v is not found!", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
