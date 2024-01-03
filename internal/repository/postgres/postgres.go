package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Postgres struct {
	DB *sql.DB
}

// NewStorage инициализирует хранилище и применяет миграции.
func NewStorage(addr string) (*Postgres, error) {
	db, err := goose.OpenDBWithDriver("pgx", addr)
	if err != nil {
		return nil, fmt.Errorf("goose: failed to open DB: %v", err)
	}

	err = goose.Up(db, "./migrations")
	if err != nil {
		return nil, fmt.Errorf("goose: failed to migrate: %v", err)
	}

	return &Postgres{
		DB: db,
	}, nil
}

// CloseDB закрывает подключение к базе данных.
func (s *Postgres) Close() error {
	return s.DB.Close()
}
