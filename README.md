# Pharmacy Shop

A web application for managing a pharmacy shop's inventory and orders.

## Features

- Product Management
  - Add, edit, and delete products
  - Track product inventory
  - Categorize products
  - View product details

- Order Management
  - Create new orders
  - View order history
  - Track order status
  - Automatic inventory updates

- Customer Interface
  - Browse available products
  - Place orders
  - View order status

## Tech Stack

- Backend: Go
- Database: PostgreSQL
- Frontend: HTML, CSS, JavaScript

## Setup

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Set up PostgreSQL using Docker:
   ```bash
   docker run --name pharmacy-db -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=pharmacy_shop -p 5432:5432 -d postgres
   ```

3. Run migrations:
   ```bash
   migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/pharmacy_shop?sslmode=disable" up
   ```

4. Start the server:
   ```bash
   go run cmd/main.go
   ```

5. Access the application:
   - Admin interface: http://localhost:8080
   - Customer interface: http://localhost:8080/customer.html

## Project Structure

```
pharmacy-shop/
├── cmd/
│   └── main.go
├── internal/
│   ├── domain/
│   │   ├── product.go
│   │   └── order.go
│   ├── repository/
│   │   ├── product.go
│   │   └── order.go
│   └── handler/
│       ├── product.go
│       └── order.go
├── migrations/
│   ├── 000001_initial_schema.up.sql
│   ├── 000001_initial_schema.down.sql
│   ├── 000002_create_orders_table.up.sql
│   └── 000002_create_orders_table.down.sql
├── static/
│   ├── index.html
│   ├── customer.html
│   ├── styles.css
│   └── script.js
├── go.mod
└── go.sum
```

## API Endpoints

### Products
- `GET /products` - Get all products
- `GET /products/{id}` - Get product by ID
- `POST /products` - Create new product
- `PUT /products/{id}` - Update product
- `DELETE /products/{id}` - Delete product

### Orders
- `GET /orders` - Get all orders
- `POST /orders` - Create new order

## License

This project is licensed under the MIT License. 