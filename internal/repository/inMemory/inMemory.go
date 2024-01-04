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
		mu: &sync.Mutex{},
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

func (s *InMemory) Close() error {
	return nil
}
