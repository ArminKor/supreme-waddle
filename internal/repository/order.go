package repository

import (
	"database/sql"
	"pharmacy-shop/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *domain.Order) error {
	query := `
        INSERT INTO orders (product_id, customer_name, customer_address, quantity, total_price)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at`

	err := r.db.QueryRow(
		query,
		order.ProductID,
		order.CustomerName,
		order.CustomerAddress,
		order.Quantity,
		order.TotalPrice,
	).Scan(&order.ID, &order.CreatedAt)

	return err
}

func (r *OrderRepository) UpdateProductQuantity(productID int, quantity int) error {
	query := `
        UPDATE products 
        SET quantity = quantity - $1
        WHERE id = $2 AND quantity >= $1`

	result, err := r.db.Exec(query, quantity, productID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *OrderRepository) CreateOrderWithTransaction(order *domain.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Update product quantity
	updateQuery := `
        UPDATE products 
        SET quantity = quantity - $1
        WHERE id = $2 AND quantity >= $1`

	result, err := tx.Exec(updateQuery, order.Quantity, order.ProductID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	// Create the order
	orderQuery := `
        INSERT INTO orders (product_id, customer_name, customer_address, quantity, total_price)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at`

	err = tx.QueryRow(
		orderQuery,
		order.ProductID,
		order.CustomerName,
		order.CustomerAddress,
		order.Quantity,
		order.TotalPrice,
	).Scan(&order.ID, &order.CreatedAt)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *OrderRepository) GetAllOrders() ([]domain.Order, error) {
	query := `
        SELECT o.id, o.product_id, o.customer_name, o.customer_address, 
               o.quantity, o.total_price, o.created_at,
               p.name as product_name, p.price as product_price
        FROM orders o
        JOIN products p ON o.product_id = p.id
        ORDER BY o.created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var order domain.Order
		var productName string
		var productPrice float64
		err := rows.Scan(
			&order.ID,
			&order.ProductID,
			&order.CustomerName,
			&order.CustomerAddress,
			&order.Quantity,
			&order.TotalPrice,
			&order.CreatedAt,
			&productName,
			&productPrice,
		)
		if err != nil {
			return nil, err
		}
		order.ProductName = productName
		order.ProductPrice = productPrice
		orders = append(orders, order)
	}

	return orders, nil
}
