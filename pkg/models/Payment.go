package models

import (
	"time"
)

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	Amount        float64   `json:"amount"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
	// Status (Paid, Unpaid, Partial)
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
