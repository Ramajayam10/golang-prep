package model

type CreateUserRequest struct {
	Name      string    `json:"name" binding:"required,min=3"`
	Email     string    `json:"email" binding:"required,email,max=255"`
	Bio       *string   `json:"bio" binding:"omitempty,max=1000"`
}