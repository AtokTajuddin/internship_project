// internal/utils/error_handler.go
package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HandleValidationError(ctx *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errors := make([]ApiError, len(ve))
		for i, fe := range ve {
			errors[i] = ApiError{
				Field:   getJSONFieldName(fe),
				Message: msgForTag(fe.Tag(), fe.Param()),
			}
		}
		ctx.JSON(400, gin.H{"errors": errors})
		return
	}
	HandleError(ctx, 400, "Invalid request")
}

func getJSONFieldName(fe validator.FieldError) string {
	field := fe.Field()
	if fieldType, ok := reflect.TypeOf(fe.Value()).Elem().FieldByName(fe.Field()); ok {
		if jsonTag := fieldType.Tag.Get("json"); jsonTag != "" {
			return strings.Split(jsonTag, ",")[0]
		}
	}
	return field
}

func msgForTag(tag string, param string) string {
	switch tag {
	case "required":
		return "Field ini wajib diisi"
	case "email":
		return "Format email tidak valid"
	case "min":
		return fmt.Sprintf("Minimal %s karakter", param)
	case "max":
		return fmt.Sprintf("Maksimal %s karakter", param)
	case "indonesian_phone":
		return "Format nomor telepon tidak valid. Gunakan format: 08xx, 628xx, atau +628xx"
	case "password_complexity":
		return "Password harus mengandung kombinasi huruf besar, huruf kecil, dan angka (min. 8 karakter)"
	default:
		return "Field tidak valid"
	}
}

func HandleError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"success": false,
		"error":   message,
	})
	ctx.Abort()
}

func HandleSuccess(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(statusCode, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}
