package middleware

import (
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		reqId,_ := c.Get("reqId")
		fmt.Printf("[%s] %s %s %s\n", reqId, c.Request.Method, c.Request.URL.Path, reqId)
		fmt.Printf("Request Id: %s\n", reqId)
		c.Next()
		duration := time.Since(start)
		status := c.Writer.Status()
		fmt.Printf("[%s] %s %s %v %v\n", reqId, c.Request.Method, c.Request.URL.Path, status, duration)
	}
}