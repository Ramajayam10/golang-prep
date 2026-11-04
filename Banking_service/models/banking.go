package models

import "time"

type BankAccount struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	AccountNumber string    `gorm:"uniqueIndex" json:"account_number" binding:"required"`
	BankName      string    `json:"bank_name" binding:"required"`
	UserID        uint      `gorm:"not null;index" json:"user_id" binding:"required"`
	Balance       float64   `gorm:"default:0" json:"balance"`
	IsActive      bool      `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
