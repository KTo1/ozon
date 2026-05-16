package usecases

import (
	"errors"
	"exampleClean/domain"
	"github.com/google/uuid"
)

type OrderInteractor struct {
	repo OrderRepository // зависит от выходной границы
}

// Проверка, что OrderInteractor реализует входную границу
var _ OrderInputBoundary = (*OrderInteractor)(nil)

func NewOrderInteractor(repo OrderRepository) *OrderInteractor {
	return &OrderInteractor{repo: repo}
}

func (uc *OrderInteractor) CreateOrder(userID string, amount int) (domain.Order, error) {
	if amount <= 0 {
		return domain.Order{}, errors.New("amount must be positive")
	}

	order := domain.Order{
		ID:     uuid.New().String(),
		UserID: userID,
		Amount: amount,
	}

	if err := uc.repo.Save(order); err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (uc *OrderInteractor) GetOrder(id string) (domain.Order, error) {
	return uc.repo.FindByID(id)
}
