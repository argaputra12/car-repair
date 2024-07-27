package models

import (
	"time"
)

type ServiceRecord struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	VehicleID   uint      `json:"vehicle_id"`
	PaymentID   uint      `json:"payment_id"`
	ServiceID   uint      `json:"service_id"`
	ServiceDate time.Time `json:"service_date"`
	Notes       string    `json:"notes"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
