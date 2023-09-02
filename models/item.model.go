package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ItemId      uint      `gorm:"primaryKey" json:"item_id"`
	ItemCode    uuid.UUID `gorm:"not_null" json:"item_code"`
	Description string    `gorm:"not_null" json:"description"`
	Quantity    int32     `gorm:"not_null" json:"quantity"`
	OrderID     Order     `gorm:"foreignKey:OrderID" json:"-"`
	CreatedAt   time.Time `gorm:"not_null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not_null" json:"updated_at"`
}

type ItemResponse struct {
	ItemID      uint      `gorm:"primaryKey" json:"item_id"`
	ItemCode    uuid.UUID `gorm:"not null" json:"item_code"`
	Description string    `gorm:"not_null" json:"description"`
	Quantity    int32     `gorm:"not_null" json:"quantity"`
	OrderID     Order     `gorm:"foreignKey:OrderID" json:"-"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// hook BeforeCreate will set a UUID rather than numeric ID.
func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	uuidItemCode := uuid.NewV4()
	item.ItemCode = uuidItemCode
	return
}
