// ports/driving/order_service.go
package ports

import "example/domain"

type OrderService interface {
	CreateOrder(userID string, amount int) (domain.Order, error)
	GetOrder(id string) (domain.Order, error)
}
