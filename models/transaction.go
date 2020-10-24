package models

import "time"

type Transaction struct {
	TransactionID uint64    `json:"transaction_id" gorm:"primary_key;serial;"`
	From          uint64    `json:"from" `
	To            uint64    `json:"to"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
}
