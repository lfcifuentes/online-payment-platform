package rules

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CustomPasswordValidator checks if the password has at least 8 characters, at most 25 characters, at least one number and at least one special character
func CustomPasswordValidator(fl validator.FieldLevel) bool {
	// Obtener el valor del campo
	val := fl.Field().String()

	// check length
	if len(val) < 8 || len(val) > 25 {
		return false
	}

	// Check if it has at least one number
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(val)
	if !hasNumber {
		return false
	}

	// Check if it has at least one special character
	hasSpecialChar := regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(val)
	if !hasSpecialChar {
		return false
	}

	return true
}

// CustomPasswordConfirmationValidator checks if the password confirmation matches the password
func CustomPasswordConfirmationValidator(fl validator.FieldLevel) bool {
	field := fl.Parent().FieldByName("Password").String()
	confirmation := fl.Field().String()

	return confirmation == field
}
