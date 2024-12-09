Product Management Backend System
This is a backend application designed for managing products. It provides APIs to create, retrieve, and manage product data, along with features like caching (Redis), asynchronous processing (RabbitMQ), and a relational database (PostgreSQL).

Features
🚀 RESTful API for product management.
🛠 Create and retrieve product details.
📦 Asynchronous image processing using RabbitMQ.
⚡ Caching with Redis for faster responses.
📝 Structured logging for easier debugging.
🛡 Robust error handling with retries.
Tech Stack
Language: Golang
Database: PostgreSQL
Caching: Redis
Message Queue: RabbitMQ
Framework: Gin Web Framework
Getting Started
1. Prerequisites
Golang (version 1.18 or higher)
PostgreSQL
Redis
RabbitMQ
2. Installation
Step 1: Clone the Repository
bash
Copy code
git clone <repository_url>
cd product-management-system
Step 2: Set Up Environment Variables
Create a .env file in the project root:

env
Copy code
DB_URL=postgres://username:password@localhost:5432/product_management
REDIS_URL=localhost:6379
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
Replace username and password with your PostgreSQL credentials.

Step 3: Install Dependencies
bash
Copy code
go mod tidy
Step 4: Set Up PostgreSQL
Run the following SQL script to set up the database schema:

sql
Copy code
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
    product_images TEXT[],
    compressed_product_images TEXT[],
    product_price DECIMAL(10, 2),
    created_at TIMESTAMP DEFAULT NOW()
);
Step 5: Start Redis
bash
Copy code
redis-server
Step 6: Start RabbitMQ
Enable the RabbitMQ management plugin:

bash
Copy code
rabbitmq-plugins enable rabbitmq_management
Access the management UI at http://localhost:15672/ (default credentials: guest / guest).

3. Running the Application
Start the Server
bash
Copy code
go run main.go
Run RabbitMQ Consumer
Start the image processing service:

bash
Copy code
go run queue/consumer.go
API Endpoints
1. POST /products
Description: Add a new product.
Request Body:
json
Copy code
{
    "user_id": 1,
    "product_name": "Sample Product",
    "product_description": "This is a test product.",
    "product_images": ["http://example.com/image1.jpg", "http://example.com/image2.jpg"],
    "product_price": 49.99
}
Response:
json
Copy code
{
    "message": "Product created successfully"
}
2. GET /products/:id
Description: Fetch product details by ID.
Example Request: GET http://localhost:8080/products/1
Response:
json
Copy code
{
    "id": 1,
    "user_id": 1,
    "product_name": "Sample Product",
    "product_description": "This is a test product.",
    "product_images": ["http://example.com/image1.jpg"],
    "compressed_product_images": ["http://example.com/compressed_image1.jpg"],
    "product_price": 49.99
}
3. GET /products
Description: Retrieve all products with optional filters.
Example Request:
bash
Copy code
GET http://localhost:8080/products?user_id=1&min_price=10&max_price=100
Response:
json
Copy code
[
    {
        "id": 1,
        "product_name": "Sample Product",
        "product_price": 49.99
    }
]
Project Structure
bash
Copy code
product-management-system/
├── cache/
│   └── redis.go           # Redis caching logic
├── database/
│   └── connection.go      # PostgreSQL connection
├── handlers/
│   └── products.go        # API handlers
├── queue/
│   ├── producer.go        # RabbitMQ producer
│   └── consumer.go        # RabbitMQ consumer
├── router/
│   └── router.go          # API routes
├── main.go                # Application entry point
├── .env                   # Environment variables
└── README.md              # Documentation
Testing the APIs
Using Postman
Import the API endpoints into Postman.
Test the following:
POST /products
GET /products/:id
GET /products
Unit Testing
Run the following command to test the application:

bash
Copy code
go test ./...
Troubleshooting
1. Redis Issues
Ensure Redis is running:
bash
Copy code
redis-server
Test Redis connection:
bash
Copy code
redis-cli ping
2. RabbitMQ Issues
Verify RabbitMQ is running: Access http://localhost:15672/.
3. Database Issues
Check PostgreSQL connection:
bash
Copy code
psql -U postgres -d product_management
Future Enhancements
Implement user authentication.
Add image upload functionality.
Introduce pagination for product listings.
Author
Rahul Vijay Singh
Email: rahulvijay_singh@srmap.edu.in
