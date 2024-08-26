package valueobjects

import "time"

// LoginResponse represents the response structure for the authentication module.
type LoginResponse struct {
	AccessToken string    `json:"access_token"` // Access token for authentication.
	TokenType   string    `json:"token_type"`   // Type of the token.
	ExpiresIn   time.Time `json:"expires_in"`   // Expiration time of the token.
}
