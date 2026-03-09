package repository

import (
	"database/sql"
	"order-system/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}
func (r *OrderRepository) Create(order model.Order) error {
	query := `
INSERT INTO orders (user_id,product,price)
VALUES ($1,$2,$3)`
	_, err := r.db.Exec(query, order.UserID, order.Product, order.Price)
	return err
}
