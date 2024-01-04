package inMemory

import (
	"sync"

	"github.com/AlexCorn999/order-data-service/internal/domain"
)

type InMemory struct {
	mu *sync.Mutex
	DB map[string]*domain.Order
}

func NewStorage() *InMemory {
	return &InMemory{
		DB: make(map[string]*domain.Order),
	}
}

func (s *InMemory) Close() error {
	return nil
}
