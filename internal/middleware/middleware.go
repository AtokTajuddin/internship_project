package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// CORS middleware provides Cross-Origin Resource Sharing support
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// RequestLogger logs request details like HTTP method, path, status code, and processing time
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start time
		startTime := time.Now()

		// Process request
		c.Next()

		// End time
		endTime := time.Now()

		// Request details
		latency := endTime.Sub(startTime)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		// Different log colors based on status code
		var statusColor string
		switch {
		case status >= 400:
			statusColor = "\033[31m" // Red
		case status >= 300:
			statusColor = "\033[33m" // Yellow
		default:
			statusColor = "\033[32m" // Green
		}

		// Reset color
		resetColor := "\033[0m"

		// Log the request
		fmt.Printf("[GIN] %v |%s %3d %s| %13v | %s %s %s\n",
			endTime.Format("2006/01/02 - 15:04:05"),
			statusColor, status, resetColor,
			latency,
			method,
			path,
			c.Request.UserAgent(),
		)
	}
}
