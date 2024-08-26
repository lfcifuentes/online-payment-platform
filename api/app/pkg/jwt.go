package pkg

import (
	"strconv"

	"github.com/spf13/viper"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

// PasswordHasher provides methods for hashing and comparing passwords
type ApiJWT struct {
}

// UserClaim represents the user claim for the JWT token
type UserClaim struct {
	Id int `json:"id"`
}

// CustomClaims represents the custom claims for the JWT token
type CustomClaims struct {
	Provider string    `json:"p"`
	User     UserClaim `json:"u"`
	jwt.RegisteredClaims
}

// NewApiJWT creates a new instance of ApiJWT
func NewApiJWT() *ApiJWT {
	return &ApiJWT{}
}

// GenerateJWT generates a JWT token
func (a *ApiJWT) GenerateJWT(id int) (string, time.Time, error) {
	key := viper.GetString("APP_KEY")
	expirationTime := time.Now().Add(24 * time.Hour * 7)
	claims := CustomClaims{
		Provider: "AuthApi",
		User: UserClaim{
			Id: id,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "AuthApi",
			Subject:   "AuthApi",
			ID:        strconv.Itoa(id),
			Audience:  []string{"AuthApi"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	mySigningKey := []byte(key)
	ss, err := token.SignedString(mySigningKey)

	return ss, expirationTime, err
}

// ValidateJWT validates a JWT token
func (a *ApiJWT) ValidateJWT(token string) (*CustomClaims, string, error) {

	if token == "" {
		return nil, "", jwt.ErrSignatureInvalid
	}
	// remove the Bearer prefix
	if len(token) > 6 && token[:7] == "Bearer " {
		token = token[7:]
	}

	key := viper.GetString("APP_KEY")
	mySigningKey := []byte(key)

	t, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, "", err
	}

	claims, ok := t.Claims.(*CustomClaims)
	if !ok {
		return nil, "", err
	}

	return claims, token, nil
}
