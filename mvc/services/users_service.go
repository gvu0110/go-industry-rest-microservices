package services

import (
	"go-industry-rest-microservices/mvc/domain"
	"go-industry-rest-microservices/mvc/utils"
)

type userService struct{}

var UserService userService

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
