package services

import (
	"assignment-2/models"
)

type OrderService interface {
	GetAllOrders() (*models.Order, error)
	CreateOrder(orderedAt string,
		items []*models.Item,
	) (*models.Order, error)
}
