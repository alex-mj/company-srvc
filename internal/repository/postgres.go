package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sqlx.DB
}

type PostgresConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresStorage(cfg PostgresConfig) (*PostgresDB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return &PostgresDB{}, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresDB{db: db}, nil
}
