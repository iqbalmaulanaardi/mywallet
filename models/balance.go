package models

import "time"

type Balance struct {
	UserID    uint64    `json:"user_id" gorm:"foreign_key:'users(user_id)';unique;"`
	Balance   float64   `json:"balance" `
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
}
