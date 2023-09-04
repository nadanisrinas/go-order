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

func (isi *ItemServiceImpl) FindItem(itemCode string) (models.Item, error) {
	// fmt.Println("itemCode", itemCode)
	item := models.Item{}
	// itemCodeUUID, errUUID := uuid.FromString(itemCode)
	errFindItem := isi.db.Where("item_code = ?", itemCode).Find(&item).Error
	if errFindItem != nil {
		log.Fatal("Error can't find item", errFindItem)
	}

	return item, errFindItem
}

func FindItem(itemCode string) (models.Item, error) {
	db := *&gorm.DB{}
	isi := *&ItemServiceImpl{}
	item, err := isi.GetAllItems()
	fmt.Println("item", item)
	// itemCodeUUID, errUUID := uuid.FromString(itemCode)
	errFindItem := db.Where("item_code = ?", itemCode).Find(&item).Error
	if err != nil || errFindItem != nil {
		log.Fatal("Error can't find item", errFindItem)
	}

	return *item, errFindItem
}
