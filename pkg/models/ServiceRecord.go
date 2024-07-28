package models

import (
	"time"
)

type ServiceRecord struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	User				User      `gorm:"foreignKey:UserID"`
	VehicleID   uint      `json:"vehicle_id"`
	Vehicle			Vehicle   `gorm:"foreignKey:VehicleID"`
	PaymentID   uint      `json:"payment_id"`
	Payment			Payment   `gorm:"foreignKey:PaymentID"`
	ServiceID   uint      `json:"service_id"`
	Service			Service   `gorm:"foreignKey:ServiceID"`
	ServiceDate time.Time `json:"service_date"`
	Notes       string    `json:"notes"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
