// internal/package/controller/file_controller.go
package controller

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileController struct {
	// Tambahkan dependencies jika diperlukan
}

func NewFileController() *FileController {
	return &FileController{}
}

func (c *FileController) UploadProductImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext

	// Save file
	if err := ctx.SaveUploadedFile(file, "./uploads/"+filename); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan file"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"url": "/uploads/" + filename,
	})
}
