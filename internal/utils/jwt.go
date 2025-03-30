package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	jwtSecretOnce sync.Once
	jwtSecret     []byte
)

func initJWTSecret() {
	jwtSecretOnce.Do(func() {
		// Try to get the secret from environment
		secret := os.Getenv("JWT_SECRET")

		// If not available, generate a new one
		if secret == "" {
			// Generate new random secret
			key := make([]byte, 64)
			if _, err := rand.Read(key); err != nil {
				// Fall back to default if generation fails
				jwtSecret = []byte("your-256-bit-secret")
				return
			}
			secret = base64.StdEncoding.EncodeToString(key)

			// Try to update .env file if it exists
			updateEnvFile(secret)
		}

		jwtSecret = []byte(secret)
	})
}

// Add this function here
// getJWTSecret returns the JWT secret, initializing it if needed
func getJWTSecret() []byte {
	// Initialize JWT secret if not already done
	initJWTSecret()
	return jwtSecret
}

// updateEnvFile attempts to update the JWT_SECRET in the .env file
// updateEnvFile attempts to update the JWT_SECRET in the .env file
func updateEnvFile(secret string) {
	// Get project root directory regardless of where the app is run from
	rootDir, err := getRootDir()
	if err != nil {
		return // Failed to find project root, just use in-memory secret
	}

	// Use absolute paths
	envFilePath := rootDir + "/.env"
	exampleEnvPath := rootDir + "/example.env"

	// Check if .env file exists
	if _, err := os.Stat(envFilePath); os.IsNotExist(err) {
		// Check if example.env exists
		if _, err := os.Stat(exampleEnvPath); err == nil {
			// Copy example.env to .env
			exampleContent, err := ioutil.ReadFile(exampleEnvPath)
			if err != nil {
				return
			}
			ioutil.WriteFile(envFilePath, exampleContent, 0644)
		} else {
			// Create a new .env file
			ioutil.WriteFile(envFilePath, []byte("JWT_SECRET="+secret+"\n"), 0644)
			return
		}
	}

	// Read existing .env file
	content, err := ioutil.ReadFile(envFilePath)
	if err != nil {
		return
	}

	// Update JWT_SECRET line
	contentStr := string(content)
	re := regexp.MustCompile(`(?m)^JWT_SECRET=.*$`)

	if re.MatchString(contentStr) {
		// Replace existing line
		contentStr = re.ReplaceAllString(contentStr, "JWT_SECRET="+secret)
	} else {
		// Add new line
		if strings.HasSuffix(contentStr, "\n") {
			contentStr += "JWT_SECRET=" + secret
		} else {
			contentStr += "\nJWT_SECRET=" + secret
		}
	}

	// Write back to file
	ioutil.WriteFile(envFilePath, []byte(contentStr), 0644)

	// Update environment variable for current process
	os.Setenv("JWT_SECRET", secret)
}

// getRootDir finds the project root directory by looking for the go.mod file
func getRootDir() (string, error) {
	// Start with the current directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Try current directory first
	if _, err := os.Stat(dir + "/go.mod"); err == nil {
		return dir, nil
	}

	// Walk up the directory tree until we find go.mod
	for {
		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root directory without finding go.mod
			return "", errors.New("could not find project root")
		}
		dir = parent

		// Check for go.mod
		if _, err := os.Stat(dir + "/go.mod"); err == nil {
			return dir, nil
		}

		// Also check for example.env as a secondary indicator
		if _, err := os.Stat(dir + "/example.env"); err == nil {
			return dir, nil
		}
	}
}

// GenerateJWT creates a new JWT token for authentication
func GenerateJWT(userID uint, email string, isAdmin bool) (string, error) {
	// Get the signing key from environment
	secret := getJWTSecret()

	// Parse expiry duration from .env
	expiryDuration, err := time.ParseDuration(os.Getenv("JWT_EXPIRY"))
	if err != nil {
		expiryDuration = 24 * time.Hour // Default 24 hours if parsing fails
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"email":    email,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(expiryDuration).Unix(),
	})

	// Sign and return the token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token and returns its claims

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validasi alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token validation failed: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Cek expiration
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("token expired")
			}
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
