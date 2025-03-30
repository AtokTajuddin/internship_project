package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new server
	server := NewServer()

	// Ensure server is not nil
	assert.NotNil(t, server)

	// Test that the router is initialized
	assert.NotNil(t, server.Router)

	// Add a simple handler for testing
	server.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Create a test request
	req := httptest.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	// Serve the request
	server.Router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}
