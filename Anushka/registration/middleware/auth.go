package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang-starter/registration/utils"
)

// JWTMiddleware verifies the JWT bearer token in the Authorization header
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

		// 2. Strip the "Bearer " prefix to extract the raw token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. Verify the token signature and expiration
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: " + err.Error(),
			})
			return
		}

		// 4. Attach claims to context so handlers can access the logged-in user
		c.Set("userID", claims["user_id"])
		c.Set("email", claims["email"])

		// 5. Token is valid! Proceed to the protected handler
		c.Next()
	}
}
