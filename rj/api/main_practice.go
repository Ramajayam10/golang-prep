package main

import "github.com/gin-gonic/gin"

type User struct {
	Name string `json:"name" binding:"required"`
	Age int `json:"age" binding:"required,min=18"`
}

func mmain() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/users", func (c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, gin.H{
			"data": name,
		})
	})
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})
	router.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(201, user)
	})
	router.Run("localhost:8080")
}