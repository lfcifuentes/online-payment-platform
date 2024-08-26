package valueobjects

type CreateClientRequest struct {
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	BankID int64  `json:"bank_id" validate:"required"`
}
