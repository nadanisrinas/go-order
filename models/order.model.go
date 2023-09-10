package models

import (
	"time"
)

type Order struct {
	ID           uint      `gorm:"column:id;primaryKey" json:"id"`
	CustomerName string    `gorm:"not_null" json:"customer_name"`
	OrderedAt    time.Time `gorm:"not_null" json:"ordered_at"`
	// Items        []*Item   `gorm:"Foreignkey:ItemID" json:"items"`
}

type OrderCreateRequestBody struct {
	// *Order
	OrderID      uint    `gorm:"primaryKey" json:"order_id"`
	OrderedAt    string  `gorm:"not_null" json:"ordered_at"`
	CustomerName string  `gorm:"not_null" json:"customer_name"`
	Items        []*Item `json:"items"`
}
