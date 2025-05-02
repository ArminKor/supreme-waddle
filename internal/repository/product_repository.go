package repository

import (
	"database/sql"
	"pharmacy-shop/internal/domain"
	"time"
)

// ProductRepository implements the domain.ProductRepository interface
type ProductRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &ProductRepository{db: db}
}

// Create implements domain.ProductRepository
func (r *ProductRepository) Create(product *domain.Product) error {
	query := `
		INSERT INTO products (name, description, price, quantity, category)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.Category,
	).Scan(&product.ID, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// GetByID implements domain.ProductRepository
func (r *ProductRepository) GetByID(id string) (*domain.Product, error) {
	query := `
		SELECT id, name, description, price, quantity, category, created_at, updated_at
		FROM products
		WHERE id = $1`

	product := &domain.Product{}
	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.Category,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetAll implements domain.ProductRepository
func (r *ProductRepository) GetAll() ([]*domain.Product, error) {
	query := `
		SELECT id, name, description, price, quantity, category, created_at, updated_at
		FROM products
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		product := &domain.Product{}
		errScan := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
			&product.Category,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if errScan != nil {
			return nil, errScan
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Update implements domain.ProductRepository
func (r *ProductRepository) Update(product *domain.Product) error {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, quantity = $4, category = $5, updated_at = $6
		WHERE id = $7
		RETURNING updated_at`

	product.UpdatedAt = time.Now()
	err := r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.Category,
		product.UpdatedAt,
		product.ID,
	).Scan(&product.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

// Delete implements domain.ProductRepository
func (r *ProductRepository) Delete(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ProductRepository) GetByCategory(category string) ([]*domain.Product, error) {
	query := `
		SELECT id, name, description, price, quantity, category, created_at, updated_at
		FROM products
		WHERE category = $1
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		product := &domain.Product{}
		errScan := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
			&product.Category,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if errScan != nil {
			return nil, errScan
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
