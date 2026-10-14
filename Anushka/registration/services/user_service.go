package services

import (
	"golang-starter/registration/config"
	"golang-starter/registration/models"
)


// Create User
func RegisterUser(user *models.User) error {
	result := config.DB.Create(user)
	return result.Error
}

// Get all User 
func GetAllUser() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

// Get user by id 
func GetUserById(id string) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return user, result.Error}

// UpdateUser finds a user and updates their information
func UpdateUser(id string, updatedData models.User) (models.User, error) {
	var user models.User

	// Step 1: Check if the user exists
	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	// Step 2: Update the fields in PostgreSQL
	config.DB.Model(&user).Updates(updatedData)
	return user, nil
}


// DeleteUser
func DeleteUser(id string) error {
	var user models.User

	// Step 1: Check if the user exists
	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}

	// Step 2: Delete from the database
	result := config.DB.Delete(&user)
	return result.Error
}
