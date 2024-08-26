package repositories

import (
	"time"

	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/payments/data/models"
	"github.com/lfcifuentes/online-payment-platform/bank/app/pkg"
)

type PaymentRepository struct {
	DB   *pgsql.DBAdapter
	Hash *pkg.Hash
}

func NewPaymentRepository(db *pgsql.DBAdapter) *PaymentRepository {
	return &PaymentRepository{
		DB:   db,
		Hash: pkg.NewHash(),
	}
}

func (r *PaymentRepository) CreatePayment(paymentMethod *models.Payment) error {
	paymentMethod.CreatedAt = time.Now()
	// Here we should insert the payment method into the database
	var id int64
	err := r.DB.DB.QueryRow("INSERT INTO payments (bank_id, payment_method_id, receiver_id, amount, status, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		paymentMethod.BankID,
		paymentMethod.PaymentMethodID,
		paymentMethod.ReceiverID,
		paymentMethod.Amount,
		paymentMethod.Status,
		paymentMethod.CreatedAt,
	).Scan(&id)

	if err != nil {
		return err
	}

	paymentMethod.ID = id

	return nil
}

func (r *PaymentRepository) UpdatePaymentStatus(paymentMethod *models.Payment) error {
	paymentMethod.UpdatedAt = time.Now()
	// Here we should update the payment method status in the database
	_, err := r.DB.DB.Exec("UPDATE payments SET status=$1, updated_at=$2 WHERE id=$3",
		paymentMethod.Status,
		paymentMethod.UpdatedAt,
		paymentMethod.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *PaymentRepository) GetPaymentByID(id int64) (*models.Payment, error) {
	var payment models.Payment
	err := r.DB.DB.QueryRow("SELECT id, bank_id, payment_method_id, receiver_id, amount, status, created_at, updated_at FROM payments WHERE id = $1", id).Scan(
		&payment.ID,
		&payment.BankID,
		&payment.PaymentMethodID,
		&payment.ReceiverID,
		&payment.Amount,
		&payment.Status,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
