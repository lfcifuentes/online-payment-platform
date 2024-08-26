package repositories

import (
	"time"

	"errors"

	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/bank/app/pkg"
)

type PaymentMethodRepository struct {
	DB   *pgsql.DBAdapter
	Hash *pkg.Hash
}

func NewPaymentMethodRepository(db *pgsql.DBAdapter) *PaymentMethodRepository {
	return &PaymentMethodRepository{
		DB:   db,
		Hash: pkg.NewHash(),
	}
}

func (r *PaymentMethodRepository) GetPaymentMethodsByUserCode(userID string) ([]models.PaymentMethod, error) {
	var paymentMethods []models.PaymentMethod

	// Here we should query the database to get the payment methods for the user
	rows, err := r.DB.DB.Query("SELECT id, user_id, brand, last_4, exp_month, exp_year, bank_id, created_at FROM payment_methods WHERE user_id = $1", userID)

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
		paymentMethods = append(paymentMethods, paymentMethod)
	}

	return paymentMethods, nil
}

func (r *PaymentMethodRepository) GetPaymentMethodByID(paymentMethodID int) (*models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod

	// Here we should query the database to get the payment method by ID
	err := r.DB.DB.QueryRow("SELECT id, user_id, brand, last_4, exp_month, exp_year, bank_id, created_at FROM payment_methods WHERE id = $1", paymentMethodID).Scan(
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
		return nil, err
	}

	return &paymentMethod, nil
}

func (r *PaymentMethodRepository) CreatePaymentMethod(paymentMethod *models.PaymentMethod) error {
	cardNumber, err := r.Hash.Encrypt(paymentMethod.CardNumber)
	if err != nil {
		return err
	}
	paymentMethod.CardNumber = cardNumber
	paymentMethod.CreatedAt = time.Now()
	// Here we should insert the payment method into the database
	var id int64
	err = r.DB.DB.QueryRow(
		"INSERT INTO payment_methods (user_id, brand, last_4, exp_month, exp_year, bank_id, card_number, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		paymentMethod.UserID,
		paymentMethod.Brand,
		paymentMethod.Last4,
		paymentMethod.ExpMonth,
		paymentMethod.ExpYear,
		paymentMethod.BankID,
		paymentMethod.CardNumber,
		paymentMethod.CreatedAt,
	).Scan(&id)

	if err != nil {
		return err
	}

	paymentMethod.ID = int(id)

	return nil
}

func (r *PaymentMethodRepository) DeletePaymentMethod(userID, paymentMethodID int) error {
	// Here we should delete the payment method from the database
	_, err := r.DB.DB.Exec("DELETE FROM payment_methods WHERE id = $1 AND user_id = $2", paymentMethodID, userID)

	if err != nil {
		return err
	}

	return nil
}
