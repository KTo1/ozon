// adapters/driven/mysql/order_repo.go
package mysql

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
	query := "INSERT INTO orders (id, user_id, amount) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, order.ID, order.UserID, order.Amount)
	return err
}

func (r *OrderRepo) FindByID(id string) (domain.Order, error) {
	var o domain.Order
	query := "SELECT id, user_id, amount FROM orders WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&o.ID, &o.UserID, &o.Amount)
	return o, err
}
