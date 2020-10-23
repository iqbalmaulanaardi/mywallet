package models

import "time"

type User struct {
	UserID    uint64    `json:"user_id" gorm:"primary_key;serial;"`
	Username  string    `json:"username" gorm:"unique;"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP"`
}
