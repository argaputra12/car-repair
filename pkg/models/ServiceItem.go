package models

import (
	"time"
)

type ServiceItem struct {
	ID        uint      `gorm:"primaryKey"`
	ServiceID uint      `json:"service_id"`
	Service		Service   `gorm:"foreignKey:ServiceID"`
	ItemID    uint      `json:"item_id"`
	Item			Item      `gorm:"foreignKey:ItemID"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}