package models

import "time"

type Order struct {
	OrderID      uint      `gorm:"primaryKey" json:"order_id"`
	CustomerName string    `gorm:"not_null" json:"customer_name"`
	OrderedAt    time.Time `gorm:"not_null" json:"ordered_at"`
	CreatedAt    time.Time `gorm:"not_null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not_null" json:"updated_at"`
}
