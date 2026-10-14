package models 

type User struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `gorm:"unique" json:"email" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}