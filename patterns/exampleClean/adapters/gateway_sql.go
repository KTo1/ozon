package adapters

import (
	"database/sql"
	"exampleClean/domain"
	"exampleClean/usecases"
)

type SQLOrderRepository struct {
	db *sql.DB
}

// Проверка, что SQLOrderRepository реализует выходную границу
var _ usecases.OrderRepository = (*SQLOrderRepository)(nil)

func NewSQLOrderRepository(db *sql.DB) *SQLOrderRepository {
	return &SQLOrderRepository{db: db}
}

func (r *SQLOrderRepository) Save(order domain.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (id, user_id, amount) VALUES (?, ?, ?)",
		order.ID, order.UserID, order.Amount)
	return err
}

func (r *SQLOrderRepository) FindByID(id string) (domain.Order, error) {
	var o domain.Order
	err := r.db.QueryRow("SELECT id, user_id, amount FROM orders WHERE id = ?", id).
		Scan(&o.ID, &o.UserID, &o.Amount)
	return o, err
}
