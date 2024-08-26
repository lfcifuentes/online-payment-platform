package models

import "time"

// PaymentMethod represents a payment method
type PaymentMethod struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Name      string    `json:"name,omitempty"`
	Brand     string    `json:"brand,omitempty"`
	Last4     int       `json:"last_4,omitempty"`
	ExpMonth  int       `json:"exp_month,omitempty"`
	ExpYear   int       `json:"exp_year,omitempty"`
	BankID    int64     `json:"bank_id,omitempty"`
	BankCode  int64     `json:"bank_code,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
