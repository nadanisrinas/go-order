package services

import "assignment-2/models"

type ItemService interface {
	GetAllItems() (*models.Item, error)
	CreateItem(description string, quantity int32, orderId int32) (*models.Item, error)
	FindItem(itemCode string) (models.Item, error)
}
