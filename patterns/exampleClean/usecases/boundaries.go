package usecases

import "exampleClean/domain"

// Input Boundary – то, что вызывает внешний мир
type OrderInputBoundary interface {
	CreateOrder(userID string, amount int) (domain.Order, error)
	GetOrder(id string) (domain.Order, error)
}

// Output Boundary (Gateway Interface) – то, что нужно интерактору от инфраструктуры
type OrderRepository interface {
	Save(order domain.Order) error
	FindByID(id string) (domain.Order, error)
}
