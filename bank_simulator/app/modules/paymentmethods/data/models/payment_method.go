package models

import "time"

type PaymentMethod struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Last4      string    `json:"last4"`
	Brand      string    `json:"brand"`
	ExpMonth   int       `json:"exp_month"`
	ExpYear    int       `json:"exp_year"`
	BankID     int       `json:"bank_id"`
	CardNumber string    `json:"card_number,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}
