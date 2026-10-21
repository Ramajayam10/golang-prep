package model

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required,min=3"`
	Email     string    `json:"email" binding:"required"`
	Bio       *string   `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
}