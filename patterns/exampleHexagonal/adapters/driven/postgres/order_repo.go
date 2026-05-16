// adapters/driven/postgres/order_repo.go
package postgres

import (
	"database/sql"
	"example/domain"
	driven "example/ports/driven"
)

type OrderRepo struct {
	db *sql.DB
}

// Проверка, что структура реализует порт
var _ driven.OrderRepository = (*OrderRepo)(nil)

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Save(order domain.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (id, user_id, amount) VALUES ($1,$2,$3)",
		order.ID, order.UserID, order.Amount)
	return err
}

func (r *OrderRepo) FindByID(id string) (domain.Order, error) {
	var o domain.Order
	err := r.db.QueryRow("SELECT id, user_id, amount FROM orders WHERE id=$1", id).
		Scan(&o.ID, &o.UserID, &o.Amount)
	return o, err
}
