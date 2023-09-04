package services

import (
	"assignment-2/models"
	"log"
	"time"

	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return &OrderServiceImpl{db}
}

func (osi *OrderServiceImpl) GetAllOrders() (*models.Order, error) {
	orders := models.Order{}
	result := osi.db.Find(&orders)
	if result.Error != nil {
		log.Fatal("Error get orders data", result.Error)
	}

	return &orders, result.Error
}

func (osi *OrderServiceImpl) CreateOrder(
	customerName string,
	orderedAt string,
	items []*models.Item,
) (*models.Order, error) {
	const layout = "01-02-2006"
	orderAtTimeParseFromInput, err := time.Parse(layout, orderedAt)

	payload := models.Order{
		OrderedAt:    orderAtTimeParseFromInput,
		CustomerName: customerName,
		Items:        items,
	}

	orders := models.Order{}

	err = osi.db.Create(&payload).Error
	if err != nil {
		log.Fatal("Error create items data", err)
	}

	return &orders, err

}

func (osi *OrderServiceImpl) DeleteOrder(id int) (*models.Order, error) {
	orders := &models.Order{}
	result := osi.db.First(orders, id)
	if result.Error != nil {
		log.Fatal("Error get orders data", result.Error)
	}
	result = osi.db.Delete(orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, result.Error
}

func (osi *OrderServiceImpl) UpdateOrder(id int, customerName string, orderedAt string, items []*models.Item) (*models.Order, error) {
	orders := &models.Order{}
	const layout = "01-02-2006"
	orderAtTimeParseFromInput, err := time.Parse(layout, orderedAt)
	if err != nil {
		log.Fatal("Errorparsing data", err)
	}
	result := osi.db.First(orders, id)
	if result.Error != nil {
		log.Fatal("Error get orders data", result.Error)
	}

	payload := models.Order{
		OrderedAt:    orderAtTimeParseFromInput,
		CustomerName: customerName,
		Items:        items,
	}
	result = osi.db.Save(&payload)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, result.Error
}
