CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE
);
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    product_name VARCHAR(100),
    product_description TEXT,
    product_images TEXT [],
    -- Array of image URLs
    compressed_product_images TEXT [],
    -- Array of processed image URLs
    product_price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT NOW()
);