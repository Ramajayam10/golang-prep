package models

// User represents the database table schema
type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

// LoginRequest holds incoming login credentials
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}