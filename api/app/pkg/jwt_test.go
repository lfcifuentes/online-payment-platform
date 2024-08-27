package pkg

import (
	"testing"
)

func TestApiJWT_ValidateJWT(t *testing.T) {
	apiJWT := NewApiJWT()

	// Generate a valid JWT token for testing
	validToken, _, _ := apiJWT.GenerateJWT(1)

	// Test case: Valid token
	claims, token, err := apiJWT.ValidateJWT(validToken)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if claims == nil {
		t.Error("Expected claims to be non-nil")
	}

	if token != validToken {
		t.Errorf("Expected token to be %s, got %s", validToken, token)
	}

	// Test case: Invalid token
	invalidToken := "invalid_token"

	claims, token, err = apiJWT.ValidateJWT(invalidToken)

	if err == nil {
		t.Error("Expected error, but got nil")
	}

	if claims != nil {
		t.Error("Expected claims to be nil")
	}

	if token != "" {
		t.Error("Expected token to be empty")
	}
}
