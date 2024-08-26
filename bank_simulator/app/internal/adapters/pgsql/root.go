package pgsql

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DBAdapter struct {
	DB *sql.DB
}

func NewDBAdapter() (*DBAdapter, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("DB_HOST_BANK"),
		viper.GetInt("DB_PORT_BANK"),
		viper.GetString("DB_USER_BANK"),
		viper.GetString("DB_PASSWORD_BANK"),
		viper.GetString("DB_NAME_BANK"),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Add your code here
	return &DBAdapter{
		DB: db,
	}, nil
}

func (db *DBAdapter) Ping() error {
	return db.DB.Ping()
}

func (db *DBAdapter) Close() error {
	if db.DB != nil {
		if err := db.DB.Close(); err != nil {
			return errors.New("error al desconectar de PsqlServer")
		}
	}
	return nil
}

func (db *DBAdapter) GetDB() *sql.DB {
	return db.DB
}

func (db *DBAdapter) GetDBAdapter() *DBAdapter {
	if db.DB == nil {
		db, _ = NewDBAdapter()
	}
	return db
}
