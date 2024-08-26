package valueobjects

// NewPaymentMethodRequest represents a new payment method request
type NewPaymentMethodRequest struct {
	Name       string `json:"name" validate:"required"`
	Brand      string `json:"brand" validate:"required"`
	CardNumber string `json:"card_number" validate:"required"`
	ExpMonth   int    `json:"exp_month" validate:"required"`
	ExpYear    int    `json:"exp_year" validate:"required"`
}
