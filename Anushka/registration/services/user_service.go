package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"golang-starter/registration/config"
	"golang-starter/registration/models"
	"golang-starter/registration/utils"
)

// RegisterUser hashes password and saves new user record
func RegisterUser(user *models.User) error {
	// 1. Hash the plain-text password securely with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// 2. Insert into PostgreSQL
	result := config.DB.Create(user)
	return result.Error
}

// LoginUser verifies email & password, and returns a signed JWT token
func LoginUser(req models.LoginRequest) (string, models.User, error) {
	var user models.User

	// 1. Check if user with this email exists
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return "", user, errors.New("invalid email or password")
	}

	// 2. Compare the plain-text password with the hashed password in PostgreSQL
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", user, errors.New("invalid email or password")
	}

	// 3. Password matches! Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", user, errors.New("failed to generate authentication token")
	}

	// Remove password hash from returned object for security
	user.Password = ""
	return token, user, nil
}

// GetAllUser fetches all registered users
func GetAllUser() ([]models.User, error) {
	var users []models.User
	result := config.DB.Omit("password").Find(&users)
	return users, result.Error
}

// GetUserById fetches a single user by primary key
func GetUserById(id string) (models.User, error) {
	var user models.User
	result := config.DB.Omit("password").First(&user, id)
	return user, result.Error
}

// UpdateUser finds a user and updates their information
func UpdateUser(id string, updatedData models.User) (models.User, error) {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	// If updating password, hash it first
	if updatedData.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		if err == nil {
			updatedData.Password = string(hashed)
		}
	}

	config.DB.Model(&user).Updates(updatedData)
	user.Password = ""
	return user, nil
}

// DeleteUser removes a user record from PostgreSQL by ID
func DeleteUser(id string) error {
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}

	result := config.DB.Delete(&user)
	return result.Error
}
