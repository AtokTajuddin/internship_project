// internal/middleware/ownership.go
package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserOwnership(resourceID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		requestedID, _ := strconv.Atoi(c.Param(resourceID))

		if uint(requestedID) != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "akses tidak diizinkan"})
			c.Abort()
			return
		}
		c.Next()
	}
}
