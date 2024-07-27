package models

import (
	"time"
)

type Vehicle struct {
	ID        uint      `gorm:"primaryKey"`
	Make      string    `json:"make"`
	Model     string    `json:"model"`
	Year      int       `json:"year"`
	Plate     string    `json:"plate"`
	UserID		uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}