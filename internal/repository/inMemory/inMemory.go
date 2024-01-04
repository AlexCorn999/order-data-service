package in_memory

import (
	"sync"

	"github.com/AlexCorn999/order-data-service/internal/domain"
)

type InMemory struct {
	mu *sync.Mutex
	DB map[string]*domain.Order
}

// NewStorage инициализирует хранилище и применяет миграции.
func NewStorage(addr string) (*InMemory, error) {
	return &InMemory{
		DB: make(map[string]*domain.Order),
	}, nil
}

// CloseDB закрывает подключение к базе данных.
func (s *InMemory) Close() error {
	return nil
}
