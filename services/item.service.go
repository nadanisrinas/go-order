package services

import "assignment-2/models"

type ItemService interface {
	GetAllItems() (*models.Item, error)
	CreateItem(description string, quantity int32) (*models.Item, error)
}
