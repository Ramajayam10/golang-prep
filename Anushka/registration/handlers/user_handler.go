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
		"instructions":    "Submit a POST request with name, email, age, and password to register.",
		"required_fields": []string{"name", "email", "age", "password"},
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

	// Save to PostgreSQL via service layer (hashes password)
	if err := services.RegisterUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	// Hide password hash in response
	newUser.Password = ""

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully!",
		"user":    newUser,
	})
}

// 3. AUTH: User Login (verifies password and issues JWT token)
func Login(c *gin.Context) {
	var loginReq models.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := services.LoginUser(loginReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   token,
		"user":    user,
	})
}

// 4. READ: Get all registered users (Protected)
func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// 5. READ: Get a single user by ID (Protected)
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := services.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found with ID: " + id})
		return
	}

	c.JSON(http.StatusOK, user)
}

// 6. UPDATE: Edit existing user information (Protected)
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedData models.User

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

// 7. DELETE: Remove a user by ID (Protected)
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
