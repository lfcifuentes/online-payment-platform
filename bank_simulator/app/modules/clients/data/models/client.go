package models

import "time"

type Client struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	BankID    int64     `json:"bank_id"`
	CreatedAt time.Time `json:"created_at"`
}
