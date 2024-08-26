package repositories

import (
	"time"

	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/merchant/data/models"
)

type MerchantRepository struct {
	DBAdapter *pgsql.DBAdapter
}

func NewMerchantRepository(db *pgsql.DBAdapter) *MerchantRepository {
	// Add your code here
	return &MerchantRepository{
		DBAdapter: db,
	}
}

func (r *MerchantRepository) GetMerchantByUserID(userID string) (*models.Merchant, error) {
	var merchant models.Merchant
	row := r.DBAdapter.DB.QueryRow(
		"SELECT id, user_id, status, created_at FROM merchants WHERE user_id = $1",
		userID,
	)
	err := row.Scan(
		&merchant.ID,
		&merchant.UserID,
		&merchant.Status,
		&merchant.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) GetMerchantById(merchantID int64) (*models.Merchant, error) {
	var merchant models.Merchant
	row := r.DBAdapter.DB.QueryRow(
		"SELECT id, user_id, status, created_at FROM merchants WHERE id = $1",
		merchantID,
	)
	err := row.Scan(
		&merchant.ID,
		&merchant.UserID,
		&merchant.Status,
		&merchant.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) CreateMerchant(userID string) error {
	_, err := r.DBAdapter.DB.Exec(
		"INSERT INTO merchants (user_id, status, created_at) VALUES ($1, $2, $3)",
		userID,
		true,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
