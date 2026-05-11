// core/order_service_impl.go
package core

import (
	"errors"
	"example/domain"
	driven "example/ports/driven"
	driving "example/ports/driving"
	"github.com/google/uuid"
)

type orderServiceImpl struct {
	repo driven.OrderRepository
}

// Проверка на этапе компиляции, что реализация соответствует порту
var _ driving.OrderService = (*orderServiceImpl)(nil)

func NewOrderService(repo driven.OrderRepository) driving.OrderService {
	return &orderServiceImpl{repo: repo}
}

func (s *orderServiceImpl) CreateOrder(userID string, amount int) (domain.Order, error) {
	if amount <= 0 {
		return domain.Order{}, errors.New("amount must be positive")
	}

	order := domain.Order{
		ID:     uuid.New().String(),
		UserID: userID,
		Amount: amount,
	}

	if err := s.repo.Save(order); err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (s *orderServiceImpl) GetOrder(id string) (domain.Order, error) {
	return s.repo.FindByID(id)
}
