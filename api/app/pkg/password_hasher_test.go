package pkg

import (
	"testing"

	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHasher_Make(t *testing.T) {
	viper.Set("APP_KEY", "test_key")
	password := "password123"
	anotherHash, _ := bcrypt.GenerateFromPassword([]byte(password+"test_key"), bcrypt.MinCost)

	hasher := NewPasswordHasher()
	hash, err := hasher.Make(password)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if hash == string(anotherHash) {
		t.Errorf("Expected diferent hash: %s, %s", string(anotherHash), hash)
	}
}

func TestPasswordHasher_Check(t *testing.T) {
	viper.Set("APP_KEY", "test_key")
	password := "password123"

	hasher := NewPasswordHasher()

	hash, err := hasher.Make(password)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	match := hasher.Check(password, string(hash))

	if !match {
		t.Errorf("Expected password and hash to match, but they didn't")
	}
}
