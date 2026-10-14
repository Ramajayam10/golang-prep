package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-starter/registration/models"
	"golang-starter/registration/services"
)

// 1. GREETING: Welcome endpoint
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":         "Welcome to the User Registration Portal!",
		"instructions":    "Submit a POST request with name, email, and age to register.",
		"required_fields": []string{"name", "email", "age"},
	})
}

// 2. CREATE: Register a new user
func RegisterUser(c *gin.Context) {
	var newUser models.User

	// Waiter 'c' unpacks and validates the incoming JSON body
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save to PostgreSQL via service layer
	if err := services.RegisterUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully!",
		"user":    newUser,
	})
}

// 3. READ: Get all registered users
func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// 4. READ: Get a single user by ID
func GetUserByID(c *gin.Context) {
	// Waiter 'c' extracts the URL parameter (e.g., /users/:id -> id)
	id := c.Param("id")

	user, err := services.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found with ID: " + id})
		return
	}

	c.JSON(http.StatusOK, user)
}

// 5. UPDATE: Edit existing user information
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedData models.User

	// Validate the updated fields from JSON body
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.UpdateUser(id, updatedData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully!",
		"user":    user,
	})
}

// 6. DELETE: Remove a user by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to delete: user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully!",
		"id":      id,
	})
}
