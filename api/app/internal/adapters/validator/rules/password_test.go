package rules

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestCustomPasswordValidator(t *testing.T) {
	validate := validator.New()

	// Registrar la validación custom
	validate.RegisterValidation("custom_password", CustomPasswordValidator)

	tests := []struct {
		password string
		valid    bool
	}{
		{"short", false}, // Less than 8 characters
		{"longpasswordwithoutnumberorspecialcharacter", false}, // More than 25 characters
		{"password1", false},      // No special character
		{"password!", false},      // No number
		{"validPassword1!", true}, // Valid password
	}

	for _, tt := range tests {
		err := validate.Var(tt.password, "custom_password")
		if tt.valid {
			assert.NoError(t, err, "Password should be valid: %s", tt.password)
		} else {
			assert.Error(t, err, "Password should be invalid: %s", tt.password)
		}
	}
}

// Definir una estructura para el test
type TestPasswordForm struct {
	Password             string `validate:"required"`
	PasswordConfirmation string `validate:"required,custom_password_confirmation"`
}

func TestCustomPasswordConfirmationValidator(t *testing.T) {
	validate := validator.New()

	// Registrar la validación custom
	validate.RegisterValidation("custom_password_confirmation", CustomPasswordConfirmationValidator)

	tests := []struct {
		password             string
		passwordConfirmation string
		valid                bool
	}{
		{"password123", "password123", true},   // Coinciden
		{"password123", "password1234", false}, // No coinciden
		{"password", "password", true},         // Coinciden
		{"password", "different", false},       // No coinciden
	}

	for _, tt := range tests {
		form := TestPasswordForm{
			Password:             tt.password,
			PasswordConfirmation: tt.passwordConfirmation,
		}

		err := validate.Struct(form)

		if tt.valid {
			assert.NoError(t, err, "Password confirmation should be valid: %s", tt.passwordConfirmation)
		} else {
			assert.Error(t, err, "Password confirmation should be invalid: %s", tt.passwordConfirmation)
		}
	}
}
