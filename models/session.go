package models

import "time"

type Session struct {
	UserID      uint64    `json:"user_id" gorm:"foreign_key:'users(user_id)';unique;"`
	AccessToken string    `json:"access_token"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
}
