// internal/utils/validators_passwd_test.go
package utils

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestPasswordComplexity(t *testing.T) {
	testCases := []struct {
		Password string
		Expected bool
		Desc     string
	}{
		{
			Password: "Passw0rd",
			Expected: true,
			Desc:     "Valid: kombinasi huruf besar, kecil, dan angka",
		},
		{
			Password: "P@ssw0rd!",
			Expected: true,
			Desc:     "Valid: mengandung special character (tidak wajib)",
		},
		{
			Password: "password",
			Expected: false,
			Desc:     "Invalid: tanpa huruf besar & angka",
		},
		{
			Password: "PASSWORD",
			Expected: false,
			Desc:     "Invalid: tanpa huruf kecil & angka",
		},
		{
			Password: "Pass123",
			Expected: false,
			Desc:     "Invalid: panjang 7 karakter",
		},
		{
			Password: "Passwords",
			Expected: false,
			Desc:     "Invalid: tanpa angka",
		},
		{
			Password: "P4ssword",
			Expected: true,
			Desc:     "Valid: kombinasi lengkap",
		},
		{
			Password: "12345678",
			Expected: false,
			Desc:     "Invalid: tanpa huruf",
		},
		{
			Password: "Abcd1234",
			Expected: true,
			Desc:     "Valid: kombinasi lengkap",
		},
	}

	v := validator.New()
	RegisterCustomValidations(v)

	for _, tc := range testCases {
		t.Run(tc.Desc, func(t *testing.T) {
			err := v.Var(tc.Password, "password_complexity")
			assert.Equal(t, tc.Expected, err == nil,
				"Password: %s, Expected: %v, Actual: %v",
				tc.Password, tc.Expected, err == nil)
		})
	}
}

// Tambahkan test untuk edge cases
func TestPasswordComplexityEdgeCases(t *testing.T) {
	v := validator.New()
	RegisterCustomValidations(v)

	// Test empty password
	assert.Error(t, v.Var("", "password_complexity"))

	// Test minimal length
	assert.NoError(t, v.Var("Aa1aaaaa", "password_complexity")) // 8 karakter
	assert.Error(t, v.Var("Aa1aaa", "password_complexity"))     // 7 karakter

	// Test unicode characters
	assert.NoError(t, v.Var("Pässwörd1", "password_complexity"))
}
