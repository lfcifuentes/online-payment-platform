package repositories

import (
	"errors"
	"time"

	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/auth/data/models"
)

type AuthRepository struct {
	DBAdapter *pgsql.DBAdapter
}

func NewAuthRepository(db *pgsql.DBAdapter) *AuthRepository {
	// Add your code here
	return &AuthRepository{
		DBAdapter: db,
	}
}

// FindByEmail retrieves a user by email.
func (r *AuthRepository) FindByEmail(email string) (*models.User, error) {

	row := r.DBAdapter.DB.QueryRow("SELECT id, email, password, status, client_id FROM users WHERE email = $1", email)
	var user models.User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Status,
		&user.ClientID,
	)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// SaveToken saves a token in the database.
func (r *AuthRepository) SaveToken(id int, token string, expirationTime time.Time) error {
	_, err := r.DBAdapter.DB.Exec("INSERT INTO tokens (user_id, token, expires_at, status) VALUES ($1, $2, $3, $4)", id, token, expirationTime, true)
	if err != nil {
		return errors.New("error saving token")
	}

	return nil
}

// ValidateToken checks if a token is valid.
func (r *AuthRepository) ValidateToken(token string) (bool, error) {
	rows := r.DBAdapter.DB.QueryRow("SELECT status FROM tokens WHERE token = $1", token)
	var status bool

	err := rows.Scan(&status)
	if err != nil {
		return false, errors.New("token not found")
	}

	return status, nil
}

// InvalidateToken invalidates a token.
func (r *AuthRepository) InvalidateToken(token string) error {
	_, err := r.DBAdapter.DB.Exec("UPDATE tokens SET status = false WHERE token = $1", token)
	if err != nil {
		return errors.New("error invalidating token")
	}

	return nil
}

// Create creates a new user.
func (r *AuthRepository) Create(username, email, password string) error {
	_, err := r.DBAdapter.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", username, email, password)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			return errors.New("email already exists")
		}
		return errors.New("error creating user")
	}

	return nil
}

// UpdateClientID updates the client_id of a user.
func (r *AuthRepository) UpdateClientID(email string, clientID int64) error {
	_, err := r.DBAdapter.DB.Exec("UPDATE users SET client_id = $1 WHERE email = $2", clientID, email)
	if err != nil {
		return errors.New("error updating client_id")
	}

	return nil
}
