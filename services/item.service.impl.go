package services

import (
	"assignment-2/models"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ItemServiceImpl struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) ItemService {
	return &ItemServiceImpl{db}
}

func (isi *ItemServiceImpl) GetAllItems() (*models.Item, error) {
	items := models.Item{}
	result := isi.db.Find(&items)
	if result.Error != nil {
		log.Fatal("Error get items data", result.Error)
	}

	return &items, result.Error
}

func (isi *ItemServiceImpl) CreateItem(description string, quantity int32) (*models.Item, error) {
	randomUuid := uuid.NewV4()
	payload := models.Item{
		Description: description,
		Quantity:    quantity,
		ItemCode:    randomUuid,
	}
	err := isi.db.Create(&payload).Error
	if err != nil {
		log.Fatal("Error create items data", err)
	}

	fmt.Println("New item data:", payload)
	return &payload, err
}
