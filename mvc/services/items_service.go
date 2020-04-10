package services

import (
	"go-industry-rest-microservices/mvc/domain"
	"go-industry-rest-microservices/mvc/utils"
	"net/http"
)

type itemService struct{}

var ItemService itemService

func (i *itemService) GetItem(itemId int64) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me!",
		StatusCode: http.StatusInternalServerError,
	}
}
