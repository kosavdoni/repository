package repository

import (
	"database/sql"
	"fmt"

	"D:\main\models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err := ur.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
	row := ur.db.QueryRow(query, id)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// Other methods for working with users
