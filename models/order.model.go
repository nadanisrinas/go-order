package models

import "time"

type Order struct {
	Id           uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `gorm:"not_null" json:"customer_name"`
	OrderedAt    time.Time `gorm:"not_null" json:"ordered_at"`
	// Items        []*Item   `gorm:"foreignkey:id; association_foreignkey:id"`
	// CreatedAt    time.Time `gorm:"not_null" json:"created_at"`
	// UpdatedAt    time.Time `gorm:"not_null" json:"updated_at"`
}

type OrderRequestBody struct {
	// *Order
	Id        uint    `gorm:"primaryKey" json:"order_id"`
	OrderedAt string  `gorm:"not_null" json:"ordered_at"`
	Items     []*Item `gorm:"foreignkey:id; association_foreignkey:id" json:"items"`
}
