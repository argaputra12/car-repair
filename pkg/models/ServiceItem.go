package models

import (
	"time"
)

type ServiceItem struct {
	ID        uint      `gorm:"primaryKey"`
	ServiceID uint      `json:"service_id"`
	ItemID    uint      `json:"item_id"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}