package domain

import (
	"errors"
	"time"
)

var ErrInsufficientStock = errors.New("insufficient stock")

type Order struct {
	ID              int       `json:"id"`
	ProductID       string    `json:"product_id"`
	ProductName     string    `json:"product_name"`
	ProductPrice    float64   `json:"product_price"`
	CustomerName    string    `json:"customer_name"`
	CustomerAddress string    `json:"customer_address"`
	Quantity        int       `json:"quantity"`
	TotalPrice      float64   `json:"total_price"`
	CreatedAt       time.Time `json:"created_at"`
}

type OrderRepository interface {
	Create(order *Order) error
	UpdateProductQuantity(productID string, quantity int) error
	CreateOrderWithTransaction(order *Order) error
	GetAllOrders() ([]Order, error)
}
