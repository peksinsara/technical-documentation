package middleware

import (
	"fmt"
	"techdocs/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs HTTP requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)

		// Get user ID if available
		userID := uint(0)
		if id, exists := c.Get("userID"); exists {
			userID = id.(uint)
		}

		// Log request
		logger.Request(
			c.Request.Method,
			c.Request.URL.Path,
			fmt.Sprintf("%d", c.Writer.Status()),
			duration,
			userID,
		)
	}
}
