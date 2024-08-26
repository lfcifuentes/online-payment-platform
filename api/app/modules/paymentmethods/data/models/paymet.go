package models

import "time"

const (
	PaymentStatusRefunded = "refunded"
)

type Payment struct {
	ID              int64     `json:"id"`
	BankID          int64     `json:"bank_id"`
	BankReference   int64     `json:"bank_reference,omitempty"`
	PaymentMethodID int64     `json:"payment_method_id"`
	ReceiverID      int64     `json:"receiver_id"`
	Amount          float64   `json:"amount"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`

	Receiver      *PaymentMethod `json:"receiver"`
	PaymentMethod *PaymentMethod `json:"payment_method"`
}
