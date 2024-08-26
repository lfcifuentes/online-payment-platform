package valueobjects

// PayParams is a struct that represents the request for the PayUseCase
type PayParams struct {
	BankID          int64   `json:"bank_id" validate:"required"`
	PaymentMethodID int64   `json:"payment_method_id" validate:"required"`
	ReceiverID      int64   `json:"receiver_id" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
}
