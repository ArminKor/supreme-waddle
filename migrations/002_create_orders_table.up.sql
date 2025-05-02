CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    product_id UUID REFERENCES products(id),
    customer_name VARCHAR(255) NOT NULL,
    customer_address TEXT NOT NULL,
    quantity INTEGER NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 