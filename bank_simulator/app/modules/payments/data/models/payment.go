package models

import "time"

const (
	PENDING  = "pending"
	APPROVED = "approved"
	REJECTED = "rejected"
	REFUND   = "refund"
)

type Payment struct {
	ID              int64     `json:"id"`
	BankID          int64     `json:"bank_id" validate:"required"`
	PaymentMethodID int64     `json:"payment_method_id" validate:"required"`
	ReceiverID      int64     `json:"receiver_id" validate:"required"`
	Amount          float64   `json:"amount" validate:"required"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
