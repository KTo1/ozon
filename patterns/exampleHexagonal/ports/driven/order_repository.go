// ports/driven/order_repository.go
package ports

import "example/domain"

type OrderRepository interface {
	Save(order domain.Order) error
	FindByID(id string) (domain.Order, error)
}
