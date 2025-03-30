// internal/http/handler/file_handler.go
package handler

import (
	"project_virtual_internship_evermos/internal/package/controller"

	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.Engine, fileController *controller.FileController) {
	fileGroup := r.Group("/files")
	{
		fileGroup.POST("/upload", fileController.UploadProductImage)
	}
}
