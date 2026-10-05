package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Item represents the data structure for our CRUD microservice
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

// In-memory data store for demonstration
var items = []Item{
	{ID: "1", Name: "Laptop", Price: 999.99},
	{ID: "2", Name: "Mouse", Price: 29.99},
}

// Health check endpoint (essential for microservices)
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"service": "crud-microservice",
		"status":  "healthy",
	})
}

// CREATE: Add a new item
func createItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items = append(items, newItem)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Item created successfully",
		"data":    newItem,
	})
}

// READ: Get all items
func getAllItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

// READ: Get a single item by ID
func getItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// UPDATE: Update an existing item by ID
func updateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem Item

	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.ID == id {
			// Maintain the original ID
			updatedItem.ID = id
			items[i] = updatedItem
			c.JSON(http.StatusOK, gin.H{
				"message": "Item updated successfully",
				"data":    updatedItem,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// DELETE: Remove an item by ID
func deleteItem(c *gin.Context) {
	id := c.Param("id")

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

func main() {
	// Initialize default Gin engine (includes Logger and Recovery middleware)
	router := gin.Default()

	// Microservice Health Check
	router.GET("/health", healthCheck)

	// API Routes for CRUD operations
	api := router.Group("/api/v1")
	{
		api.GET("/items", getAllItems)       // Read all
		api.GET("/items/:id", getItemByID)   // Read one
		api.POST("/items", createItem)       // Create
		api.PUT("/items/:id", updateItem)    // Update
		api.DELETE("/items/:id", deleteItem) // Delete
	}

	// Start HTTP server on port 8080
	router.Run(":8080")
}
