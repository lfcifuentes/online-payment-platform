package pkg

import (
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher provides methods for hashing and comparing passwords.
type PasswordHasher struct{}

// NewPasswordHasher creates a new instance of PasswordHasher.
func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (h *PasswordHasher) Make(password string) (string, error) {
	key := viper.GetString("APP_KEY")
	hash, err := bcrypt.GenerateFromPassword([]byte(password+key), bcrypt.MinCost)
	return string(hash), err
}

func (h *PasswordHasher) Check(password string, hash string) bool {
	key := viper.GetString("APP_KEY")
	byteHash := []byte(hash)
	bytePasswordHash := []byte(password + key)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePasswordHash)
	return err == nil
}
