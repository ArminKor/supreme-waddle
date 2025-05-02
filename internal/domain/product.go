package domain

import "time"

// Product represents a pharmacy product
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProductRepository defines the interface for product data access
type ProductRepository interface {
	Create(product *Product) error
	GetByID(id string) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(id string) error
	GetByCategory(category string) ([]*Product, error)
}
