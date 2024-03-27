package repositories

import (
	"database/sql"
	"fmt"

	"github.com/your-package/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) CreateProduct(product *models.Product) error {
	query := "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id"
	err := pr.db.QueryRow(query, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

func (pr *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	query := "SELECT id, name, price, created_at, updated_at FROM products WHERE id = $1"
	row := pr.db.QueryRow(query, id)
	product := &models.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	return product, nil
}

// Other methods for working with products
