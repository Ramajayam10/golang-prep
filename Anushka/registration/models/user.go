package models
import "strings"
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


// Check if a user is eligible to register (Age >= 18 and Email not empty)
func (u User) IsEligible() bool {
    return u.Age >= 18 && u.Age <= 120 && strings.Contains(u.Email, "@")
}