package models

import (
	"time"
)

type Service struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description *string    `json:"description"`
	TotalPrice  *float64   `json:"total_price"`
	CreatedAt	 time.Time `json:"created_at"`
	UpdatedAt	 time.Time `json:"updated_at"`
}