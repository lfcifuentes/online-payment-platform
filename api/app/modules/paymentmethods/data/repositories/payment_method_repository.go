package repositories

import (
	"errors"

	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
)

type PaymentMethodRepository struct {
	DB      *pgsql.DBAdapter
	bankApi pkg.BankApi
}

func NewPaymentMethodRepository(db *pgsql.DBAdapter) *PaymentMethodRepository {
	return &PaymentMethodRepository{
		DB:      db,
		bankApi: *pkg.NewBankApi(),
	}
}

func (r *PaymentMethodRepository) GetPaymentMethodsByUserCode(userID string) ([]*models.PaymentMethod, error) {
	var paymentMethods []*models.PaymentMethod

	// Here we should query the database to get the payment methods for the user
	rows, err := r.DB.DB.Query("SELECT id, user_id, brand, last_four, exp_month, exp_year, bank_id, created_at FROM payment_methods WHERE user_id = $1", userID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var paymentMethod models.PaymentMethod
		err := rows.Scan(
			&paymentMethod.ID,
			&paymentMethod.UserID,
			&paymentMethod.Brand,
			&paymentMethod.Last4,
			&paymentMethod.ExpMonth,
			&paymentMethod.ExpYear,
			&paymentMethod.BankID,
			&paymentMethod.CreatedAt,
		)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return nil, errors.New("no payment methods found")
			}
			return nil, err
		}
		paymentMethods = append(paymentMethods, &paymentMethod)
	}

	return paymentMethods, nil
}

func (r *PaymentMethodRepository) CreatePaymentMethod(paymentMethod models.PaymentMethod) error {
	_, err := r.DB.DB.Exec(
		"INSERT INTO payment_methods (name,user_id, brand, last_four, exp_month, exp_year, bank_id, bank_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		paymentMethod.Name,
		paymentMethod.UserID,
		paymentMethod.Brand,
		paymentMethod.Last4,
		paymentMethod.ExpMonth,
		paymentMethod.ExpYear,
		paymentMethod.BankID,
		paymentMethod.BankCode,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *PaymentMethodRepository) GetPaymentMethodByID(userID, paymentMethodID string) (*models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod

	err := r.DB.DB.QueryRow("SELECT id, user_id, brand, last_four, exp_month, exp_year, bank_id, bank_code, created_at FROM payment_methods WHERE user_id = $1 AND id = $2", userID, paymentMethodID).Scan(
		&paymentMethod.ID,
		&paymentMethod.UserID,
		&paymentMethod.Brand,
		&paymentMethod.Last4,
		&paymentMethod.ExpMonth,
		&paymentMethod.ExpYear,
		&paymentMethod.BankID,
		&paymentMethod.BankCode,
		&paymentMethod.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &paymentMethod, nil
}

func (r *PaymentMethodRepository) DeletePaymentMethod(userID, paymentMethodID string) error {
	_, err := r.DB.DB.Exec("DELETE FROM payment_methods WHERE user_id = $1 AND id = $2", userID, paymentMethodID)
	if err != nil {
		return err
	}
	return nil
}
