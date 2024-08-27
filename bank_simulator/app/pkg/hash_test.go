package pkg

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestHash_Encrypt(t *testing.T) {
	viper.SetDefault("HASH_KEY", "12345678901234567890123456789012")
	hash := NewHash()
	plaintext := "Hello, World!"

	ciphertext, err := hash.Encrypt(plaintext)
	assert.NoError(t, err)

	decryptedPlaintext, err := hash.Decrypt(ciphertext)
	assert.NoError(t, err)

	assert.Equal(t, plaintext, decryptedPlaintext)
}
