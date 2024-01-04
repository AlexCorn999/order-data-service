package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/AlexCorn999/order-data-service/internal/domain"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Postgres struct {
	DB *sql.DB
}

func NewStorage(addr string) (*Postgres, error) {
	db, err := goose.OpenDBWithDriver("pgx", addr)
	if err != nil {
		return nil, fmt.Errorf("goose: failed to open DB: %w", err)
	}

	err = goose.Up(db, "./migrations")
	if err != nil {
		return nil, fmt.Errorf("goose: failed to migrate: %w", err)
	}

	return &Postgres{
		DB: db,
	}, nil
}

func (p *Postgres) AddOrder(order *domain.Order) error {

	data, err := json.Marshal(*order)
	if err != nil {
		return fmt.Errorf("postgres: can't marshal json object: %w", err)
	}

	result, err := p.DB.Exec("INSERT INTO orders (order_uid, order_info) VALUES ($1, $2) on conflict (order_uid) do nothing", order.OrderID, data)
	if err != nil {
		return fmt.Errorf("postgres: couldn't add the order details: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("postgreSQL: addOrder %w", err)
	}

	if rowsAffected == 0 {
		return domain.ErrAlreadyUploaded
	}
	return nil
}

func (p *Postgres) GetOrderByID(id string) (*domain.Order, error) {
	var order domain.Order
	var data []byte
	if err := p.DB.QueryRow("SELECT order_info from orders WHERE order_uid=$1", id).Scan(&data); err != nil {
		return nil, fmt.Errorf("postgres: failed to retrieve the value from the database: %w", err)
	}

	if err := json.Unmarshal(data, &order); err != nil {
		return nil, fmt.Errorf("postgres: can't unmarshal data to json object: %w", err)
	}
	return &order, nil
}

func (s *Postgres) Close() error {
	return s.DB.Close()
}
