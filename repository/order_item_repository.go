package repositories

import (
	"database/sql"
	"fmt"

	"github.com/your-package/models"
)

type OrderItemRepository struct {
	db *sql.DB
}

func NewOrderItemRepository(db *sql.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

func (oir *OrderItemRepository) CreateOrderItem(orderItem *models.OrderItem) error {
	query := "INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)"
	_, err := oir.db.Exec(query, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity)
	if err != nil {
		return fmt.Errorf("failed to create order item: %w", err)
	}
	return nil
}

// Other methods for working with order items
