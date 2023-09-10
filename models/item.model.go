package models

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Item struct {
	ID          uint      `gorm:"primaryKey" json:"-"`
	ItemID      uint      `gorm:"not_null" json:"-"`
	ItemCode    uuid.UUID `gorm:"not_null" json:"item_code"`
	Description string    `gorm:"not_null" json:"description"`
	Quantity    int32     `gorm:"not_null" json:"quantity"`
	OrderId     int64     `gorm:"foreignkey:fk_orders_items;" json:"order_id"`
}

type ItemResponse struct {
	ItemID      uint      `json:"item_id"`
	ItemCode    uuid.UUID `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int32     `json:"quantity"`
}

type ItemRequestBody struct {
	Description string
	Quantity    int32
}

// hook BeforeCreate will set a UUID rather than numeric ID.
func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	uuidItemCode := uuid.NewV4()
	item.ItemCode = uuidItemCode
	if len(item.Description) < 4 {
		log.Fatal("item description is too short")
	}
	return
}
