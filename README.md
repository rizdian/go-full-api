```markdown
# Go Full API

This project is a full-featured REST API built with **Go**, using **Gin** as the web framework, **GORM** for ORM, **PostgreSQL** as the database, and **Docker** for containerization. It demonstrates the implementation of CRUD operations, authentication, product ordering with stock validation, and error handling.

## Features

- User authentication (JWT)
- Product management with stock tracking
- Order management with stock validation
- RESTful API for managing users, products, and orders
- Full-featured API with validation and error handling
- Docker and Docker Compose for containerization

## Installation

### Prerequisites

- **Go** (v1.18+)
- **PostgreSQL**
- **Docker** (optional but recommended for containerization)

### Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/rizdian/go-full-api.git
    cd go-full-api
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Configure the `.env` file for environment variables. Example:

    ```ini
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_user
    DB_PASSWORD=your_password
    DB_NAME=your_db
    JWT_SECRET=your_jwt_secret_key
    ```

4. (Optional) Run the application using Docker Compose:

    ```bash
    docker-compose up --build
    ```

5. If you are not using Docker, you can run the application locally:

    ```bash
    go run cmd/main.go
    ```

6. The API will be running on `http://localhost:8080`.

## API Endpoints

### Authentication

- `POST /login` - Login with email and password and get a JWT token.

  Example request:

  ```json
  {
    "email": "user@example.com",
    "password": "password"
  }
  ```

  Example response:

  ```json
  {
    "token": "your_jwt_token"
  }
  ```

### Users

- `POST /users` - Create a new user.

  Example request:

  ```json
  {
    "email": "user@example.com",
    "password": "password",
    "name": "John Doe"
  }
  ```

  Example response:

  ```json
  {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
  ```

- `GET /users/:id` - Get user by ID.

  Example request:

  ```http
  GET /users/1
  ```

  Example response:

  ```json
  {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
  ```

### Products

- `POST /products` - Create a new product.

  Example request:

  ```json
  {
    "name": "Product A",
    "price": 20.5,
    "stock": 100
  }
  ```

  Example response:

  ```json
  {
    "id": 1,
    "name": "Product A",
    "price": 20.5,
    "stock": 100
  }
  ```

- `GET /products` - Get all products.

  Example response:

  ```json
  [
    {
      "id": 1,
      "name": "Product A",
      "price": 20.5,
      "stock": 100
    }
  ]
  ```

### Orders

- `POST /orders` - Create a new order and validate stock.

  Example request:

  ```json
  {
    "user_id": 1,
    "product_id": 1,
    "quantity": 2
  }
  ```

  Example response:

  ```json
  {
    "id": 1,
    "user_id": 1,
    "product_id": 1,
    "quantity": 2,
    "total": 41.0
  }
  ```

- `GET /orders` - Get all orders.

  Example response:

  ```json
  [
    {
      "id": 1,
      "user_id": 1,
      "product_id": 1,
      "quantity": 2,
      "total": 41.0
    }
  ]
  ```

## Docker

To run the application in a Docker container, you can use the following steps:

1. Build and run with Docker Compose:

    ```bash
    docker-compose up --build
    ```

2. This will start the application and PostgreSQL in containers. The application will be available on `http://localhost:8080`.

## Testing

You can test the API using Postman or curl to make HTTP requests to the provided endpoints.

### Example with curl:

- **Create a user**:

    ```bash
    curl -X POST http://localhost:8080/users -d '{"email": "user@example.com", "password": "password", "name": "John Doe"}' -H "Content-Type: application/json"
    ```

- **Login**:

    ```bash
    curl -X POST http://localhost:8080/login -d '{"email": "user@example.com", "password": "password"}' -H "Content-Type: application/json"
    ```

- **Create an order**:

    ```bash
    curl -X POST http://localhost:8080/orders -d '{"user_id": 1, "product_id": 1, "quantity": 2}' -H "Content-Type: application/json" -H "Authorization: Bearer <your_jwt_token>"
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```
