package services

import (
	"assignment-2/models"
)

type OrderService interface {
	GetAllOrders() (*models.Order, error)
	CreateOrder(customerName string, orderedAt string, items []*models.Item) (*models.Order, error)
	DeleteOrder(id int) (*models.Order, error)
	UpdateOrder(id int, customerName string, orderedAt string, items []*models.Item) (*models.Order, error)
}
