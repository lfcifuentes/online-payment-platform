package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/spf13/viper"
)

type Hash struct {
	Key []byte
}

func NewHash() *Hash {
	return &Hash{
		Key: []byte(viper.GetString("HASH_KEY")),
	}
}

// Función para cifrar los datos
func (h *Hash) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(h.Key)
	if err != nil {
		return "", err
	}

	// GCM es un modo de operación de AES que proporciona autenticación y confidencialidad
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Cifrar los datos
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)

	// Devolver los datos cifrados como una cadena codificada en base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Función para descifrar los datos
func (h *Hash) Decrypt(ciphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(h.Key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := data[:nonceSize], string(data[nonceSize:])

	// Descifrar los datos
	plaintext, err := aesGCM.Open(nil, nonce, []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
