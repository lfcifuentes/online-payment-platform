package valueobjects

// PayParams Pagar un producto
type PayParams struct {
	MerchantID              int64   `json:"merchant_id" validate:"required" documentation:"ID del comercio"`
	PaymentMethodID         int64   `json:"payment_method_id" validate:"required" documentation:"ID del método de pago"`
	MerchantPaymentMethodID int64   `json:"merchant_payment_method_id" validate:"required" documentation:"ID del método de pago del comercio"`
	Amount                  float64 `json:"amount" validate:"required" documentation:"Monto a pagar"`
}
