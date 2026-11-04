package models

import "time"

type Transaction struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	UserID            uint      `gorm:"not null;index" json:"user_id" binding:"required"`
	SenderAccountID   uint      `gorm:"not null;index" json:"sender_account_id" binding:"required"`
	ReceiverAccountID uint      `gorm:"not null;index" json:"receiver_account_id" binding:"required"`
	ReceiverName      string    `json:"receiver_name"`
	Amount            float64   `gorm:"not null" json:"amount" binding:"required,gt=0"`
	TransactionType   string    `gorm:"not null" json:"transaction_type" binding:"required"`
	Source            string    `gorm:"not null" json:"source" binding:"required"`          
	Status            string    `gorm:"not null;default:'PENDING'" json:"status"`          
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
}
