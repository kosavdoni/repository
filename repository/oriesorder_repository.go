package repositories

import (
	"database/sql"
	"fmt"

	"github.com/your-package/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) CreateOrder(order *models.Order) error {
	query := "INSERT INTO orders (user_id, total_price) VALUES ($1, $2) RETURNING id"
	err := or.db.QueryRow(query, order.UserID, order.TotalPrice).Scan(&order.ID)
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}
	return nil
}

func (or *OrderRepository) GetOrderByID(id int) (*models.Order, error) {
	query := "SELECT id, user_id, total_price, created_at, updated_at FROM orders WHERE id = $1"
	row := or.db.QueryRow(query, id)
	order := &models.Order{}
	err := row.Scan(&order.ID, &order.UserID, &order.TotalPrice, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	return order, nil
}

// Other methods for working with orders
