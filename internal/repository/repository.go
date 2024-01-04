package repository

import "github.com/AlexCorn999/order-data-service/internal/domain"

type Storage interface {
	AddOrder(order *domain.Order) error
	GetOrderByID(id string) (*domain.Order, error)
	Close() error
}
