package entities

// Bank Entidad Banco
type Bank struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Webhook   string `json:"webhook"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
