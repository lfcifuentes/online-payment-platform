package models

import (
	"time"
)

type Merchant struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	UserID    int64     `json:"user_id"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
