package inMemory

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/AlexCorn999/order-data-service/internal/domain"
	"github.com/AlexCorn999/order-data-service/internal/repository/postgres"
)

type InMemory struct {
	mu       *sync.Mutex
	DB       map[string]*domain.Order
	postgres *postgres.Postgres
}

func NewStorage(db *postgres.Postgres) *InMemory {
	return &InMemory{
		DB:       make(map[string]*domain.Order),
		mu:       &sync.Mutex{},
		postgres: db,
	}
}

func (s *InMemory) AddOrder(order *domain.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.DB[order.OrderID]
	if ok {
		return domain.ErrAlreadyUploaded
	}

	s.DB[order.OrderID] = order
	return nil
}

func (s *InMemory) GetOrderByID(id string) (*domain.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.DB[id]
	if !ok {
		return nil, domain.ErrIncorrectOrder
	}

	return order, nil
}

func (s *InMemory) RestoreCacheFromDB() error {
	rows, err := s.postgres.DB.Query("SELECT order_info from orders")
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("cache could not be restored: %w", err)
		}
	}
	defer rows.Close()

	for rows.Next() {
		var order domain.Order
		if err := rows.Scan(&order); err != nil {
			return err
		}
		s.AddOrder(&order)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
