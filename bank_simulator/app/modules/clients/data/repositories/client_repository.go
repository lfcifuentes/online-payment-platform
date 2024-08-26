package repositories

import (
	"time"

	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/bank/app/modules/clients/data/models"
)

type ClientRepository struct {
	DBAdapter *pgsql.DBAdapter
}

func NewClientRepository(dbAdapter *pgsql.DBAdapter) *ClientRepository {
	return &ClientRepository{DBAdapter: dbAdapter}
}

func (cr *ClientRepository) GetClientByID(id int64) (*models.Client, error) {
	// get client by id
	client := &models.Client{}
	err := cr.DBAdapter.DB.QueryRow(
		"SELECT id, name, email, code, bank_id, created_at FROM clients WHERE id = $1",
		id,
	).Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.Code,
		&client.BankID,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (cr *ClientRepository) GetClientByCode(code string) (*models.Client, error) {
	// get client by code
	client := &models.Client{}
	err := cr.DBAdapter.DB.QueryRow(
		"SELECT id, name, email, code, bank_id, created_at FROM clients WHERE code = $1",
		code,
	).Scan(
		&client.ID,
		&client.Name,
		&client.Email,
		&client.Code,
		&client.BankID,
		&client.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (cr *ClientRepository) CreateClient(client *models.Client) error {
	client.CreatedAt = time.Now()
	_, err := cr.DBAdapter.DB.Exec(
		"INSERT INTO clients (name, email, code, bank_id, created_at) VALUES ($1, $2, $3, $4, $5)",
		client.Name,
		client.Email,
		client.Code,
		client.BankID,
		client.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
