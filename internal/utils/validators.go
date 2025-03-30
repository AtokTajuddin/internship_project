// internal/utils/validators.go
package utils

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func IndonesianPhone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()

	// Bersihkan karakter non-digit kecuali +
	cleaned := strings.ReplaceAll(phone, "-", "")
	cleaned = strings.ReplaceAll(cleaned, " ", "")

	// Regex pattern untuk nomor Indonesia
	pattern := `^(\+62|62|0)8[1-9][0-9]{6,9}$`

	match, _ := regexp.MatchString(pattern, cleaned)
	return match
}

func PasswordComplexity(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Minimal 8 karakter
	if len(password) < 8 {
		return false
	}

	// Harus mengandung:
	hasLower := false
	hasUpper := false
	hasNumber := false

	for _, c := range password {
		switch {
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsNumber(c):
			hasNumber = true
		}
	}

	return hasLower && hasUpper && hasNumber
}

func RegisterCustomValidations(v *validator.Validate) {
	v.RegisterValidation("indonesian_phone", IndonesianPhone)
	v.RegisterValidation("password_complexity", PasswordComplexity)
}
