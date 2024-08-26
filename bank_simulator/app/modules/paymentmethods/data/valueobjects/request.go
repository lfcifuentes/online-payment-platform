package valueobjects

// Request is the value object for the request of the list payment methods use case
type NewPaymentMethodRequest struct {
	Brand      string `json:"brand" validate:"required"`
	ExpMonth   int    `json:"exp_month" validate:"required"`
	ExpYear    int    `json:"exp_year" validate:"required"`
	BankID     int    `json:"bank_id" validate:"required"`
	CardNumber string `json:"card_number,omitempty" validate:"required"`
}
