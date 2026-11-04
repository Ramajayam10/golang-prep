package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Age         int       `json:"age" binding:"required"`
	Email       string    `gorm:"uniqueIndex" json:"email" binding:"required,email"`
	Address     string    `json:"address"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
