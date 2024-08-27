package pgsql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewDBAdapter(t *testing.T) {
	// Simula las configuraciones de Viper
	viper.Set("DB_HOST", "localhost")
	viper.Set("DB_PORT", 5432)
	viper.Set("DB_USER", "testuser")
	viper.Set("DB_PASSWORD", "password")
	viper.Set("DB_NAME", "testdb")

	// Mock de la base de datos
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	adapter, err := &DBAdapter{
		DB: db,
	}, nil

	//
	assert.NoError(t, err)
	assert.NotNil(t, adapter)

	// Verifica si se ha ejecutado Ping con éxito
	mock.ExpectPing()
	err = adapter.Ping()
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBAdapter_Close(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	adapter := &DBAdapter{
		DB: db,
	}

	// Simula un cierre exitoso
	mock.ExpectClose()
	err = adapter.Close()
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDBAdapter_GetDB(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)

	adapter := &DBAdapter{
		DB: db,
	}

	assert.Equal(t, db, adapter.GetDB())
}
