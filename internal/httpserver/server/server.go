package server

import (
	"project_virtual_internship_evermos/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupServer() *gin.Engine {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		utils.RegisterCustomValidations(v)
	}

	return router
}

// Server represents the HTTP server
type Server struct {
	Router *gin.Engine
}

// NewServer creates and returns a new server instance with configured router
func NewServer() *Server {
	router := gin.Default()
	return &Server{
		Router: router,
	}
}
