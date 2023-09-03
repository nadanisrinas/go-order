package services

import (
	"assignment-2/models"
	"fmt"
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
	orderedAt string,
	items []*models.Item,
) (*models.Order, error) {
	// itemCodeUUID, errUUID := uuid.FromString("ec9232c6-7e93-40ce-867a-2974b216a0f9")
	const layout = "2006-01-02"
	orderAtTimeParseFromInput, err := time.Parse(layout, orderedAt)

	// if errUUID != nil {
	// 	log.Fatal("Error create items data", errUUID)
	// }
	payload := models.OrderRequestBody{
		OrderedAt: orderAtTimeParseFromInput.String(),
		Items:     items,
	}

	orders := models.Order{}

	err = osi.db.Create(&payload).Error
	if err != nil {
		log.Fatal("Error create items data", err)
	}

	fmt.Println("New item data:", payload)
	return &orders, err
}
