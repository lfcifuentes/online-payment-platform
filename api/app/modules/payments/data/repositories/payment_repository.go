package repositories

import (
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/paymentmethods/data/models"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/payments/data/valueobjects"
)

type PaymentRepository struct {
	DBAdapter *pgsql.DBAdapter
}

func NewPaymentRepository(db *pgsql.DBAdapter) *PaymentRepository {
	// Add your code here
	return &PaymentRepository{
		DBAdapter: db,
	}
}

func (r *PaymentRepository) GetTransaction(userID, transactionID int64) (*models.Payment, error) {
	var payment models.Payment
	err := r.DBAdapter.DB.QueryRow(
		"SELECT id, bank_id, payment_method_id, receiver_id, amount, status, created_at, updated_at FROM payments WHERE id = $1 AND user_id = $2",
		transactionID,
		userID,
	).Scan(
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
func (r *PaymentRepository) GetPaymentByBankPayID(bankPaymentId string) (*models.Payment, error) {
	var payment models.Payment
	err := r.DBAdapter.DB.QueryRow(
		"SELECT id, bank_id, payment_method_id, receiver_id, amount, status, created_at, updated_at FROM payments WHERE bank_reference = $1",
		bankPaymentId,
	).Scan(
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

func (r *PaymentRepository) CreatePayment(
	userID string,
	request valueobjects.PayParams,
	paymentBankId int64,
	paymentBankStatus string,
) error {

	_, err := r.DBAdapter.DB.Exec(
		"INSERT INTO payments (user_id, payment_method_id, receiver_id, amount, bank_id, bank_reference, status) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		userID,
		request.PaymentMethodID,
		request.MerchantPaymentMethodID,
		request.Amount,
		"1",
		paymentBankId,
		paymentBankStatus,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepository) GetTransactions(userID int64) ([]models.Payment, error) {
	transactions := make([]models.Payment, 0)
	query := `
		SELECT 
			p.id, 
			p.payment_method_id, 
			p.receiver_id,
			p.amount,
			p.status,
			p.created_at,
			p.updated_at,
			pm.name as payment_method_name,
			pm.user_id as payment_method_user_id,
			pm2.name as receiver_name,
			pm2.user_id as receiver_user_id
		FROM payments AS p 
			INNER JOIN payment_methods AS pm ON p.payment_method_id = pm.id
			INNER JOIN payment_methods AS pm2 ON p.receiver_id = pm2.id
		WHERE p.user_id = $1 
		ORDER BY p.created_at DESC
	`
	rows, err := r.DBAdapter.DB.Query(query,
		userID,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transaction models.Payment
		var paymentMethod models.PaymentMethod
		var receiver models.PaymentMethod
		err = rows.Scan(
			&transaction.ID,
			&transaction.PaymentMethodID,
			&transaction.ReceiverID,
			&transaction.Amount,
			&transaction.Status,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&paymentMethod.Name,
			&paymentMethod.UserID,
			&receiver.Name,
			&receiver.UserID,
		)
		if err != nil {
			return nil, err
		}
		receiver.ID = transaction.ReceiverID
		transaction.Receiver = &receiver

		paymentMethod.ID = transaction.PaymentMethodID
		transaction.PaymentMethod = &paymentMethod

		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *PaymentRepository) GetTransactionsReceives(userID int64) ([]models.Payment, error) {
	transactions := make([]models.Payment, 0)
	query := `
		SELECT 
			p.id, 
			p.payment_method_id, 
			p.receiver_id,
			p.amount,
			p.status,
			p.created_at,
			p.updated_at,
			pm.name as payment_method_name,
			pm.user_id as payment_method_user_id,
			pm2.name as receiver_name,
			pm2.user_id as receiver_user_id
		FROM payments AS p 
			INNER JOIN payment_methods AS pm ON p.payment_method_id = pm.id
			INNER JOIN payment_methods AS pm2 ON p.receiver_id = pm2.id
		WHERE pm2.user_id = $1
		ORDER BY p.created_at DESC
	`
	rows, err := r.DBAdapter.DB.Query(query,
		userID,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var transaction models.Payment
		var paymentMethod models.PaymentMethod
		var receiver models.PaymentMethod
		err = rows.Scan(
			&transaction.ID,
			&transaction.PaymentMethodID,
			&transaction.ReceiverID,
			&transaction.Amount,
			&transaction.Status,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&paymentMethod.Name,
			&paymentMethod.UserID,
			&receiver.Name,
			&receiver.UserID,
		)
		if err != nil {
			return nil, err
		}
		receiver.ID = transaction.ReceiverID
		transaction.Receiver = &receiver

		paymentMethod.ID = transaction.PaymentMethodID
		transaction.PaymentMethod = &paymentMethod

		transactions = append(transactions, transaction)
	}
	return transactions, nil
}

func (r *PaymentRepository) GetTransactionReceive(userID, transactionID int64) (*models.Payment, error) {
	var payment models.Payment
	var paymentMethod models.PaymentMethod
	var receiver models.PaymentMethod
	query := `
		SELECT 
			p.id, 
			p.payment_method_id, 
			p.receiver_id,
			p.amount,
			p.status,
			p.created_at,
			p.updated_at,
			p.bank_reference,
			pm.name as payment_method_name,
			pm.user_id as payment_method_user_id,
			pm2.name as receiver_name,
			pm2.user_id as receiver_user_id
		FROM payments AS p 
			INNER JOIN payment_methods AS pm ON p.payment_method_id = pm.id
			INNER JOIN payment_methods AS pm2 ON p.receiver_id = pm2.id
		WHERE p.id = $1
			AND pm2.user_id = $2
		ORDER BY p.created_at DESC
	`
	err := r.DBAdapter.DB.QueryRow(
		query,
		transactionID,
		userID,
	).Scan(
		&payment.ID,
		&payment.PaymentMethodID,
		&payment.ReceiverID,
		&payment.Amount,
		&payment.Status,
		&payment.CreatedAt,
		&payment.UpdatedAt,
		&payment.BankReference,
		&paymentMethod.Name,
		&paymentMethod.UserID,
		&receiver.Name,
		&receiver.UserID,
	)
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) RefundTransaction(transaction *models.Payment) error {
	_, err := r.DBAdapter.DB.Exec(
		"UPDATE payments SET status = $2 WHERE id = $1",
		transaction.ID,
		models.PaymentStatusRefunded,
	)
	if err != nil {
		return err
	}
	transaction.Status = models.PaymentStatusRefunded
	transaction.BankReference = 0
	return nil
}
