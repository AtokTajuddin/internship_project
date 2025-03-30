// internal/middleware/auth.go
package middleware

import (
	"project_virtual_internship_evermos/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header required"})
			return
		}

		// Pastikan format Bearer token
		if !strings.HasPrefix(tokenHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization format. Use Bearer <token>"})
			return
		}

		// Ekstrak token
		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

		// Validasi token
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			// Tambahkan pesan error spesifik
			c.AbortWithStatusJSON(401, gin.H{
				"error":  "Invalid token",
				"detail": err.Error(), // Hanya untuk development, disable di production
			})
			return
		}

		// Cek existence user_id dalam claims
		userID, ok := claims["user_id"]
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token missing user_id claim"})
			return
		}

		// Konversi user_id ke uint dengan aman
		userIDFloat, ok := userID.(float64)
		if !ok {
			c.AbortWithStatusJSON(401, gin.H{
				"error":         "Invalid user_id format in token",
				"expected_type": "number",
			})
			return
		}
		userIDUint := uint(userIDFloat)

		// Set userID ke context
		c.Set("userID", userIDUint)
		c.Next()
	}
}
