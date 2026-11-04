package middleware

import (
	"net/http"
	"strings"

	"banking-service/utils"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware verifies the JWT bearer token on protected banking routes
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Missing or invalid Authorization header (format: Bearer <token>)",
			})
			return
		}

		// 2. Extract the raw token string (remove "Bearer ")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. Verify the token signature and expiration
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: " + err.Error(),
			})
			return
		}

		// 4. Attach claims to context so all protected handlers know the logged-in user!
		c.Set("userID", claims["user_id"])
		c.Set("email", claims["email"])

		// 5. Token is valid! Proceed to the actual handler
		c.Next()
	}
}
